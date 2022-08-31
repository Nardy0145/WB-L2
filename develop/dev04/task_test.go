package main

import (
	"strings"
	"testing"
)

func TestMakeMap(t *testing.T) {
	line := "слиток тяпка листок пятак столик пятка я в поо"
	lLine := strings.ToLower(line)
	sLine := strings.Split(lLine, " ")
	anagrams, err := MakeMap(sLine)
	if err != nil {
		t.Fatal(err)
	}
	expectedRes := make(map[string][]string)
	expectedRes["слиток"] = append(expectedRes["слиток"], "листок")
	expectedRes["слиток"] = append(expectedRes["слиток"], "слиток")
	expectedRes["слиток"] = append(expectedRes["слиток"], "столик")
	expectedRes["тяпка"] = append(expectedRes["слиток"], "пятак")
	expectedRes["тяпка"] = append(expectedRes["слиток"], "пятка")
	expectedRes["тяпка"] = append(expectedRes["слиток"], "тяпка")
	fSlice := makeSlice(anagrams)
	sSlice := makeSlice(anagrams)
	fStr := strings.Join(fSlice, "")
	sStr := strings.Join(sSlice, "")
	if fStr != sStr {
		t.Fatal("error: unexpected result!")
	}
}

func makeSlice(x map[string][]string) []string {
	var tmpSlice []string
	for _, v := range x {
		tmpSlice = append(tmpSlice, v[0])
	}
	return tmpSlice
}
