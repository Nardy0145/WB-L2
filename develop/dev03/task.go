package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var (
		n, r, u bool
		k       int
	)

	flag.BoolVar(&n, "n", false, "num value sort")
	flag.BoolVar(&r, "r", false, "reverse sort")
	flag.BoolVar(&u, "u", false, "no repeated-str sort")
	flag.IntVar(&k, "k", 0, "table value sort")
	flag.Parse()
	lines, err := readFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	sorted, err := DoSort(lines, n, r, u, k)
	if err != nil {
		log.Fatal(err)
	}
	saveFile(sorted)
}

func readFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	scan := bufio.NewScanner(f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("file to read hasn't been closed")
		}
	}(f)
	var strs []string
	for scan.Scan() {
		str := scan.Text()
		strs = append(strs, str)
	}
	return strs, nil
}

func saveFile(lines []string) {
	f, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("save file hasn't been closed")
		}
	}(f)
	for _, v := range lines {
		_, err := f.WriteString(v + "\n")
		if err != nil {
			return
		}
	}
}

// DoSort ...sorts input string slice
func DoSort(lines []string, n, r, u bool, k int) ([]string, error) {
	var nr bool
	tmpSlice := make([]string, len(lines))
	copy(tmpSlice, lines)
	if k > 0 {
		var nSlice []string
		for _, v := range tmpSlice {
			lineSlice := strings.Split(v, " ")
			if len(lineSlice) < k {
				continue
			} else {
				lineStr := strings.Join(lineSlice[k-1:], " ")
				nSlice = append(nSlice, lineStr)
			}
		}
		tmpSlice = tmpSlice[:0]
		tmpSlice = nSlice[:]
	}
	if n {
		var tmpWords, tmpInts []string
		var tmpWord []rune
		for _, v := range tmpSlice {
			tmpWord = []rune(v)
			if unicode.IsLetter(tmpWord[0]) {
				tmpWords = append(tmpWords, v)
			} else {
				tmpInts = append(tmpInts, v)
			}
		}
		if r {
			sort.Sort(sort.Reverse(sort.StringSlice(tmpInts)))
			nr = true
		} else {
			sort.Strings(tmpInts)
		}
		tmpSlice = append(tmpInts, tmpWords...)
	}
	if u {
		allKeys := make(map[string]bool)
		var list []string
		for _, item := range tmpSlice {
			if _, value := allKeys[item]; !value {
				allKeys[item] = true
				list = append(list, item)
			}
		}
		fmt.Println(list)
	}
	if !n {
		sort.Strings(tmpSlice)
	}
	if r && nr == false {
		sort.Sort(sort.Reverse(sort.StringSlice(tmpSlice)))
	}
	return tmpSlice, nil
}
