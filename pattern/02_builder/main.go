package main

import "fmt"

func main() {
	defPhone := getBuilder("Default")
	defPhone.setMemory()
	defPhone.setRam()
	defPhone.setCameraMp()
	defPhone.setDisplayType()
	phoneObj := defPhone.getSmartphone()
	fmt.Printf("\nDefault phone chars:\nram: %d\ncamera mp: %d\nmemory: %d\ndisplay type: %s\n\n",
		phoneObj.ram,
		phoneObj.cameraMp,
		phoneObj.memory,
		phoneObj.displayType)
	shitPhone := getBuilder("Shit")
	shitPhone.setRam()
	shitPhone.setMemory()
	shitPhone.setCameraMp()
	shitPhone.setDisplayType()
	phoneObj = shitPhone.getSmartphone()
	fmt.Printf("\nShit phone chars:\nram: %d\ncamera mp: %d\nmemory: %d\ndisplay type: %s\n\n",
		phoneObj.ram,
		phoneObj.cameraMp,
		phoneObj.memory,
		phoneObj.displayType)
}

// Паттерн "Строитель".
// Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
