package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Println("Gun: ", g.getName())
	fmt.Println("Power: ", g.getPower())
}
