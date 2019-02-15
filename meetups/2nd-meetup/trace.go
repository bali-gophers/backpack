package main

import (
	"fmt"
	"log"

	"github.com/juju/errgo"
)

type TraceError struct {
	Code    string
	Message string
}

func (h TraceError) Error() string {
	return fmt.Sprintf("%s: %s", h.Code, h.Message)
}

func doSomethingForTrace() error {
	return errgo.Mask(TraceError{"AlreadyExists", "item is already exists"})
}

func doSomethingForTrace1() error {
	return errgo.Mask(TraceError{"EntityNotFound", "item couldn't be found"})
}

func logErr(err error) {
	if e, ok := err.(errgo.Locationer); ok {
		log.Printf("%s - %s", e.Location(), err.Error())
	} else {
		log.Println(err.Error())
	}
}

func main() {
	err := doSomethingForTrace()
	if err != nil {
		logErr(err)
	}
	err = doSomethingForTrace1()
	if err != nil {
		logErr(err)
	}
}
