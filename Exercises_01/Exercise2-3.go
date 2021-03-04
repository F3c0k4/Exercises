package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x, " ", y, " ", z)

	//Exercise 3
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Printf("%v", s)

}
