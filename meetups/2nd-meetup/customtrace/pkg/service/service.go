package service

import (
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/errors"
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/repo"
)

type Service struct {
	Repo repo.Repo
}

func (s Service) GetSomething() (res string, err error) {
	var op = "service/Service.GetSomething"
	res, err = s.Repo.GetSomething()
	if err != nil {
		return "", errors.E(op, err)
	}
	return res, nil
}
