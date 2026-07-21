package main

import (
	"context"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// CalculationResult holds the return values for the frontend
type CalculationResult struct {
	P             float64 `json:"p"`
	Q             float64 `json:"q"`
	External      float64 `json:"external"`
	GrooveType    string  `json:"grooveType"`
	SuggestedName string  `json:"suggestedName"`
	SuggestedP    float64 `json:"suggestedP"`
	SuggestedW    float64 `json:"suggestedW"`
	SuggestedGW   float64 `json:"suggestedGw"`
	ErrMsg        string  `json:"errMsg"`
}

// CalculatePandQ performs calculations and returns a CalculationResult struct
func (a *App) CalculatePandQ(external float64, internal float64, grooveType string) CalculationResult {
	errMsg := ""
	var p float64
	var q float64
	if external <= internal {
		errMsg = "Zewnętrzny musi być większy od wewnętrznego"
		return CalculationResult{
			P:          0,
			Q:          0,
			ErrMsg:     errMsg,
		}
	}
	if external <= 0 || internal <= 0 {
		errMsg = "Wymiary muszą być większe od 0"
		return CalculationResult{
			P:          0,
			Q:          0,
			ErrMsg:     errMsg,
		}
	}

	q = (external - internal + 15) / 2
	p = external - q

	closest := FindClosestGroove(p, q, grooveType)

	resGrooveType := grooveType
	if resGrooveType == "" || resGrooveType == "AUTO" {
		if strings.HasPrefix(closest.Name, "BX") {
			resGrooveType = "BX"
		} else {
			resGrooveType = "R"
		}
	}

	return CalculationResult{
		P:             p,
		Q:             q,
		External:      external,
		GrooveType:    resGrooveType,
		SuggestedName: closest.Name,
		SuggestedP:    closest.PitchDia,
		SuggestedW:    closest.Width,
		SuggestedGW:   closest.GrooveWidth,
		ErrMsg:        "",
	}
}
