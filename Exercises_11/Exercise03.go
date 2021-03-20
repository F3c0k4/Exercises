package main

import "fmt"

type customErr struct {
}

func (ce customErr) Error() string {
	return "customErr error"
}

func foo(err error) string {
	return err.Error()
}

func main() {
	var ce customErr
	fmt.Println(foo(ce))
}
