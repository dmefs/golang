package main

import "fmt"

const (
	KB uint64 = 1000
	MB        = 1000 * KB
	GB        = 1000 * MB
	TB        = 1000 * GB
	PB        = 1000 * TB
	EB        = 1000 * PB
	ZB        = 1000 * EB
	YB        = 1000 * ZB
)

const (
	_ = 1 << (10 * iota)

	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)

func main() {
	fmt.Println("KB = ", KB)
	fmt.Println("MB = ", MB)
	fmt.Println("MB = ", GB)
	fmt.Println("MB = ", TB)
	fmt.Println("MB = ", PB)
	fmt.Println("MB = ", EB)

	fmt.Println("KiB = ", KiB)
	fmt.Println("MiB = ", MiB)
	fmt.Println("GiB = ", GiB)
	fmt.Println("TiB = ", TiB)
	fmt.Println("PiB = ", PiB)
	fmt.Println("EiB = ", EiB)
}
