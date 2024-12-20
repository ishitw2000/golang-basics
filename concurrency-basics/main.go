package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan)

}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)

	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)

	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ...  are ... you...?", done)
	go greet("I hope you're liking the course!", done)
	// for _, done := range done {
	// 	<-done
	// }
	for range done {
		// fmt.Println(doneChan)
	}
}
