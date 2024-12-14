package main

import "fmt"

func main() {

	result := add(4, 5)

	fmt.Print(result)

}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
