package main

type director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *director {
	return &director{builder: b}
}

func (d *director) setBuilder(b IBuilder) {
	d.builder = b
}
func (d *director) buildPhone() Smartphone {
	d.builder.setMemory()
	d.builder.setDisplayType()
	d.builder.setCameraMp()
	d.builder.setRam()
	return d.builder.getSmartphone()
}
