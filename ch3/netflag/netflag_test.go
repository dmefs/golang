package netflag_test

import (
	"testing"

	"gopl.io/ch3/netflag"
)

func Test(t *testing.T) {
	var v netflag.Flags = netflag.FlagMulticast | netflag.FlagUp
	var solution int
	solution = 0b10001
	if v != solution {
		t.Fatal(v, "!=", solution)
	}
	TurnDown(&v)
	solution = 0b10000
	if v != solution {
		t.Fatal(v, "!=", solution)
	}
	SetBroadcast(&v)
	solution = 0b10010
	if v != solution {
		t.Fatal(v, "!=", solution)
	}
}
