package main

import "fmt"

func main() {
	//Exercise 1
	fmt.Println("\n\tExercise 1\n")
	arr := [5]int{4, 6, 2, 35, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)

	//Exercise 2
	fmt.Println("\n\tExercise 2\n")
	sli := []int{54, 332, 654, 327, 756, 2345, 234234, 5445, 34, 23}

	for _, v := range sli {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", sli)

	//Exercise 3
	fmt.Println("\n\tExercise 3\n")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := x[0:5]
	b := x[5:]
	c := x[2:7]
	d := x[1:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//Exercise 4
	fmt.Println("\n\tExercise 4\n")

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = append(y, 52)

	fmt.Println(y)

	y = append(y, 53, 54, 55)
	fmt.Println(y)
	z := []int{56, 57, 58, 59, 60}
	y = append(y, z...)
	fmt.Println(y)

	//Exercise 5
	fmt.Println("\n\tExercise 5\n")

	u := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	o := u[0:3]
	o = append(o, u[6:]...)

	fmt.Println(o)
}
