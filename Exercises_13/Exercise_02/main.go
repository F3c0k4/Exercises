package main

import (
	"fmt"

	"github.com/F3c0k4/Exercises_01/Exercises_13/Exercise_02/quote"
	"github.com/F3c0k4/Exercises_01/Exercises_13/Exercise_02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
