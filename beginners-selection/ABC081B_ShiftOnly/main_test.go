package main

import (
	"strconv"
	"testing"
)

func TestIsEvenAll(t *testing.T) {
	if IsEvenAll([]int{1, 2, 3}) {
		t.Fatalf("failed test : 1,2,3")
	}
	if !IsEvenAll([]int{2, 4, 6}) {
		t.Fatalf("failed test : 2,4,6")
	}
}

func TestCountHalfLimit(t *testing.T) {
	if CountHalfLimit([]int{2, 4, 6}) != 1 {
		t.Fatalf("failed test : 2,4,6")
	}
	if CountHalfLimit([]int{4, 8, 12}) != 2 {
		t.Fatalf("failed test : 4,8,12")
	}
	if CountHalfLimit([]int{5, 6, 8, 10}) != 0 {
		t.Fatalf("failed test : 4,8,12")
	}
	act := CountHalfLimit([]int{
		382253568, 723152896, 37802240, 379425024, 404894720, 471526144,
	})
	if act != 8 {
		t.Fatalf("failed test " + strconv.Itoa(act) + " : 382253568, 723152896, 37802240, 379425024, 404894720, 471526144")
	}
}
