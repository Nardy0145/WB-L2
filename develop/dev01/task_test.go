package main

import (
	"testing"
)

func TestGetTime(t *testing.T) {
	fakeHost := "https://nothinghere.com"
	curTime, err := GetTime(host)
	if curTime == "" || err != nil {
		t.Fatal("Error: ", err)
	}
	curTime, err = GetTime(fakeHost)
	if curTime == "" || err != nil {
		t.Fatal("Error: ", err)
	}
}
