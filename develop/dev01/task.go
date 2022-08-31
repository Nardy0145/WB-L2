package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

const host = "ntp1.stratum2.ru"

func main() {
	curTime, err := GetTime(host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(curTime)
	os.Exit(1)
}

// GetTime ...get actual time +3.00GMT
func GetTime(host string) (string, error) {
	curTime, err := ntp.Time(host)
	if err != nil {
		return "", err
	}
	return curTime.Format(time.RFC3339), nil
}
