package main

import "fmt"

type levelZero struct {
	next examination
}

func (i *levelZero) execute(l *intern) {
	fmt.Println(l.name, " successfully passed L0!")
	l.levelZeroDone = true
	i.next.execute(l)
}
func (i *levelZero) setNext(next examination) {
	i.next = next
}
