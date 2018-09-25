package main

import "fmt"
import "errors"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Siap- siap panicking")
	panic(errors.New("Pesan saat panic"))
}
