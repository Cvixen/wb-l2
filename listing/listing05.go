//Что выведет программа? Объяснить вывод программы.
//
//```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		return &customError{msg: "lol"}
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

//```
//
//Ответ:error
//```
//...
//
//```
