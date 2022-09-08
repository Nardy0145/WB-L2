package main

type shape interface {
	accept(visitor)
}
