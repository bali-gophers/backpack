package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("ditulis oleh defer")
	}()
	fmt.Println("Hello World")
}
