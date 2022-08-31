package main

type IBuilder interface {
	setRam()
	setMemory()
	setCameraMp()
	setDisplayType()
	getSmartphone() Smartphone
}

func getBuilder(builderType string) IBuilder {
	if builderType == "Default" {
		return NewDefaultBuilder()
	}
	if builderType == "Shit" {
		return NewShitBuilder()
	}
	return nil
}
