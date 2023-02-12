package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			fmt.Println(line)
		}
	}

	fmt.Println(len(seen))
	if err := input.Err(); err != nil {
		panic(err)
	}
}
