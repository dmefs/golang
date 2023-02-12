package main

import "fmt"

func main() {
	seen := map[int]struct{}{}
	l := []int{1, 2, 3, 4, 5}
	for _, k := range l {
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			fmt.Println(k)
		}
	}
}
