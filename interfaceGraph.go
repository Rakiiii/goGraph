package graphlib

//IGraph implaements functions for grpah
type IGraph interface {
	Init(v, e int)
	AmountOfVertex() int
	AmountOfEdges() int
	AddVertex()
	AddEdgesToVertex(vertex int, edges []int)
	AddVertexWithEdges(edges []int)
	GetEdges(ver int) []int
	Size() int
	Print()
}
