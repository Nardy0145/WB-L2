package main

type cat struct {
	pet
}

func newCat() IPet {
	return &pet{
		name:    "Murka",
		petType: "Cat",
	}
}
