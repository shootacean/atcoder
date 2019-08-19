package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	if solve(-13, 3) != -10 {
		t.Fatalf("failed test")
	}
	if solve(1, -33) != 34 {
		t.Fatalf("failed test")
	}
	if solve(13, 3) != 39 {
		t.Fatalf("failed test")
	}
}
