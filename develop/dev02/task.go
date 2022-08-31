package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Fatal("input error: ", err)
	}
	v, err := FormatStr(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}

// FormatStr ...Formats string
func FormatStr(s string) (string, error) {

	if s == "" {
		return "", nil
	}
	var check int
	rString := []rune(s)
	for _, v := range rString {
		if unicode.IsLetter(v) {
			check++
		}
	}
	if check == 0 {
		err1 := errors.New("invalid string")
		return "", err1
	}
	if len(rString) == check {
		return string(rString), nil
	}
	var output []rune
	for i, v := range rString {
		if unicode.IsLetter(v) {
			if i != len(rString)-1 {
				if z, err := strconv.Atoi(string(rString[i+1])); err == nil {
					for iters := 0; iters < z; iters++ {
						output = append(output, v)
					}
				} else {
					output = append(output, v)
				}
			} else {
				output = append(output, v)
			}
		} else {
			switch i {
			case 0:
				continue
			case len(rString) - 1:
				if !unicode.IsLetter(rString[i-1]) {
					output = append(output, rString[i])
					break
				}
			default:
				if !unicode.IsLetter(rString[i-1]) {
					output = append(output, rString[i])
				}
			}
		}
	}
	return string(output), nil
}
