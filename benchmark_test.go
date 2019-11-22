package graphlib

import (
	"log"
	"testing"
)

var grpahname = "graphBig"

func BenchmarkFullAccsesGraphCRS(b *testing.B) {
	var parser Parser
	Graph, err := parser.ParseUnweightedUndirectedGraphCRSFromFile(grpahname)
	if err != nil {
		log.Println(err)
		return
	}
	for j := 0; j < b.N; j++ {
		counter := -32345
		for i := 0; i < Graph.AmountOfVertex()-1; i++ {
			for _, edge := range Graph.GetEdges(i) {
				counter += edge
			}
		}
	}
}

func BenchmarkFullAccsesGraph(b *testing.B) {
	var parser Parser
	Graph, err := parser.ParseUnweightedUndirectedGraphCRSFromFile(grpahname)
	if err != nil {
		log.Println(err)
		return
	}
	for j := 0; j < b.N; j++ {
		counter := -32345
		for i := 0; i < Graph.AmountOfVertex()-1; i++ {
			for _, edge := range Graph.GetEdges(i) {
				counter += edge
			}
		}
	}
}
