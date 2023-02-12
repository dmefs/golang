package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func main() {
	l := []int{1, 2, 3, 4, 5}
	rotate(l, 2)
	fmt.Println(l)
}
