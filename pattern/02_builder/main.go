package main

import "fmt"

func main() {
	defPhoneBuilder := getBuilder("Default")
	//defPhone.setMemory()
	//defPhone.setRam()
	//defPhone.setCameraMp()
	//defPhone.setDisplayType()
	//phoneObj := defPhone.getSmartphone()
	director := newDirector(defPhoneBuilder)
	phoneObj := director.buildPhone()
	fmt.Printf("\nDefault phone chars:\nram: %d\ncamera mp: %d\nmemory: %d\ndisplay type: %s\n\n",
		phoneObj.ram,
		phoneObj.cameraMp,
		phoneObj.memory,
		phoneObj.displayType)
	shitPhoneBuilder := getBuilder("Shit")
	director.setBuilder(shitPhoneBuilder)
	phoneObj = director.buildPhone()
	//shitPhone.setRam()
	//shitPhone.setMemory()
	//shitPhone.setCameraMp()
	//shitPhone.setDisplayType()
	//phoneObj = shitPhone.getSmartphone()
	fmt.Printf("\nShit phone chars:\nram: %d\ncamera mp: %d\nmemory: %d\ndisplay type: %s\n\n",
		phoneObj.ram,
		phoneObj.cameraMp,
		phoneObj.memory,
		phoneObj.displayType)
}

// Паттерн "Строитель".
// Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

/*
 	Позволяет создавать продукты пошагово.
+	Позволяет использовать один и тот же код для создания различных продуктов.
 	Изолирует сложный код сборки продукта от его основной бизнес-логики.
*/
/*
-	Усложняет код программы из-за введения дополнительных классов.
-	Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/
