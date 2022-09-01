Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Бесконечные нули
Думаю здесь дело в том, что в мердж мы кидаем два закрытых канала
При чтении с закрытого канала мы получаем zero-value, поэтому в горутине с фор мы будем писать нули, когда пройдет итерация 1-8
Отсюда и бесконечный вывод нулей:
1)main горутина, итерирующая по ch
2)merge горутина, бесконечно пишущая в ch
```
