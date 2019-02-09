package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var act string
	act = solve(3, 2)
	if act != "YES" {
		t.Fatalf("Failed test:  %s: 3,2", act)
	}
	act = solve(5,5)
	if act != "NO" {
		t.Fatalf("Failed test:  %s: 5,5", act)
	}
	act = solve(31,10)
	if act != "YES" {
		t.Fatalf("Failed test:  %s: 31,10", act)
	}
	act = solve(10, 90)
	if act != "NO" {
		t.Fatalf("Failed test:  %s: 10,90", act)
	}
}
