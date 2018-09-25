package tematik

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	total = Sum(5, -10)
	if total != -5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, -5)
	}
}
