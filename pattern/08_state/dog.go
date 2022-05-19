package main

type dog struct {
	activity state
	tired    state

	currentState state
}

func (d *dog) setState(s state) {
	d.currentState = s
}

func newDog() *dog {
	d := &dog{}
	a := &activity{d}
	t := &tired{d}
	d.activity = a
	d.tired = t
	d.setState(a)
	return d
}

func (d *dog) play() {
	d.currentState.play()
}

func (d *dog) rest() {
	d.currentState.rest()
}
