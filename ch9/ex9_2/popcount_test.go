package popcount_test

import (
	"fmt"
	"testing"

	popcount "gopl.io/ch9/ex9_2"
)

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}
func TestPopcount(t *testing.T) {
	for i := 1; i < 16; i++ {
		result := popcount.PopCount(uint64(i))
		fmt.Println("[Popcount] i: ", i, "result: ", result)
	}
}
