package main

import (
	"strconv"
	"testing"
)

func TestSolveCoins(t *testing.T) {
	var act int
	act = solveCoins(2, 2, 2, 100)
	if act != 2 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 2,2,2,100")
	}
	act = solveCoins(5, 1, 0, 150)
	if act != 0 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 5,1,0,150")
	}
	act = solveCoins(30, 40, 50, 6000)
	if act != 213 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 30,40,50,6000")
	}
}
