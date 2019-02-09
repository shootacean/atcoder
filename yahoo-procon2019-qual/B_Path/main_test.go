package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var act string
	act = solve(4, 2, 1, 3, 2, 3)
	if act != "YES" {
		t.Fatalf("Failed test:  %s: 4,2,1,3,2,3", act)
	}
	act = solve(3, 2, 2, 4, 1, 2)
	if act != "NO" {
		t.Fatalf("Failed test:  %s: 3,2,2,4,1,2", act)
	}
	act = solve(2, 1, 3, 2, 4, 3)
	if act != "YES" {
		t.Fatalf("Failed test:  %s: 2,1,3,2,4,3", act)
	}
	act = solve(2, 1, 3, 2, 4, 3)
	if act != "YES" {
		t.Fatalf("Failed test:  %s: 2,1,3,2,4,3", act)
	}
}
