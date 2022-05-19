package main

//Реализует интерфес комманды. Содержат девайс(телик,телефон)
//Команды в виде отдельных объектов реализуются и инкапсулирует методы в самом телевизоре
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}
