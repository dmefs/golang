package removedup_test

import (
	"testing"

	removedup "gopl.io/ch4/ex4_5"
)

func TestDup(t *testing.T) {
	s := []string{"a", "a", "b", "b", "b", "c"}
	removedup.RemoveDup(s)
	expect := []string{"a", "b", "c"}
	for i, v := range expect {
		if expect[i] != v {
			t.Fatalf("expect[%d] = %s, s[%d] = %s\n", i, expect[i], i, v)
		}
	}
}
