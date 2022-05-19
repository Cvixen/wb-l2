package main

//Содержит метод для каждого элемента
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}
