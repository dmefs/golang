package popcount

import "sync"

var initpcOnce sync.Once

// pc[i] is the population count of i.
var pc [256]byte

func initPc() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func CompSha256(x, y *[32]byte) int {
	var z [32]byte
	for i := 0; i < len(z); i++ {
		z[i] = x[i] ^ y[i]
	}
	return PopCount(&z)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x *[32]byte) int {
	initpcOnce.Do(initPc)
	res := 0
	for i := 0; i < 32; i++ {
		res += int(pc[x[i]])
	}
	return res
}
