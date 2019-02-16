package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var act string
	act = solve("ch@ku@ai", "choku@@i")
	if act != "You can win" {
		t.Fatalf("failed test: %v\n", act)
	}
	act = solve("aoki", "@ok@")
	if act != "You will lose" {
		t.Fatalf("failed test: %v\n", act)
	}
}
