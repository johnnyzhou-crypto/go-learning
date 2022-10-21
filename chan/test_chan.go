package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("seed: %v\n", value)
	values <- value

}

func main() {

	defer close(values)
	go send()
	fmt.Printf("wait\n")
	value := <-values
	fmt.Printf("received: %v\n", value)

	fmt.Printf("end\n")

}
