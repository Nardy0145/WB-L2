package main

import (
	"testing"
)

func TestFormatStr(t *testing.T) {
	strs := []string{"a4bc2d5e", "abcd", "", "45"}
	for i, v := range strs {
		switch i {
		case 0:
			if _, err := FormatStr(v); err != nil {
				t.Fatalf("Error: %s, String: %s", err, v)
			}
		case 1:
			if _, err := FormatStr(v); err != nil {
				t.Fatalf("Error: %s, String: %s", err, v)
			}
		case 2:
			if _, err := FormatStr(v); err != nil {
				t.Fatalf("Error: %s, String: %s", err, v)
			}
		case 3:
			if _, err := FormatStr(v); err != nil {
				t.Fatalf("Error: %s, String: %s", err, v)
			}
		}
	}
}
