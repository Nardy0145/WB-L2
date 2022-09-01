package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	f int
	d string
	s bool
)

func main() {
	flag.IntVar(&f, "f", 0, "fields")
	flag.StringVar(&d, "d", "", "delimiter")
	flag.BoolVar(&s, "s", false, "separated")
	flag.Parse()
	if f <= 0 {
		log.Fatal("error: flag f must be bigger than 0")
	}
	parseFile("text.txt")
}

func parseFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("error: file open: ", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("error: file close: ", err)
		}
	}(file)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		str := scan.Text()
		sl := strings.Split(str, " ")
		// Проверка на индекс поля к заданному параметру
		if len(sl) < f {
			continue
		} else {
			// Если подходит, выводим строку после обработки
			fmt.Println(Parse(str))
		}
	}
}

// Parse ...
func Parse(str string) string {
	// Если задан флаг -s и полученная строка не содержит кастом разделитель - возвращаем пустоту
	if s && !strings.Contains(str, d) {
		return ""
	}
	// Разбиваем строку на кастом разделитель & сверяем индексы и если меньше заданного, возвращаем инпут
	sl := strings.Split(str, d)
	if len(sl) < f {
		return str
	}
	// Либо по заданному "индексу"
	return sl[f-1]
}
