package main

type dog struct {
	pet
}

func newDog() IPet {
	return &pet{
		name:    "Sharik",
		petType: "Dog",
	}
}
