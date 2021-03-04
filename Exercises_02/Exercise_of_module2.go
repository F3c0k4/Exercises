package main

import "fmt"

const (
	a     = 4
	b int = 82
)

func main() {
	//Exc 1
	var num int = 18
	fmt.Printf("decimal: %v, binary: %b, hexadecimal: %X\n", num, num, num)

	//Exc2
	fmt.Println(1 == 1)
	fmt.Println(2 >= 3)
	fmt.Println(3 <= 2)
	fmt.Println(3 != 4)
	fmt.Println(5 < 6)
	fmt.Println(6 > 7)

	//Exc3
	fmt.Println(a, b)

	//Exc4
	c := 56
	fmt.Printf("decimal: %v,  binary:  %b, hexadecimal: %X\n", c, c, c)
	d := c << 1
	fmt.Printf("decimal: %v, binary: %b, hexadecimal: %X\n", d, d, d)
	//Exc5
	e := `it is a string as you can see "string in string" here it is`
	fmt.Println(e)

	//Exc6
	const (
		f = iota + 2022
		g = iota + 2022
		h = iota + 2022
		i = iota + 2022
	)

	fmt.Println(f, g, h, i)
}
