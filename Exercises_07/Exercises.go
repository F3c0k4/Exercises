package main

import "fmt"

func main() {
	myvar := 67
	fmt.Println(&myvar)

	p1 := person{
		first: "Jon",
		last:  "Doe",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "Jane"
}
