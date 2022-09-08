package main

import "fmt"

type levelSecond struct {
	next examination
}

func (i *levelSecond) execute(l *intern) {
	fmt.Println(l.name, " successfully passed L2! Congratulations!")
	l.levelSecondDone = true
}

func (i *levelSecond) setNext(next examination) {
	i.next = next
}
