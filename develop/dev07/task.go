package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“done after %v”, time.Since(start))
*/

// Or ../
func Or(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})

	// Итерируя по каналам, создаем для каждого горутину, проверяющую канал и закрывающую горутину
	for _, channel := range channels {
		go func(ch <-chan interface{}) {
			select {
			// Если хотя-бы с одного канала можно прочитать (var sig func), пишем в done и закрываем доступ остальным рутинам.
			// В этот момент тикает <-done в конце функции
			case <-ch:
				done <- "stop"
				close(done)
			}
		}(channel)
	}
	// Здесь. И поток main оффается
	<-done
	return done
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-Or(
		sig(4*time.Second),
		sig(5*time.Second),
		sig(15*time.Second),
		sig(7*time.Second),
		sig(12*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
