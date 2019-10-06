package graphlib

//Graph is struct that describe graph with vertex edges popose in map
type Graph struct {
	graph      [][]int
	//lastVertex int
}

//Init is initialize graph with nonnil map
func (g *Graph)Init(){
	g.graph = make([][]int,0)
	//g.lastVertex = 0
}

//AddVertex function is adding new vertex without edges t graph
func (g *Graph) AddVertex() {

	g.graph = append(g.graph,make([]int,1))

}

//AddEdgesToVertex function is setting edges of vertex with num @vertex with edges countained in @edges
func (g *Graph) AddEdgesToVertex(vertex int, edges []int) {
	g.graph[vertex] = edges

}

//AddVertexWithEdges function is adding new vertex with edges contained in @edges
func (g *Graph) AddVertexWithEdges(edges []int) {
	g.graph = append(g.graph,edges)

}

//GetEdges returning edges of vertex number @ver
func (g *Graph) GetEdges(ver int) []int {
	return g.graph[ver]
}

//Size return amount of added vertex
func (g *Graph) Size() int {
	return len(g.graph)
}
