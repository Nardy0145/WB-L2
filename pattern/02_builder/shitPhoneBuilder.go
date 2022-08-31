package main

type ShitBuilder struct {
	ram         int
	memory      int
	cameraMp    int
	displayType string
}

func NewShitBuilder() *ShitBuilder {
	return &ShitBuilder{}
}

func (b *ShitBuilder) setRam() {
	b.ram = 1
}

func (b *ShitBuilder) setMemory() {
	b.memory = 8
}

func (b *ShitBuilder) setCameraMp() {
	b.cameraMp = 2
}

func (b *ShitBuilder) setDisplayType() {
	b.displayType = "Lamp"
}

func (b *ShitBuilder) getSmartphone() Smartphone {
	return Smartphone{
		ram:         b.ram,
		memory:      b.memory,
		cameraMp:    b.cameraMp,
		displayType: b.displayType,
	}
}
