package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
Он позволяет создать цепочку обработчиков запросов. Для каждого входящего запроса он проходит через цепочку и каждый из обработчиков:

Обрабатывает запрос или пропускает обработку.
Решает, передать ли запрос следующему обработчику в цепочке или нет.
*/

func main() {
	handler3 := ConcreteHandlerC{next: nil}
	handler2 := ConcreteHandlerB{next: &handler3}
	handler1 := ConcreteHandlerA{next: &handler2}
	//Везде отправляем первому обработчику сообщение, но тем не менее получаем ответ из других обработчиков
	fmt.Println(handler1.SendRequest(1))
	fmt.Println(handler1.SendRequest(2))
	fmt.Println(handler1.SendRequest(3))
	fmt.Println(handler1.SendRequest(4))
}
