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

	//Exercise 5
	fmt.Println("\n\t Exercise 5\n")
	s1 := square{length: 8}
	c1 := circle{radius: 7}
	info(s1)
	info(c1)

	//Exercise 6
	fmt.Println("\n\t Exercise 6\n")
	func(num int) {
		fmt.Println(num)
	}(2)

	//Exercise 7
	fmt.Println("\n\t Exercise 7\n")
	myfunc := func(str string) {
		fmt.Println(str)
	}
	myfunc("Hello")

	//Exercise 8
	fmt.Println("\n\t Exercise 8\n")
	f := func(i int) func() {
		return func() { fmt.Println(i) }
	}
	g := f(6)

	g()

	//Exercise 9
	fmt.Println("\n\t Exercise 9\n")

	exc9_foo(g)

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

//Exercise 2
func exc2_bar(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

//Exercise 3
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

//Exercise 5
type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return 2 * c.radius * 3.14
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

//Exercise 9
func exc9_foo(myfunc func()) {
	myfunc()
}

//Exercise 10

func exc10_foo() {
	if 2 > 1 {
		str := "2 is bigger than one"
		fmt.Println(str)
	}
}
