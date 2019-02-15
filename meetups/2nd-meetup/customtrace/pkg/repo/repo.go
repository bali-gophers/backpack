package repo

import (
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/errors"
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/util"
)

type Repo struct{}

func (r Repo) GetSomething() (res string, err error) {
	var op = "repo/Repo.GetSomething"
	if _, err := util.UtilThing(); err != nil {
		return "", errors.E(op, err)
	}
	return "", errors.Error{"EntityNotFound", "entity couldn't be found", op, nil}
}

func (r Repo) StoreSomething() error {
	var op = "repo/Repo.StoreSomething"
	return errors.Error{"DuplicateItem", "item is duplicate", op, nil}
}
