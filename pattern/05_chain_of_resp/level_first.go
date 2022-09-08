package main

import "fmt"

type levelFirst struct {
	next examination
}

func (i *levelFirst) execute(l *intern) {
	if l.levelZeroDone != true {
		fmt.Println(l.name, " hasn't complete L0!")
		return
	}
	fmt.Println(l.name, " successfully passed L1!")
	l.levelFirstDone = true
	i.next.execute(l)
}
func (i *levelFirst) setNext(next examination) {
	i.next = next
}
