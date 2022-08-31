package main

type DefaultBuilder struct {
	ram         int
	memory      int
	cameraMp    int
	displayType string
}

func NewDefaultBuilder() *DefaultBuilder {
	return &DefaultBuilder{}
}

func (b *DefaultBuilder) setRam() {
	b.ram = 4
}

func (b *DefaultBuilder) setMemory() {
	b.memory = 64
}

func (b *DefaultBuilder) setCameraMp() {
	b.cameraMp = 12
}

func (b *DefaultBuilder) setDisplayType() {
	b.displayType = "Retina"
}

func (b *DefaultBuilder) getSmartphone() Smartphone {
	return Smartphone{
		ram:         b.ram,
		memory:      b.memory,
		cameraMp:    b.cameraMp,
		displayType: b.displayType,
	}
}
