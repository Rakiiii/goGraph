package graphlib

import (
	"fmt"
	"log"
	"testing"
)

func TestGetEdgesCRS(t *testing.T){
	fmt.Println("Start TestGetEdgesCrs")
	var parser Parser
	graph1, err := parser.ParseUnweightedUndirectedGraphCRSFromFile("testgraphoo")
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
		fmt.Println("TestGetEdgesCrs=[ok]")
	}
}