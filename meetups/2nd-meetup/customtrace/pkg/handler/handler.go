package handler

import (
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/errors"
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/service"
)

type Handler struct {
	Svc service.Service
}

func (h Handler) HandleThing() (res string, err error) {
	op := "handler/Handler.HandleThing"
	if _, err := h.Svc.GetSomething(); err != nil {
		return "", errors.E(op, err)
	}
	return "", err
}
