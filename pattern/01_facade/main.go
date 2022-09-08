package main

import "fmt"

type Card struct {
	balance int
	code    int
}

type Game struct {
	name string
	cost int
}

func BuyGame(card Card, game Game) {
	if card.balance < game.cost {
		fmt.Println("not enough money")
	} else {
		fmt.Println(" congratulations, you've just bought: ", game.name)
	}
}

func main() {
	myCard := Card{
		balance: 500,
		code:    1234,
	}
	someGame := Game{
		name: "Super Gopher Game",
		cost: 450,
	}
	BuyGame(myCard, someGame)
}

// Паттерн фасад.
// Он позволяет клиенту работать с десятками компонентов, используя при этом простой интерфейс.

// + Изолирует клиентов от компонентов сложной подсистемы.
// - Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
