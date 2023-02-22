package sum_test

import (
	"testing"

	"gopl.io/ch5/sum"
)

func TestSum(t *testing.T) {
	values := []int{1, 2, 3, 4}
	if sum.Sum(values...) != 10 {
		t.Fatal()
	}
}
