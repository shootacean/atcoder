package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	// Yes
	if ! solve(3, 1) {
		t.Fatalf("failed test")
	}
	// No
	if solve(1, 2) {
		t.Fatalf("failed test")
	}
	if solve(2, 2) {
		t.Fatalf("failed test")
	}
}
