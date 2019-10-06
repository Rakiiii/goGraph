package graph

//Graph is struct that describe graph with vertex edges popose in map
type Graph struct {
	graph      map[int][]int
	lastVertex int
}

//AddVertex function is adding new vertex without edges t graph
func (g *Graph) AddVertex() {

	g.graph[g.lastVertex] = make([]int, 1)

	g.lastVertex++
}

//AddEdgesToVertex function is setting edges of vertex with num @vertex with edges countained in @edges
func (g *Graph) AddEdgesToVertex(vertex int, edges []int) {
	g.graph[vertex] = edges

}

//AddVertexWithEdges function is adding new vertex with edges contained in @edges
func (g *Graph) AddVertexWithEdges(edges []int) {
	g.graph[g.lastVertex] = edges

	g.lastVertex++
}

//GetVertexEdges returning edges of vertex number @ver
func (g *Graph) GetVertexEdges(ver int) []int {
	return g.graph[ver]
}

//Size return amount of added vertex
func (g *Graph) Size() int {
	return g.lastVertex
}

//lates add some comm for update package


