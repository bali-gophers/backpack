package util

import (
	"github.com/bali-gophers/backpack/meetups/2nd-meetup/customtrace/pkg/errors"
)

func UtilThing() (string, error) {
	var op = "util/UtilThing"
	return "", errors.Error{"SomethingBad", "something bad", op, nil}
}
