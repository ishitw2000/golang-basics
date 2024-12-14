package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["next"] = 4.3

	// courseRatings.output()

	for key, value := range courseRatings {
		fmt.Printf("%v %v\n", key, value)
	}

}

// func main() {
// 	websites := map[string]string{
// 		"google": "https://google.com",
// 		"AWS":    "https://aws.com",
// 	}
// 	fmt.Println(websites["google"])
// 	websites["google"] = "hello"
// 	fmt.Println(websites["google"])
// 	websites["linkedin"] = " https://linkedin.com"
// 	fmt.Println(websites)

// 	delete(websites, "google")
// 	fmt.Println(websites)

// }
