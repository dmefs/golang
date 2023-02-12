package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(n *html.Node, mapping map[string]int) {
	if n.Type == html.ElementNode {
		fmt.Println("Add.")
		mapping[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(n, mapping)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	mapping := make(map[string]int)
	visit(doc, mapping)
	for name, count := range mapping {
		fmt.Printf("%s: %d\n", name, count)
	}
}
