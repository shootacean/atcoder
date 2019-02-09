package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var x, y, z int
	// TDD
	x, y, z = solve(1, 10000)
	if x != 1 || y != 0 || z != 0 {
		t.Fatalf("failed test %d, %d, %d: 1,10000", x, y, z)
	}
	x, y, z = solve(2, 20000)
	if x != 2 || y != 0 || z != 0 {
		t.Fatalf("failed test %d, %d, %d: 2, 20000", x, y, z)
	}
	x, y, z = solve(2, 15000)
	if x != 1 || y != 1 || z != 0 {
		t.Fatalf("failed test %d, %d, %d: 2, 15000", x, y, z)
	}
	x, y, z = solve(3, 20000)
	if x != 1 || y != 2 || z != 0 {
		t.Fatalf("failed test %d, %d, %d: 3, 20000", x, y, z)
	}
	x, y, z = solve(4, 21000)
	if x != 1 || y != 2 || z != 1 {
		t.Fatalf("failed test %d, %d, %d: 4, 21000", x, y, z)
	}
	x, y, z = solve(2, 5000)
	if x != -1 || y != -1 || z != -1 {
		t.Fatalf("failed test %d, %d, %d: 2, 20000", x, y, z)
	}
	// Sample
	x, y, z = solve(20, 196000)
	if x != -1 || y != -1 || z != -1 {
		t.Fatalf("failed test %d, %d, %d: 20, 1960000", x, y, z)
	}
	x, y, z = solve(2000, 20000000)
	if x != 2000 || y != 0 || z != 0 {
		t.Fatalf("failed test %d, %d, %d: 2000, 20000000", x, y, z)
	}
	x, y, z = solve(1000, 1234000)
	if x == -1 || y == -1 || z == -1 {
		t.Fatalf("failed test %d, %d, %d: 1000, 1234000", x, y, z)
	}
}
