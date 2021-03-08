package main

import "fmt"

func main() {
	//Exercise 1
	for i := 1; i <= 10001; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//Exercise 2
	for i := 65; i <= 90; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}

	}

	//Exercise 3

	year := 1994

	for year <= 2021 {
		fmt.Println(year)
		year++
	}

	//Exercise 4

	year = 1994

	for {
		if year == 2022 {
			break
		}
		fmt.Println(year)
		year++
	}

	//Exercise 5

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

	//Exercise 6
	myvar := 5

	if myvar == 5 {
		fmt.Println("myvar equals 5")
	} else if myvar < 5 {
		fmt.Println("myvar is less than 5")
	} else {
		fmt.Println("myvar is greater than 5")
	}

	//Exercise 7

	switch {
	case true:
		fmt.Println("It's true")
	case false:
		fmt.Println("It is not true")
	}
	//Exercise 8
	var favSport string
	favSport = "Football"
	switch favSport {
	case "Football":
		fmt.Println("My favourite sport is football")
	case "Basketball":
		fmt.Println("My favourite sport is basketball")
	}

	/* Exercise 9
	true
	false
	true
	true
	false
	*/
}
