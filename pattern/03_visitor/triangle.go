package main

type triangle struct {
	border int
	lower  int
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}
