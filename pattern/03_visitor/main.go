package main

func main() {
	square := Square{side: 5}
	circle := Circle{radius: 5}
	rectangle := Rectangle{
		l: 10,
		b: 5,
	}
	areaCalculator := &AreaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)
}

// Паттерн "Посетитель"
// Посетитель позволяет вам добавлять поведение в структуру без ее изменения.
