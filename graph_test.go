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

func TestRenumVertex(t *testing.T){
	fmt.Println("Start TestRenumVertex")
	var parser Parser
	graph1, err := parser.ParseUnweightedUndirectedGraphFromFile("testgraphoo")
	if err != nil {
		log.Println(err)
		return
	}

	graph2, err := parser.ParseUnweightedUndirectedGraphFromFile("testgraphno")
	if err != nil {
		log.Println(err)
		return
	}
	rightorder := []int{2,3,0,1,5,6,7,4}
	graph1.RenumVertex(rightorder)

	checkFlag := true
	for i := 0 ; i < graph1.AmountOfVertex();i++{
		for j,e := range graph1.GetEdges(i){
			if e != graph2.GetEdges(i)[j]{
				t.Error("Wrong edge:",e," expected:",graph2.GetEdges(i)[j])
				checkFlag = false
			}
		}
	}
	if checkFlag{
		fmt.Println("TestRenumVertex=[ok]")
	}
}
