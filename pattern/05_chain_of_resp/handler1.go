package main

//Содержит в себе другой обработчик запроса, если этот обработчик не подзодит
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message) //передача следующему обработчику
	}
	return
}
