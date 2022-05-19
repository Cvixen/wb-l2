package main

import "fmt"

type activity struct {
	dog *dog
}

func (a *activity) play() {
	fmt.Println("I want play")
	a.dog.setState(a.dog.tired)
}

func (a *activity) rest() {
	fmt.Println("I dont want rest")
}
