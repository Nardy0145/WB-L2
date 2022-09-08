package main

import "fmt"

type interview struct {
	next examination
}

func (i *interview) execute(l *intern) {
	fmt.Println(l.name, " successfully passed interview!")
	l.interviewDone = true
	i.next.execute(l)
}
func (i *interview) setNext(next examination) {
	i.next = next
}
