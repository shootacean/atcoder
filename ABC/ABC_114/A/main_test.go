package main

import (
	"testing"
)

func TestIs753(t *testing.T) {
	// True
	if ! is753(7) {
		t.Fatalf("failed test %#v", 7)
	}
	if ! is753(5) {
		t.Fatalf("failed test %#v", 5)
	}
	if ! is753(3) {
		t.Fatalf("failed test %#v", 3)
	}
	// False
	if is753(1) {
		t.Fatalf("failed test %#v", 1)
	}
	if is753(2) {
		t.Fatalf("failed test %#v", 2)
	}
	if is753(6) {
		t.Fatalf("failed test %#v", 6)
	}
	if is753(9) {
		t.Fatalf("failed test %#v", 9)
	}
	if is753(10) {
		t.Fatalf("failed test %#v", 10)
	}
	if is753(11) {
		t.Fatalf("failed test %#v", 11)
	}
	if is753(15) {
		t.Fatalf("failed test %#v", 15)
	}
}
