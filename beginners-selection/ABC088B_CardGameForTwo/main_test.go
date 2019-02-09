package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	var act int
	act = solve([]int{3, 1})
	if act != 2 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 3,1")
	}
	act = solve([]int{2, 7, 4})
	if act != 5 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 2,7,4")
	}
	act = solve([]int{20, 18, 2, 18})
	if act != 18 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 20,18,2,18")
	}
	act = solve([]int{1,2,3,4,5})
	if act != 3 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 1,2,3,4,5")
	}
}
