// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}

func comma(s string) string {
	var buf bytes.Buffer
	b := []byte(s)
	if r := len(b) % 3; len(b) != 0 {
		if r == 0 {
			r = 3
		}
		if _, err := buf.Write(b[:r]); err != nil {
			log.Fatal(err)
		}
		b = b[r:]
	}
	sep := ","
	for i := 0; i < len(b); i += 3 {
		buf.WriteString(sep)
		buf.Write(b[i : i+3])
	}
	return buf.String()
}
