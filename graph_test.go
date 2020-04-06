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

func TestGetEdges(t *testing.T){
	fmt.Println("Start TestGetEdges")
	var parser Parser
	graph1, err := parser.ParseUnweightedUndirectedGraphFromFile("testgraphoo")
	if err != nil {
		log.Println(err)
		return
	}
	checkFlag := true

	if graph1.GetEdges(2)[0] != 5{
		t.Error("Wrong edge at pos:2,0 expected 5 found",graph1.GetEdges(2)[0])
		checkFlag = false
	}
	if graph1.GetEdges(7)[1] != 4{
		t.Error("Wrong edge at pos:7,1 expected 4 found",graph1.GetEdges(7)[1])
		checkFlag = false
	}
	if graph1.GetEdges(5)[0] != 2{
		t.Error("Wrong edge at pos:5,0 expected 2 found",graph1.GetEdges(5)[0])
		checkFlag = false
	}
	edg := graph1.GetEdges(2)
	edg[0] = 12

	if graph1.GetEdges(2)[0] != 5{
		t.Error("retest Wrong edge at pos:2,0 expected 5 found",graph1.GetEdges(2)[0])
		checkFlag = false
	}
	if checkFlag{
		fmt.Println("TestGetEdges=[ok]")
	}
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
