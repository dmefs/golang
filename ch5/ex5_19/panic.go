package main

import "fmt"

func main() {
	fmt.Println(hello())
}

func hello() (s string) {
	defer func() {
		switch p := recover(); p {
		case 0:
			s = "hello world"
		default:
			panic("err")
		}
	}()
	panic(0)
}
