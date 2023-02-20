package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/sha3"
)

var algo *string = flag.String("sha", "", "sha algo")

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		switch *algo {
		case "sha384":
			fmt.Printf("%x\n", sha3.Sum384([]byte(arg)))
		case "sha512":
			fmt.Printf("%x\n", sha3.Sum512([]byte(arg)))
		default:
			fmt.Printf("%x\n", sha3.Sum256([]byte(arg)))
		}
	}
}
