package main

import "fmt"

type tired struct {
	dog *dog
}

func (a *tired) play() {
	fmt.Println("I dont want play")
}

func (a *tired) rest() {
	fmt.Println("I want rest")
	a.dog.setState(a.dog.activity)
}
