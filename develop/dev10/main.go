package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go run .\telnet.go --timeout=10s 127.0.0.1 8000

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		fmt.Println("Waiting for connection...")
		conn, err := l.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error accepting: %v\n", err)
			return
		}
		fmt.Println("New connection. Waiting for messages...")
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	for {
		defer conn.Close()
		fmt.Printf("Serving new conn %v\n", conn)
		connReader := bufio.NewReader(conn)
		for {
			message, err := connReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Printf("Connection %v closed.\n", conn)
					return
				}
				fmt.Fprintf(os.Stderr, "error reading from conn: %v\n", err)
				return
			}
			message = strings.TrimSpace(message)

			fmt.Printf("From: %v Received: %s\n", conn, string(message))

			newmessage := strings.ToLower(message)
			_, err = conn.Write([]byte(newmessage + "\n"))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error writing to conn: %v\n", err)
				break
			}
		}
		fmt.Printf("Done serving client %v\n", conn)
	}

}
