package main

import (
	"strconv"
	"testing"
)

func TestSlice(t *testing.T) {
	var act int
	act = sumSlice("11")
	if act != 2 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 11")
	}
	act = sumSlice("111")
	if act != 3 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 111")
	}
	act = sumSlice("152")
	if act != 8 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 152")
	}
}

func TestSolve(t *testing.T) {
	var act int
	act = solve(20, 2, 5)
	if act != 84 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 20,2,5")
	}
	act = solve(10, 1, 2)
	if act != 13 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 10,1,2")
	}
	act = solve(100, 4, 16)
	if act != 4554 {
		t.Fatalf("failed test " + strconv.Itoa(act) + ": 100,4,16")
	}
}
