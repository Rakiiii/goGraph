package graphlib

import "fmt"

//GraphCRS struct to hold graph in CRS struct
type GraphCRS struct {
	vertex []int
	edges  []int
}

//Init initialize graph with @v amount of vertex and @e amount of edges
func (g *GraphCRS) Init(v, e int) {
	g.vertex = make([]int, v)
	g.edges = make([]int, e*2)
}

//AmountOfVertex returns amount of vertex in graph
func (g *GraphCRS) AmountOfVertex() int {
	return len(g.vertex)
}

//AmountOfEdges returns amount of edges in graph
func (g *GraphCRS) AmountOfEdges() int {
	return len(g.edges) / 2
}

//AddVertex function is adding new vertex without edges t graph
func (g *GraphCRS) AddVertex() {
	g.vertex = append(g.vertex, len(g.edges))
	g.edges = append(g.edges, -1)
}

//AddEdgesToVertex function is setting edges of vertex with num @vertex with edges countained in @edges
func (g *GraphCRS) AddEdgesToVertex(vertex int, edges []int) {
	start := g.vertex[vertex]
	end := g.vertex[vertex+1]
	for i := vertex; i < len(g.vertex); i++ {
		g.vertex[i] += len(edges)
	}
	subSlice := append(g.edges[start:end], edges...)
	g.edges = append(append(g.edges[:start], subSlice...), g.edges[end:]...)
}

//AddVertexWithEdges function is adding new vertex with edges contained in @edges
func (g *GraphCRS) AddVertexWithEdges(edges []int) {
	g.vertex = append(g.vertex, len(g.edges))
	g.edges = append(g.edges, edges...)
}

//GetEdges returning edges of vertex number @ver
func (g *GraphCRS) GetEdges(ver int) []int {
	return g.edges[g.vertex[ver]:g.vertex[ver+1]]
}

//Size return amount of added vertex
func (g *GraphCRS) Size() int {
	return len(g.vertex)
}

//Print printing graph as struct of addiction to stdout
func (g *GraphCRS) Print() {
	for i, st := range g.vertex {
		var end int
		if i == len(g.vertex)-1 {
			end = len(g.edges)
		} else {
			end = g.vertex[i+1]
		}
		for _, elem := range g.edges[st:end] {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
}
