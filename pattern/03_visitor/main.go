package main

import "fmt"

func main() {
	square := &square{side: 5}
	circle := &circle{radius: 3}
	triangle := &triangle{
		border: 3,
		lower:  5,
	}
	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	fmt.Println(areaCalculator.area)
	circle.accept(areaCalculator)
	fmt.Println(areaCalculator.area)
	triangle.accept(areaCalculator)
	fmt.Println(areaCalculator.area)
}

// Паттерн "Посетитель"
// Посетитель позволяет вам добавлять поведение в структуру без ее изменения.

/*
	Упрощает добавление операций, работающих со сложными структурами объектов.
+	Объединяет родственные операции в одном классе.
	Посетитель может накапливать состояние при обходе структуры элементов.
*/

/*
-	Паттерн не оправдан, если иерархия элементов часто меняется.
-	Может привести к нарушению инкапсуляции элементов.
*/