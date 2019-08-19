package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var act float64
	
	act = solve([]float64{3, 4})
	if act != 3.5 {
		t.Fatalf("failed test")
	}
	
	act = solve([]float64{138, 138, 138, 138, 138})
	if act != 138 {
		t.Fatalf("failed test")
	}
	
	act = solve([]float64{500, 200, 300})
	if act != 375 {
		t.Fatalf("failed test")
	}
}
