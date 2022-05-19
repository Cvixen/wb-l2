//Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.
//
//```go
package main

import (
	"fmt"
)

func test1() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test1())
	fmt.Println(anotherTest())
}

//```
//
//Ответ:
//2
//1
//...
//
//```
