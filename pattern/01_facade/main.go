package main

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
Этот шаблон проектирования предназначен для того, чтобы скрыть сложности базовой системы и предоставить клиенту простой интерфейс. Он предоставляет унифицированный интерфейс для множества интерфейсов в системе, что упрощает его использование с точки зрения клиента. По сути, он обеспечивает более высокий уровень абстракции над сложной системой.


*/

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
