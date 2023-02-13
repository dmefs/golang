package tempconv_test

import (
	"testing"

	tempconv "gopl.io/ch2/tempconv0"
)

func Test(t *testing.T) {
	if f := tempconv.FToC(32); tempconv.FreezingC != f {
		t.Fatalf("FreezingC != %v", f)
	}
}
