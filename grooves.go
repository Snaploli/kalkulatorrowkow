package main

import (
	"math"
	"strings"
)

type Groove struct {
	Name        string  `json:"name"`
	PitchDia    float64 `json:"pitchDia"`
	Width       float64 `json:"width"`
	GrooveWidth float64 `json:"grooveWidth"`
}

var GrooveDatabase = []Groove{
	// Type R Gaskets (R20 to R99)
	{Name: "R20", PitchDia: 68.27, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R21", PitchDia: 72.23, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R22", PitchDia: 82.55, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R23", PitchDia: 82.55, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R24", PitchDia: 95.25, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R25", PitchDia: 101.60, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R26", PitchDia: 101.60, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R27", PitchDia: 107.95, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R28", PitchDia: 111.13, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R29", PitchDia: 114.30, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R30", PitchDia: 117.48, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R31", PitchDia: 123.83, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R32", PitchDia: 127.00, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R33", PitchDia: 131.78, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R34", PitchDia: 131.78, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R35", PitchDia: 136.53, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R36", PitchDia: 149.23, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R37", PitchDia: 149.23, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R38", PitchDia: 157.18, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R39", PitchDia: 161.93, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R40", PitchDia: 171.45, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R41", PitchDia: 171.45, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R42", PitchDia: 177.80, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R43", PitchDia: 184.15, Width: 7.95, GrooveWidth: 8.74},
	{Name: "R44", PitchDia: 184.15, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R45", PitchDia: 190.50, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R46", PitchDia: 203.20, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R47", PitchDia: 215.90, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R48", PitchDia: 228.60, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R49", PitchDia: 241.30, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R50", PitchDia: 254.00, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R51", PitchDia: 266.70, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R52", PitchDia: 282.58, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R53", PitchDia: 304.80, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R54", PitchDia: 323.85, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R55", PitchDia: 342.90, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R56", PitchDia: 333.38, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R57", PitchDia: 368.30, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R58", PitchDia: 390.53, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R59", PitchDia: 396.88, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R60", PitchDia: 419.10, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R61", PitchDia: 431.80, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R62", PitchDia: 450.85, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R63", PitchDia: 469.90, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R64", PitchDia: 488.95, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R65", PitchDia: 508.00, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R66", PitchDia: 533.40, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R67", PitchDia: 558.80, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R68", PitchDia: 539.75, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R69", PitchDia: 565.15, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R70", PitchDia: 596.90, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R71", PitchDia: 622.30, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R72", PitchDia: 596.90, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R73", PitchDia: 622.30, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R74", PitchDia: 654.05, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R75", PitchDia: 679.45, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R76", PitchDia: 711.20, Width: 12.70, GrooveWidth: 13.49},
	{Name: "R77", PitchDia: 736.60, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R78", PitchDia: 768.35, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R79", PitchDia: 793.75, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R80", PitchDia: 673.10, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R81", PitchDia: 698.50, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R82", PitchDia: 812.80, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R83", PitchDia: 329.57, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R84", PitchDia: 863.60, Width: 15.88, GrooveWidth: 16.66},
	{Name: "R85", PitchDia: 889.00, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R86", PitchDia: 914.40, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R87", PitchDia: 939.80, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R88", PitchDia: 965.20, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R89", PitchDia: 990.60, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R90", PitchDia: 1016.00, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R91", PitchDia: 260.35, Width: 31.75, GrooveWidth: 33.32},
	{Name: "R92", PitchDia: 228.60, Width: 11.11, GrooveWidth: 11.91},
	{Name: "R93", PitchDia: 749.30, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R94", PitchDia: 800.10, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R95", PitchDia: 857.25, Width: 19.05, GrooveWidth: 19.84},
	{Name: "R96", PitchDia: 914.40, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R97", PitchDia: 965.20, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R98", PitchDia: 1022.35, Width: 22.23, GrooveWidth: 23.01},
	{Name: "R99", PitchDia: 234.95, Width: 11.11, GrooveWidth: 11.91},

	// Type BX Gaskets (BX150 to BX172, and BX303)
	{Name: "BX150", PitchDia: 62.89, Width: 9.30, GrooveWidth: 11.43},
	{Name: "BX151", PitchDia: 66.77, Width: 9.63, GrooveWidth: 11.84},
	{Name: "BX152", PitchDia: 74.44, Width: 10.24, GrooveWidth: 12.65},
	{Name: "BX153", PitchDia: 89.56, Width: 11.38, GrooveWidth: 14.07},
	{Name: "BX154", PitchDia: 104.44, Width: 12.40, GrooveWidth: 15.39},
	{Name: "BX155", PitchDia: 133.74, Width: 14.22, GrooveWidth: 17.65},
	{Name: "BX156", PitchDia: 219.30, Width: 18.62, GrooveWidth: 23.01},
	{Name: "BX157", PitchDia: 273.48, Width: 20.98, GrooveWidth: 25.93},
	{Name: "BX158", PitchDia: 328.90, Width: 23.14, GrooveWidth: 28.58},
	{Name: "BX159", PitchDia: 401.02, Width: 25.70, GrooveWidth: 31.75},
	{Name: "BX160", PitchDia: 388.85, Width: 13.74, GrooveWidth: 16.97},
	{Name: "BX161", PitchDia: 475.20, Width: 16.21, GrooveWidth: 20.02},
	{Name: "BX162", PitchDia: 461.27, Width: 14.22, GrooveWidth: 17.65},
	{Name: "BX163", PitchDia: 538.79, Width: 17.37, GrooveWidth: 21.51},
	{Name: "BX164", PitchDia: 545.97, Width: 24.59, GrooveWidth: 30.43},
	{Name: "BX165", PitchDia: 606.22, Width: 18.49, GrooveWidth: 22.91},
	{Name: "BX166", PitchDia: 613.89, Width: 26.14, GrooveWidth: 32.39},
	{Name: "BX167", PitchDia: 746.25, Width: 13.11, GrooveWidth: 16.21},
	{Name: "BX168", PitchDia: 749.20, Width: 16.05, GrooveWidth: 19.86},
	{Name: "BX169", PitchDia: 160.58, Width: 12.93, GrooveWidth: 15.98},
	{Name: "BX170", PitchDia: 203.81, Width: 14.22, GrooveWidth: 17.65},
	{Name: "BX171", PitchDia: 253.22, Width: 14.22, GrooveWidth: 17.65},
	{Name: "BX172", PitchDia: 318.85, Width: 14.22, GrooveWidth: 17.65},
	{Name: "BX303", PitchDia: 835.78, Width: 16.97, GrooveWidth: 21.03},
}

// FindClosestGroove finds the closest matching R or BX groove based on pitch diameter (p) and width (q).
// grooveType can be "R", "BX", or "AUTO" (or empty string for all).
func FindClosestGroove(p, q float64, grooveType string) Groove {
	var bestMatch Groove
	minDist := math.MaxFloat64

	for _, g := range GrooveDatabase {
		if grooveType == "R" && !strings.HasPrefix(g.Name, "R") {
			continue
		}
		if grooveType == "BX" && !strings.HasPrefix(g.Name, "BX") {
			continue
		}

		// Compare distance to typical groove width
		distGroove := math.Pow(p-g.PitchDia, 2) + math.Pow(q-g.GrooveWidth, 2)
		// Compare distance to typical gasket width
		distGasket := math.Pow(p-g.PitchDia, 2) + math.Pow(q-g.Width, 2)

		dist := math.Min(distGroove, distGasket)
		if dist < minDist {
			minDist = dist
			bestMatch = g
		}
	}
	return bestMatch
}
