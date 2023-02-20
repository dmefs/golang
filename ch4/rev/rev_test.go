package rev_test

import (
	"testing"

	"gopl.io/ch4/rev"
)

func TestPointerLen(t *testing.T) {
	p := new([32]byte)
	p[0] = 1
	if len(p) != 32 {
		t.Fatal(len(p), "!=", 32)
	}
}

func TestRev(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	rev.Reverse(s)
	expect := []int{5, 4, 3, 2, 1}
	for i := range expect {
		if expect[i] != s[i] {
			t.Fatalf("s[%d] = %d, expect = %d", i, s[i], expect[i])
		}
	}
}
