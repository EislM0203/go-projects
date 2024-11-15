package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true
	close(done)
}

func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool)

	//dones[0] =  make(chan bool)
	go greet("Nice to meet you!", done)
	//dones[1] =  make(chan bool)
	go greet("How are you?", done)
	//dones[2] =  make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	//dones[3] =  make(chan bool)
	go greet("Go rocks!", done)

	//for _, done := range dones {
	//	<- done
	//}
	for range done {}
}
