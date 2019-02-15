package main

import (
	"log"

	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/handler"
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/repo"
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/service"
)

func main() {
	repo := repo.Repo{}
	svc := service.Service{repo}
	handler := handler.Handler{Svc: svc}
	if _, err := handler.HandleThing(); err != nil {
		log.Println(err.Error())
	}
}
