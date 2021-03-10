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

	//Exercise 6
	fmt.Println("\n\tExercise 6\n")

	p := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	for i := 0; i < len(p); i++ {
		fmt.Println(i, p[i])
	}

	fmt.Printf("Length: %v, Capacity: %v\n", len(p), cap(p))

	//Exercise 7
	fmt.Println("\n\tExercise 7\n")

	sli1 := []string{"James", "Bond", "Shaken, not stirred"}
	sli2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	sli3 := [][]string{sli1, sli2}
	for i, sli_2d := range sli3 {
		for j, v := range sli_2d {
			fmt.Println("Idx1: ", i, " Idx2: ", j, " Value: ", v)
		}
	}

	//Exercise 8
	fmt.Println("\n\tExercise 8")

	mymap := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range mymap {
		for i, str := range v {
			fmt.Println("Key in map: ", k, " Index in slice: ", i, " Value in slice: ", str)
		}
	}
}
