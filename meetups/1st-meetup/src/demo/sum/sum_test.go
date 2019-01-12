package sum_test

import (
	"testing"

	"github.com/bali-gophers/backpack/meetups/1st-meetup/src/demo/sum"
)

func TestSum(t *testing.T) {
	expected := 8
	res := sum.Sum(3, 5)
	if res != expected {
		t.Errorf("Expect %d, got %d", expected, res)
	}
}
