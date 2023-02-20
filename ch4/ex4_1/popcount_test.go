package popcount_test

import (
	"testing"

	popcount "gopl.io/ch4/ex4_1"
)

func TestCompSha256(t *testing.T) {
	x := [32]byte{1}
	y := [32]byte{2}
	expect := 2
	if solution := popcount.CompSha256(&x, &y); solution != expect {
		t.Fatal(expect, "!=", solution)
	}
	x = [32]byte{0xF, 0xF, 0xF, 0xF}
	y = [32]byte{0, 0, 0, 0, 0xF, 0xF, 0xF, 0xF}
	expect = 32
	if solution := popcount.CompSha256(&x, &y); solution != expect {
		t.Fatal(expect, "!=", solution)
	}
}
