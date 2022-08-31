package main

import (
	"fmt"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	area := s.side
	fmt.Printf("Calc area operations for square: %d\n", area)
}
func (a *AreaCalculator) visitForCircle(c *Circle) {
	area := c.radius * 2
	fmt.Printf("Calc diameter operations for circle: %d\n", area)
}
func (a *AreaCalculator) visitForRectangle(t *Rectangle) {
	area := t.l + t.b + t.b
	fmt.Printf("Calc area operations for rectangle: %d\n", area)
}
