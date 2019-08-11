package main

import (
	"testing"
)

func TestIs753(t *testing.T) {
	// Sample
	if solve(81, 58) != 110 {
		t.Fatalf("failed test %v, %v", 81, 58)
	}
	if solve(4, 54) != 31 {
		t.Fatalf("failed test %v, %v", 4, 54)
	}
	
	if solve(20, 10) != 25 {
		t.Fatalf("failed test %v, %v", 20, 10)
	}
}
