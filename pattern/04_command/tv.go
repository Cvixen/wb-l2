package main

import "fmt"

//Конкретный девайс, который релализует методы, а комманды обращаюится
// к этим методам
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
