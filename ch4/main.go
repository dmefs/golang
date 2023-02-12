package main

import "fmt"

type Employee struct {
	Id     int
	Salary int
}

func main() {
	var m1 []int
	m2 := []int{}

	fmt.Println(m1 == m2)
}
