package main

import (
	"fmt"
)

type CustomError struct {
	Code    string
	Message string
}

func (h CustomError) Error() string {
	return fmt.Sprintf("%d: %s", h.Code, h.Message)
}

func doSomething() error {
	return CustomError{"AlreadyExists", "item is already exists"}
}

func main() {
	err := doSomething()
	if e, ok := err.(CustomError); ok {
		fmt.Println(e.Message)
		// handle custom error
	} else {
		// handle a different error
	}
}
