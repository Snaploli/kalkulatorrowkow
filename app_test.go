package main

import (
	"testing"
)

func TestFindClosestGroove(t *testing.T) {
	// Test R20 matching
	// R20 standard: PitchDia: 68.27, Width: 7.95, GrooveWidth: 8.74
	g1 := FindClosestGroove(68.3, 8.7, "AUTO")
	if g1.Name != "R20" {
		t.Errorf("Expected R20, got %s", g1.Name)
	}

	g2 := FindClosestGroove(68.3, 8.0, "R")
	if g2.Name != "R20" {
		t.Errorf("Expected R20, got %s", g2.Name)
	}

	// Test BX150 matching
	// BX150 standard: PitchDia: 62.89, Width: 9.30, GrooveWidth: 11.43
	g3 := FindClosestGroove(63.0, 9.3, "AUTO")
	if g3.Name != "BX150" {
		t.Errorf("Expected BX150, got %s", g3.Name)
	}

	g4 := FindClosestGroove(63.0, 11.4, "BX")
	if g4.Name != "BX150" {
		t.Errorf("Expected BX150, got %s", g4.Name)
	}
}
