package main

import (
	"reflect"
	"testing"
)

func TestDoSort(t *testing.T) {
	lines, _ := ReadFile("text.txt")
	var n, r, u bool
	var k int
	sort, err := DoSort(lines, n, r, u, k)
	if err != nil {
		t.Fatal(err)
	}
	resLines, _ := ReadFile("test_result.txt")
	if !reflect.DeepEqual(sort, resLines) {
		t.Fatal("error: sorted output is wrong!")
	}
}
