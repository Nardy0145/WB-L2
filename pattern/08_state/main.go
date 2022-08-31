package main

import "fmt"

type something struct {
	value int
	state states
}

type states struct {
	first  bool
	second bool
	third  bool
	fourth bool
}

func doOperationsBasedOnState(whatever something) {
	if whatever.state.first {
		fmt.Println("job based on first state")
	}
	if whatever.state.second {
		fmt.Println("job based on first state")
	}
	if whatever.state.third {
		fmt.Println("job based on third state")
	}
	if whatever.state.fourth {
		fmt.Println("job based on fourth state")
	}
}

func main() {
	something := something{
		value: 500,
		state: states{first: true},
	}
	doOperationsBasedOnState(something)
	something.state.first = false
	something.state.third = true
	doOperationsBasedOnState(something)
}

// Состояние
