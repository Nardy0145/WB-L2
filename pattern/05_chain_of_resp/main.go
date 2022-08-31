package main

import "fmt"

type WildBerries interface {
	execute(intern *Intern)
	setNext(berries WildBerries)
}

type L0 struct {
	next WildBerries
}
type L1 struct {
	next WildBerries
}
type L2 struct {
	next WildBerries
}

func (l *L0) execute(i *Intern) {
	fmt.Println("L0 done")
	i.L0Done = true
	l.next.execute(i)
}
func (l *L0) setNext(next WildBerries) {
	l.next = next
}
func (l *L1) execute(i *Intern) {
	fmt.Println("1 done")
	i.L1Done = true
	l.next.execute(i)
}
func (l *L1) setNext(next WildBerries) {
	l.next = next
}
func (l *L2) execute(i *Intern) {
	fmt.Println("L2 done")
	i.L2Done = true
	l.next.execute(i)
}
func (l *L2) setNext(next WildBerries) {
	l.next = next
}

type Intern struct {
	L0Done bool
	L1Done bool
	L2Done bool
}

func main() {
	l2 := &L2{}
	l1 := &L1{}
	l1.setNext(l2)
	l0 := &L0{}
	l0.setNext(l1)
	me := Intern{}
	l0.execute(&me)
}

// Цепочка обязанностей
