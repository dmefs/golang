package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gettext: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
		return
	} else if n.Type == html.TextNode {
		fmt.Println(strings.Trim(n.Data, " \n"))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
