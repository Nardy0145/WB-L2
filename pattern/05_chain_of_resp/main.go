package main

func main() {
	l2 := &levelSecond{}

	l1 := &levelFirst{}
	l1.setNext(l2)

	l0 := &levelZero{}
	l0.setNext(l1)

	interV := &interview{}
	interV.setNext(l0)

	inter := &intern{name: "Dmitry"}
	interV.execute(inter)
}

// Цепочка обязанностей

/*
  	Уменьшает зависимость между клиентом и обработчиками.
+ 	Реализует принцип единственной обязанности.
  	Реализует принцип открытости/закрытости.
*/

/*
-	Запрос может остаться никем не обработанным.
*/
