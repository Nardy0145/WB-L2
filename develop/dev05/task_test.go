package main

import (
	"testing"
)

func TestRefactorByFlags(t *testing.T) {
	i = true
	n = true
	F = true
	lines := ReadFile("text.txt")
	pattern := "test"
	x := RefactorByFlags(lines, pattern)
	if x[8] != "test" || x[20] != "test" {
		t.Fatal("unexpected result!")
	}
}
