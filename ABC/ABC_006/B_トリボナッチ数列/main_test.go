package main

import "testing"

func TestSolve(t *testing.T) {
	var act int
	
	act = solve(1)
	if act != 0 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(2)
	if act != 0 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(3)
	if act != 1 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(4)
	if act != 1 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(5)
	if act != 2 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(6)
	if act != 4 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(7)
	if act != 7 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(8)
	if act != 13 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(9)
	if act != 24 {
		t.Fatalf("failed test %d", act)
	}
	
	act = solve(100000)
	if act != 7927 {
		t.Fatalf("failed test %d", act)
	}
}

//func Testtori_memo(t *testing.T) {
//	var act int
//
//	act = tori_memo(1)
//	if act != 0 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(2)
//	if act != 0 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(3)
//	if act != 1 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(4)
//	if act != 1 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(5)
//	if act != 2 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(6)
//	if act != 4 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(7)
//	if act != 7 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(8)
//	if act != 13 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(9)
//	if act != 24 {
//		t.Fatalf("failed test %d", act)
//	}
//
//	act = tori_memo(100000)
//	if act != 5320677146059475665 {
//		t.Fatalf("failed test %d", act)
//	}
//}
