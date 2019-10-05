package graph

//Graph is struct that describe graph with vertex edges popose in map
type Graph struct {
	graph      map[int][]int
	lastVertex int
}

//addVertex function is adding new vertex without edges t graph
func (g *Graph) addVertex() {

	g.graph[g.lastVertex] = make([]int, 1)

	g.lastVertex++
}

//addEdgesToVertex function is setting edges of vertex with num @vertex with edges countained in @edges
func (g *Graph) addEdgesToVertex(vertex int, edges []int) {
	g.graph[vertex] = edges

}

//addVertexWithEdges function is adding new vertex with edges contained in @edges
func (g *Graph) addVertexWithEdges(edges []int) {
	g.graph[g.lastVertex] = edges

	g.lastVertex++
}

//getVertexEdges returning edges of vertex number @ver
func (g *Graph) getVertexEdges(ver int) []int {
	return g.graph[ver]
}

//Size return amount of added vertex
func (g *Graph) Size() int {
	return g.lastVertex
}
