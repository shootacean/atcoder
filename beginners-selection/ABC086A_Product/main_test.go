package main

import "testing"

func TestIsEven(t *testing.T) {
    // even
    if ! IsEven(2) {
        t.Fatalf("failed test %#v", 2)
    }
    if ! IsEven(2*5) {
        t.Fatalf("failed test %#v", 2*5)
    }
    // odd
    if IsEven(3) {
        t.Fatalf("failed test %#v", 3)
    }
    if IsEven(13*5) {
        t.Fatalf("failed test %#v", 13*5)
    }
}

func TestGetEvenOdd(t *testing.T) {
    // even
    if GetEvenOdd(2) != "Even" {
        t.Fatalf("failed test %#v", 2)
    }
    if GetEvenOdd(2*5) != "Even" {
        t.Fatalf("failed test %#v", 2*5)
    }
    // odd
    if GetEvenOdd(3) != "Odd" {
        t.Fatalf("failed test %#v", 3)
    }
    if GetEvenOdd(13*5) != "Odd" {
        t.Fatalf("failed test %#v", 13*5)
    }
}