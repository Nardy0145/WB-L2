package main

type examination interface {
	execute(*intern)
	setNext(examination)
}
