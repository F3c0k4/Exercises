package main

func main() {

	c := make(chan int)
	go func() {
		receive(c)
		send(c)
	}()

}

func send(c <-chan int) {
	<-c
}

func receive(c chan<- int) {
	c <- 1
}
