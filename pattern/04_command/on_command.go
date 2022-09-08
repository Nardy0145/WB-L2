package main

type onCommand struct {
	device device
}

func (on *onCommand) execute() {
	on.device.on()
}
