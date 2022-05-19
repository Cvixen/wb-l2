package main

//это инвокер, то есть управляет коммандами
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
