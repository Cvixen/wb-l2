package main

//Это интрефейс со старой реализацией. Чтобы добавить новую, мы просто используем ассепт
type shape interface {
	getType() string
	accept(visitor)
}
