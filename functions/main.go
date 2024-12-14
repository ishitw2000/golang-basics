package main

import "fmt"

type transformFun func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := tranformNumbers(&numbers, getTransformerFunction(&numbers))
	numbers = numbers[1:]
	tripled := tranformNumbers(&numbers, getTransformerFunction(&numbers))
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func tranformNumbers(numbers *[]int, transform transformFun) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
