package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 88, -45)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// functions running with any amount of parameters
func sumup(startingVale int, numbers ...int) int {
	sum := 0
	fmt.Println("Starting val = ", startingVale)
	fmt.Println("Slice of numbers = ", numbers)
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// func sumup(numbers []int) int {
// 	sum := 0

// 	for _, val := range numbers {
// 		sum += val
// 	}

// 	return sum
// }
