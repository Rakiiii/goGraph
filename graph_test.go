package graphlib

import (
	"fmt"
	"log"
	"testing"
)

func TestGraphCRS(t *testing.T) {
	var parser Parser
	Graph, err := parser.ParseUnweightedUndirectedGraphCRSFromFile("testGraph1")
	fmt.Println(Graph.AmountOfVertex(), Graph.AmountOfEdges())
	if err != nil {
		log.Println(err)
		return
	}
	Graph.Print()
}
