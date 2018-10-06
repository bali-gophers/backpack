package main

import (
	"errors"
	"fmt"
)

type CreateOrder struct {
	Email    string
	SKU      string
	Quantity int
}

type Order struct {
	ID        int
	OrderNo   string
	Email     string
	SKU       string
	Quantity  int
	CreatedAt time.Time
}

func funcWithDefer() {
	defer func() {
		fmt.Println("ditulis oleh defer")
	}()
	fmt.Println("Hello World")
}

func funcWithPanic() {
	fmt.Println("Siap- siap panicking")
	panic(errors.New("Pesan saat panic"))
}

func funcWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Siap- siap panicking")
	panic(errors.New("Pesan saat panic"))
}
