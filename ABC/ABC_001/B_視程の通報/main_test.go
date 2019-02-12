package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var act string
	act = solve(100)
	if act != "00" {
		t.Fatalf("failed test %v : 100", act)
	}
	act = solve(200)
	if act != "02" {
		t.Fatalf("failed test %v : 200", act)
	}
	act = solve(2000)
	if act != "20" {
		t.Fatalf("failed test %v : 2000", act)
	}
	act = solve(30000)
	if act != "80" {
		t.Fatalf("failed test %v : 30000", act)
	}
	act = solve(40000)
	if act != "82" {
		t.Fatalf("failed test %v : 40000", act)
	}
	act = solve(75000)
	if act != "89" {
		t.Fatalf("failed test %v : 75000", act)
	}
}
