package main

import (
	"errors"
	"log"

	"github.com/juju/errgo"
)

func handleSomething() error {
	return errgo.Mask(errors.New("item is nil"))
}

func main() {
	err := handleSomething()
	if err != nil {
		if e, ok := err.(errgo.Locationer); ok {
			log.Printf("%s - %s", e.Location(), err.Error())
		} else {
			log.Println(err.Error())
		}
	}
}
