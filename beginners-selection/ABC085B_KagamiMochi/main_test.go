package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	var act int
	act = solve([]int{10, 8, 8, 6})
	if act != 3 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 10,8,8,6")
	}
	act = solve([]int{15, 15, 15})
	if act != 1 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 1,15,15,15")
	}
	act = solve([]int{50, 30, 50, 100, 50, 80, 30})
	if act != 4 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 50,30,50,100,50,80,30")
	}
}
