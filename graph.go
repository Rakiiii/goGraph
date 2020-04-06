package graphlib

import (
	"fmt"
	gosort "github.com/Rakiiii/goSort"
)
//Graph is struct that describe graph with vertex edges popose in map
type Graph struct {
	graph          [][]int
	amountOfVertex int
	amountOfEdges  int
}

//Init is initialize graph with nonnil map
func (g *Graph) Init(v, e int) {
	g.amountOfVertex = v
	g.amountOfEdges = e
	g.graph = make([][]int, v)
}

//AmountOfVertex returns amount of graph edges
func (g *Graph) AmountOfVertex() int {
	return g.amountOfVertex
}

//AmountOfEdges returns amount of edges in graph
func (g *Graph) AmountOfEdges() int {
	return g.amountOfEdges
}

//AddVertex function is adding new vertex without edges t graph
func (g *Graph) AddVertex() {

	g.graph = append(g.graph, make([]int, 1))
	g.amountOfVertex++

}

//AddEdgesToVertex function is setting edges of vertex with num @vertex with edges countained in @edges
func (g *Graph) AddEdgesToVertex(vertex int, edges []int) {
	g.graph[vertex] = edges
	g.amountOfEdges += len(edges)
}

//AddVertexWithEdges function is adding new vertex with edges contained in @edges
func (g *Graph) AddVertexWithEdges(edges []int) {
	g.graph = append(g.graph, edges)
	g.amountOfVertex++
	g.amountOfEdges += len(edges)
}

//GetEdges returning edges of vertex number @ver
func (g *Graph) GetEdges(ver int) []int {
	edges := make([]int,len(g.graph[ver]))
	copy(edges,g.graph[ver])
	return edges
}

//Size return amount of added vertex
func (g *Graph) Size() int {
	return len(g.graph) - 1
}

func (g *Graph) ConvertToCRS() *GraphCRS {
	var newgraph GraphCRS
	for _, row := range g.graph {
		newgraph.AddVertexWithEdges(row)
	}
	return &newgraph
}

//Print printing graph as struct of addiction to stdout
func (g *Graph) Print() {
	for _, row := range g.graph {
		for _, elem := range row {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
}

//todo:: add renumbering edges at grpah
//RenumVertex renumbered vertexes in grpah as in @neworder where @graph[i] = @graph[@neworder[i]]
func (g *Graph) RenumVertex(neworder []int) {
	newSet := make([][]int, len(g.graph))

	optPointers := make([]int,g.AmountOfVertex())
	for i,num := range neworder{
		optPointers[num] = i
	}
	
	for i, num := range neworder {
		newSet[i] = make([]int,len(g.graph[num]))
		for j,vertex := range g.graph[num]{
			newSet[i][j] = optPointers[vertex] 
		}
		newSet[i] = gosort.QuicksortInt(newSet[i])
	}

	g.graph = newSet
}
