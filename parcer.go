package graphlib

import (
	"bufio"
	"os"
)

//Parser is object for parsing graph from file
type Parser struct {
	e error
}

//ParseUnweightedUndirectedGraphFromFile is parsing unweighted undirected graph from file with struct
func (p *Parser) ParseUnweightedUndirectedGraphFromFile(path string) (*Graph, error) {
	result := new(Graph)
	result.Init()

	file, err := os.Open(path)

	if err != nil {
		p.e = err
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		edges := make([]int, countNumbers(scanner.Text()))
		numbers := strings.Fields(scanner.Text())
		for i, num := range numbers {
			edges[i],err = strconv.Atoi(num)
			if err != nil{
				p.e = err
				return result,err
			}
		}

		result.AddVertexWithEdges(edges)
}

//Err return last parser error
func (p *Parser) Err() error {
	return p.e
}
