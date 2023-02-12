package main

import "fmt"

func removeDuplicate(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}
	index := 0
	for i, s := range strings {
		if strings[index] != s && (index+1) != i {
			index++
			strings[index] = s
		}
	}
	return strings[:index+1]
}

func main() {
	strings := []string{"hi", "hi"}
	fmt.Println(strings)
	strings = removeDuplicate(strings)
	fmt.Println(strings)
}
