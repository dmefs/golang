package comma_test

import (
	"testing"

	comma "gopl.io/ch3/ex3_11"
)

func TestDecimal(t *testing.T) {
	check(t, "1", "1")
	check(t, "123", "123")
	check(t, "1234", "1,234")
	check(t, "123456", "123,456")
}

func TestFloatingPoint(t *testing.T) {
	check(t, "1.", "1.")
	check(t, "123.45", "123.45")
	check(t, "1234.45", "1,234.45")
	check(t, "123456.4567", "123,456.4567")
}

func TestSign(t *testing.T) {
	check(t, "+1.", "+1.")
	check(t, "-123.45", "-123.45")
	check(t, "+1234.45", "+1,234.45")
	check(t, "-123456.4567", "-123,456.4567")
}

func check(t *testing.T, origin string, expect string) {
	if result := comma.Comma(origin); result != expect {
		t.Fatal(result, "!=", expect)
	}
}
