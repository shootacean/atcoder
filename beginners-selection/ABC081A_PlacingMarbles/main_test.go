package main

import (
    "strings"
    "testing"
)

func TestStringsCount(t *testing.T) {
    if strings.Count("101", "1") != 2 {
        t.Fatalf("failed test %#v", 2)
    }
    if strings.Count("000", "1") != 0 {
        t.Fatalf("failed test %#v", 0)
    }
}
