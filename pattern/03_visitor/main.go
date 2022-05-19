package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
Шаблон проектирования, который позволяет вам добавлять поведение к структуре, фактически не изменяя структуру.
*/
func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
