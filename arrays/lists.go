package main

import "fmt"

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[1])
// 	prices[1] = 11.99

// 	prices = append(prices, 15) // to append new element

// 	prices = prices[1:] // delete first element
// 	fmt.Println(prices)

// }

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{55.9, 99.9, 0.67, 78.89}
// 	fmt.Println(prices)
// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.9
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(featuredPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

type product struct {
	title string
	id    int
	price float64
}

func New(title string, id int, price float64) (*product, error) {

	// if title == "" || id == nil || price == nil {
	// 	return nil, errors.New("invalid Input")
	// }

	return &product{
		title: title,
		id:    id,
		price: price,
	}, nil
}

func main() {
	usernames := make([]string, 2, 5) // make(array, inital_capacity, maximum_capacity)

	usernames[0] = "Hulie"
	usernames[1] = "Julie"
	usernames = append(usernames, "Max")
	usernames = append(usernames, "Manuel")

	fmt.Println(usernames)

	for index, value := range usernames {
		fmt.Printf("%d %v\n", index, value)
	}
}

// func main() {

// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[0:1])
// 	prices[1] = 9.99

// 	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
// 	prices = prices[1:]
// 	fmt.Println(prices)

// 	discountPrices := []float64{101.99, 80.99, 20.59}
// 	prices = append(prices, discountPrices...)

// 	fmt.Println(prices)

// }

// func main() {
// 	hobbies := []string{"badminton", "chess", "coding"}
// 	fmt.Println(hobbies)
// 	fmt.Println(hobbies[0])
// 	h2 := hobbies[1:]
// 	h1 := hobbies[:2]
// 	h3 := h1[1:3]
// 	fmt.Println(h2)
// 	fmt.Println(h1)
// 	fmt.Println(h3)
// 	course_goals := []string{"learn go", "learn go2"}
// 	course_goals[1] = "learn chess"
// 	course_goals = append(course_goals, "learn coding")

// 	p1, _ := New("P1", 1, 123.4)
// 	p2, _ := New("P2", 2, 455.3)
// 	products := []*product{p1, p2}

// 	fmt.Println(products[0])

// }

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
