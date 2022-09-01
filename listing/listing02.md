Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```
~~~
Ответ:
```
В test() у нас return идёт с именованной переменной

В anotherTest() же у нас return идёт any int, соответственно нужно указать возвращаемое.
Отсюда вытекает предварительная инициализация переменной, присваивание ей значения и последующий его возврат
В общем и целом, под капотом прикол где-то в инициализации
```
