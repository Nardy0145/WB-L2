package main

import "fmt"

const (
	Cat string = "Cat"
	Dog string = "Dog"
)

func GetPet(petType string) {
	pet, _ := getPet(petType)
	getInfo(pet)
}

func getInfo(p IPet) {
	fmt.Println("Pet name: ", p.getName())
	fmt.Println("Type: ", p.getType())
}
func main() {
	GetPet(Cat)
	GetPet(Dog)
}
