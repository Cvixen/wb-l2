package main

type Handler interface {
	SendRequest(message int) string
}
