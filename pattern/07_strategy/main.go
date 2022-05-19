package main

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
*/

func main() {
	digits := []int{2, 1, 6, 4, 3, 0}
	bubble := BubbleSort{}
	insert := InsertionSort{}

	context := Context{strategy: &bubble}
	context.Sort(digits)
	context.Algorithm(&insert)
	context.Sort(digits)
}
