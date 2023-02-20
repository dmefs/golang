package graph_test

import (
	"testing"

	"gopl.io/ch4/graph"
)

func TestGraph(t *testing.T) {
	graph.AddEdge("1", "2")
	if !graph.HasEdge("1", "2") {
		t.Fatal("1->2 not in graph")
	}
	graph.AddEdge("2", "3")
	if !graph.HasEdge("2", "3") {
		t.Fatal("2->3 not in graph")
	}
}
