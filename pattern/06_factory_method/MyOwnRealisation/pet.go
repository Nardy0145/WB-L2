package main

type pet struct {
	name    string
	petType string
}

func (p *pet) getName() string { return p.name }

func (p *pet) getType() string { return p.petType }
