package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "Line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	// B ..flag
	B int
	// C ..flag
	C int
	// A ..flag
	A int
	// F ..flag
	F          bool
	c, i, v, n bool
)

// Line ...just a model
type Line struct {
	text string
}

func main() {
	parseFlags()
	// Слово для поиска
	pattern := "Test"
	lines := ReadFile("text.txt")
	_ = RefactorByFlags(lines, pattern)
}

// ReadFile ...reads txt file into model
func ReadFile(name string) []Line {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("reading file hasn't been closed")
		}
	}(f)
	i := 0
	var lines []Line
	for scan.Scan() {
		i++
		lines = append(lines, Line{text: scan.Text()})
	}
	return lines
}

func parseFlags() {
	flag.IntVar(&A, "A", 0, "after")
	flag.IntVar(&B, "B", 0, "before")
	flag.IntVar(&C, "C", 0, "a + b")
	flag.BoolVar(&c, "c", false, "count")
	flag.BoolVar(&i, "i", false, "ignore-case")
	flag.BoolVar(&v, "v", false, "invert")
	flag.BoolVar(&F, "F", false, "fixed")
	flag.BoolVar(&n, "n", false, "Line num")
	flag.Parse()
}

// RefactorByFlags ..grep analogue
func RefactorByFlags(lines []Line, patter string) map[int]string {
	// -c - count - количествово строк
	if c {
		fmt.Println(len(lines))
		return nil
	}
	source := make(map[int]string)
	result := make(map[int]string)
	var pattern string

	// -i - ignore-case - игнорировать регистр
	if i {
		pattern = strings.ToLower(patter)
		for ind, val := range lines {
			lines[ind].text = strings.ToLower(val.text)
		}
		for ind, val := range lines {
			source[ind+1] = val.text
		}
	} else {
		pattern = patter
		for ind, val := range lines {
			source[ind+1] = val.text
		}
	}

	// -A - "after" печатать +N строк после совпадения
	if A > 0 && A < len(lines) && C == 0 {
		for ind, val := range lines {
			if strings.Contains(val.text, pattern) {
				result[ind+1] = val.text
				for x := 1; x <= A; x++ {
					if source[ind+1+x] == "" {
						break
					} else {
						result[ind+1+x] = source[ind+1+x]
					}
				}
			}
		}
	}

	// -B - "before" печатать -N строк после совпадения
	if B > 0 && B < len(lines) && C == 0 {
		for ind, val := range lines {
			if strings.Contains(val.text, pattern) {
				result[ind+1] = val.text
				for x := 1; x <= B; x++ {
					if source[ind+1-x] == "" {
						break
					} else {
						result[ind+1-x] = source[ind+1-x]
					}
				}
			}
		}
	}

	// -C - "context" печатать +-N строк после совпадения
	if C > 0 && C < len(lines) {
		for ind, val := range lines {
			if strings.Contains(val.text, pattern) {
				result[ind+1] = val.text
				for x := 1; x <= C; x++ {
					if source[ind+1+x] == "" || source[ind+1-x] == "" {
						if source[ind+1+x] == "" && source[ind+1-x] == "" {
							break
						}
						continue
					} else {
						result[ind+1+x] = source[ind+1+x]
						result[ind+1-x] = source[ind+1-x]
					}
				}
			}
		}
	}

	// -n - "Line num", печатать номер строки
	// -F - "fixed", точное совпадение со строкой, не паттерн
	if F {
		for ind, val := range lines {
			if val.text == pattern {
				result[ind+1] = fmt.Sprintf("%s", val.text)
			}
		}
	}
	// -v - "invert" (вместо совпадения, исключать)
	if v {
		for ind, val := range lines {
			if !strings.Contains(val.text, pattern) {
				result[ind+1] = fmt.Sprintf("%s", val.text)
			}
		}
	}
	// просто фильтрация
	if A == 0 && B == 0 && C == 0 && !c && !v && !F {
		for ind, val := range lines {
			if strings.Contains(val.text, pattern) {
				result[ind+1] = fmt.Sprintf("%s", val.text)
			}
		}
	}
	var sortedKeys []int
	for k := range result {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)
	// -n нумерация строк
	if n {
		for ind := range sortedKeys {
			fmt.Println(sortedKeys[ind], result[sortedKeys[ind]])
		}
	} else {
		for ind := range sortedKeys {
			fmt.Println(result[sortedKeys[ind]])
		}
	}
	return result
}
