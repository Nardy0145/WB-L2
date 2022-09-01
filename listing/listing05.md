Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Выведет "error"
Тоже самое, что и в listing03
Как бы да, мы создали кастомный еррор
Опять же, из-за возвращаемого типа, даже если ошибки нет, но структура внутри что-то имеет
еррор не будет равен нил
```
