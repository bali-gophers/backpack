package main

import "fmt"
import "errors"

func main() {
	fmt.Println("Siap- siap panicking")
	panic(errors.New("Pesan saat panic"))
}
