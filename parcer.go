package graph

import (
	"bufio"
	"os"
)

//Parser is object for parsing graph from file
type Parser struct{}

//ParseUnweightedUndirectedGraphFromFile is parsing unweighted undirected graph from file with struct
func (p *Parser) ParseUnweightedUndirectedGraphFromFile(path string) (*Graph, error) {

	result := new(Graph)

	file, err := os.Open(path)

	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		edges := make([]int, countNumbers(scanner.Text()))

		for i, num := range scanner.Text() {
			if num != ' ' {
				edges[i] = int(num)
			}
		}

		result.AddVertexWithEdges(edges)
	}

	if scanner.Err() != nil {
		return result, scanner.Err()
	}

	return result, nil

}
