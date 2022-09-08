package main

type State interface {
	publish() error
}
