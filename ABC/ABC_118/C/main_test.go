package main

import (
	"testing"
)

func TestGcd(t *testing.T) {
	var act int
	act = gcd(4, 40)
	if act != 4 {
		t.Fatalf("failed test: %v\n", act)
	}
}

func TestSolveSample(t *testing.T) {
	var act int
	act = solve(4, []int{2, 10, 8, 40})
	if act != 2 {
		t.Fatalf("failed test: %v\n", act)
	}
	act = solve(4, []int{5, 13, 8, 1000000000})
	if act != 1 {
		t.Fatalf("failed test: %v\n", act)
	}
	act = solve(3, []int{1000000000, 1000000000, 1000000000})
	if act != 1000000000 {
		t.Fatalf("failed test: %v\n", act)
	}
}

func TestSolve(t *testing.T) {
	var act int
	act = solve(1, []int{9999999999, 9999999999, 9999999999})
	if act != 9999999999 {
		t.Fatalf("failed test: %v\n", act)
	}
	act = solve(1, []int{9999999999})
	if act != 9999999999 {
		t.Fatalf("failed test: %v\n", act)
	}
	act = solve(2, []int{3, 6})
	if act != 3 {
		t.Fatalf("failed test: %v\n", act)
	}
}
