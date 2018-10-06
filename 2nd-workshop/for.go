package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello %d \n", i)
	}

	curr := 0
	for curr <= 10 {
		if curr == 7 || curr == 3 {
			fmt.Println("Yay anda beruntung")
		} else {
			fmt.Printf("Curr %d \n", curr)
		}
		curr++
	}
}
