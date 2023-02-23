package ex515_test

import (
	"testing"

	ex515 "gopl.io/ch5/ex5_15"
)

func TestMin(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	if ex515.Min(7, values...) != 1 {
		t.Fatal()
	}
}

func TestMax(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	if ex515.Max(7, values...) != 7 {
		t.Fatal()
	}
}
