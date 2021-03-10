package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	fav_flavor []string
}

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	//Exercise 1
	fmt.Println("\n\tExercise 1\n")
	p1 := person{
		"Jon",
		"Smith",
		[]string{"Vanilla", "Pistachio", "Cherry"},
	}

	p2 := person{
		"John",
		"Doe",
		[]string{"Chocolate", "Grape", "Lemon"},
	}
	fmt.Println(p1.first_name, p1.last_name)
	for _, v := range p1.fav_flavor {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println(p2.first_name, p2.last_name)
	for _, v := range p2.fav_flavor {
		fmt.Println(v)
	}

	//Exercise 2
	fmt.Println("\n\tExercise 2\n")

	mymap := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, v := range mymap["Smith"].fav_flavor {
		fmt.Println(v)
	}

	for _, v := range mymap["Doe"].fav_flavor {
		fmt.Println(v)
	}

	//Exercise 3
	fmt.Println("\n\tExercise 3\n")

	mytruck := truck{
		vehicle:   vehicle{4, "blue"},
		fourWheel: true,
	}

	mysedan := sedan{
		vehicle: vehicle{5, "white"},
		luxury:  false,
	}

	fmt.Println(mytruck)
	fmt.Println()
	fmt.Println(mysedan)

	fmt.Println(mytruck.vehicle.color)
	fmt.Println(mysedan.vehicle.doors)

	//Exercise 4
	fmt.Println("\n\tExercise 4\n")
	mystruct := struct {
		var1 string
		var2 int
	}{
		var1: "hello",
		var2: 4,
	}

	fmt.Println(mystruct.var1)
}
