package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age : ", age)
	fmt.Println(getAdultYears(&age))
}

func getAdultYears(age *int) int {
	return *age - 18
}
