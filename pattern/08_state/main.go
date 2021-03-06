package main

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
позволяет объектам менять поведение в зависимости от своего состояния. Извне создаётся впечатление, что изменился класс объекта.
*/

func main() {
	d := newDog()
	//Собака активная, не хочет отдыхать
	d.rest()
	//Активная собака поиграла и теперь состояние уставшее
	d.play()
	//Уставшая собака не хочет больше играть
	d.play()
	//Уставшая собака отдыхает и становится активной
	d.rest()
}
