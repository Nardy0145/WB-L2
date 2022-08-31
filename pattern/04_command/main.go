package main

import "fmt"

type Device interface {
	on()
	off()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Radio struct {
	isRunning bool
}

func (r *Radio) on() {
	r.isRunning = true
	fmt.Println("Radio ON")
}
func (r *Radio) off() {
	r.isRunning = false
	fmt.Println("Radio OFF")
}

func main() {
	tv := &Radio{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}
	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}

// Паттерн команда.

// Позволяет не создавать кучу классов при реализации одной и той же инструкции
