package main

import (
	"fmt"
)

func getPet(petType string) (IPet, error) {
	if petType == "Cat" {
		return newCat(), nil
	}
	if petType == "Dog" {
		return newDog(), nil
	}
	return nil, fmt.Errorf("no such type")
}
