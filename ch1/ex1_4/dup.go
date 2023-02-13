// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		counts["stdin"] = countLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts[arg] = countLines(f)
			f.Close()
		}
	}
	for name, content := range counts {
		for line, n := range content {
			if n > 1 {
				fmt.Printf("%s: %d\t%s\n", name, n, line)
			}
		}
	}
}

func countLines(f *os.File) map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts
	// NOTE: ignoring potential errors from input.Err()
}
