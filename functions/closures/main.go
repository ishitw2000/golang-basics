package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, createTransformer(2))
	tripled := transformNumbers(&numbers, createTransformer(3))

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// every anonymous function is a closure

// an example of closure
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
