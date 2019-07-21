package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var act float64
	act = solve(6, 2)
	if act != 2 {
		t.Fatalf("failed test: %v\n", act)
	}
	act = solve(14, 3)
	if act != 2 {
		t.Fatalf("failed test: %v\n", act)
	}
	act = solve(20, 4)
	if act != 3 {
		t.Fatalf("failed test: %v\n", act)
	}
}
