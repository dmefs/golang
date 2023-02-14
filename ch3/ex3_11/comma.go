// comma inserts commas in a non-negative decimal integer string.
package comma

import (
	"bytes"
	"log"
)

func Comma(s string) string {
	var buf bytes.Buffer
	b := []byte(s)
	var floatingIndex int
	// Handle floating point
	if floatingIndex = bytes.Index(b, []byte(".")); floatingIndex != -1 {
		b = b[:floatingIndex]
	}
	// Handle sign
	if b[0] == '+' || b[0] == '-' {
		buf.WriteByte(b[0])
		b = b[1:]
	}

	if len(b) == 0 {
		return s
	}

	r := len(b) % 3
	if r == 0 {
		r = 3
	}
	if _, err := buf.Write(b[:r]); err != nil {
		log.Fatal(err)
	}
	b = b[r:]

	sep := ","
	for i := 0; i < len(b); i += 3 {
		buf.WriteString(sep)
		buf.Write(b[i : i+3])
	}
	if floatingIndex != -1 {
		buf.WriteString(s[floatingIndex:])
	}
	return buf.String()
}
