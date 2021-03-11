package main

import "fmt"

func main() {
	//foo()
	//bar()
	//Exercise 1
	fmt.Println("\n\t Exercise 1\n")
	fmt.Println(foo())
	fmt.Println(bar())

	//Exercise 2
	fmt.Println("\n\t Exercise 2\n")
	fmt.Println(exc2_foo(1, 2, 3))
	fmt.Println(exc2_bar([]int{1, 2, 3}))

	//Exercise 3
	fmt.Println("\n\t Exercise 3\n")
	exc3()

	//Exercise 4
	fmt.Println("\n\t Exercise 4\n")
	var p1 person = person{
		first: "John",
		last:  "Doe",
		age:   34,
	}
	p1.speak()
}

//Exercise 1
func foo() int {
	return 5
}

func bar() (string, int) {
	return "asd", 6
}

//Exercise 1
func exc2_foo(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func exc2_bar(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func exc3() {
	defer fmt.Println("This runs last")
	fmt.Println("This runs first")
	fmt.Println("This runs second")
}

//Exercise 4
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Name: ", p.first, p.last, ", Age: ", p.age)
}
