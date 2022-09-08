package main

type offCommand struct {
	device device
}

func (off *offCommand) execute() {
	off.device.off()
}
