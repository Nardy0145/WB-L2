package main

import "fmt"

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
	a.area = s.side * 4
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("Calculating ~ area for circle")
	a.area = c.radius * 2 * 3
}

func (a *areaCalculator) visitForTriangle(t *triangle) {
	fmt.Println("Calculating area for triangle")
	a.area = t.border + t.border + t.lower
}
