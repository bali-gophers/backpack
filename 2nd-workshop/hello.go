package main

import (
	"fmt"

	"github.com/bali-gophers/backpack/2nd-workshop/hellopkg"
)

type Query interface {
	Query(q string) (string, error)
	MultiQuery(q string) ([]string, error)
}

type CacheQuery struct{}

func (c CacheQuery) Query(q string) (string, error) {
	return "", nil
}

func (c CacheQuery) MultiQuery(q string) ([]string, error) {
	return []string{}, nil
}

type Hello interface {
	Hello() string
}

type HelloBali struct {
	Name  string
	Value int
}

func (h HelloBali) Hello() string {
	return fmt.Sprintf("Hello %s, %d", h.Name, h.Value)
}

func SayHello(hello Hello) {
	fmt.Println(hello.Hello())
}

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("%d %s", c.Code, c.Message)
}

func executeVal(val int) (int, error) {
	return 0, CustomError{1, "Error"}
}

func main() {
	var isHealth bool
	var emailAddress string

	emailAddress = "raka@gmail.com"
	shippingAddress := "Jl Wr. Supratman"
	fmt.Println(emailAddress)
	fmt.Println(shippingAddress)

	isHealth = true
	firstName := "Raka"
	fmt.Printf("My first name: %s \n", firstName)
	fmt.Printf("isHealth: %s\n", isHealth)

	fmt.Println(hellopkg.Hello("Raka"))
	fmt.Println("Hello World")

	var letters []string
	letters = []string{
		"hello",
		"hi",
		"hallo",
	}

	for _, letter := range letters {
		fmt.Println(letter)
	}

	letterHi := letters[1]
	lettersA := letters[1:]
	fmt.Println(letterHi)

	letters = append(letters, "Morning")
	for _, letter := range letters {
		fmt.Printf("%s\n", letter)
	}

	fmt.Println("-==")
	for _, letter := range lettersA {
		fmt.Printf("%s\n", letter)
	}

	val, err := executeVal(17)
	if err != nil {
		fmt.Printf("fungsi tadi error %s\n", err)
	} else {
		fmt.Printf("nilai %s \n", val)
	}

	h := HelloBali{"bali", 46}
	fmt.Println(h.Hello())
	SayHello(h)

	valString := funcWithDefer()
	fmt.Println(valString)

}

func funcWithDefer() string {
	defer func() {
		fmt.Println("ditulis oleh defer 1")
	}()
	defer func() {
		fmt.Println("ditulis oleh defer 2")
	}()
	if true {
		fmt.Println("Hello true")
		return "true"
	} else {
		fmt.Println("Hello false")
		return "false"
	}
}
