package main

import "fmt"

func main() {

	c := make(chan int)
	go receive(c)
	for v := range c {
		fmt.Println(v)
	}

}

func send(c <-chan int) {
	<-c
}

func receive(c chan<- int) {
	for i := 0; i < 20; i++ {
		c <- 1
	}
	close(c)

}
