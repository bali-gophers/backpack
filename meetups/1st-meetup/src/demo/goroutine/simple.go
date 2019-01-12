package main

import (
	"fmt"
	"math/rand"
	"time"
)

func someComputation() int {
	fmt.Println("doing some computation ...")
	time.Sleep(500 * time.Millisecond)
	return rand.Intn(100)
}

func compute(ch chan int) {
	ch <- someComputation()
}

func main() {
	ch1 := make(chan int)
	go compute(ch1)
	result1 := <-ch1
	fmt.Println(result1)
}
