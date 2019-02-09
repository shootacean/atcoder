package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var act int
	act = solve(4, 2, 6)
	if act != 7 {
		t.Fatalf("Failed test: %v: 4,2,6", act)
	}
	act = solve(7,3,4)
	if act != 8 {
		t.Fatalf("Failed test: %v: 7,3,4", act)
	}
	act = solve(314159265,35897932,384626433)
	if act != 48518828981938099 {
		t.Fatalf("Failed test: %v: 314159265,35897932,384626433", act)
	}
}
