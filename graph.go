package graphlib

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

}

//AddEdgesToVertex function is setting edges of vertex with num @vertex with edges countained in @edges
func (g *Graph) AddEdgesToVertex(vertex int, edges []int) {
	g.graph[vertex] = edges

}

//AddVertexWithEdges function is adding new vertex with edges contained in @edges
func (g *Graph) AddVertexWithEdges(edges []int) {
	g.graph = append(g.graph, edges)

}

//GetEdges returning edges of vertex number @ver
func (g *Graph) GetEdges(ver int) []int {
	return g.graph[ver]
}

//Size return amount of added vertex
func (g *Graph) Size() int {
	return len(g.graph) - 1
}
