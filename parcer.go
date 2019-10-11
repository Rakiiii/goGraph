package graphlib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Parser is object for parsing graph from file
type Parser struct {
	e error
}

//ParseUnweightedUndirectedGraphFromFile is parsing unweighted undirected graph from file with struct
func (p *Parser) ParseUnweightedUndirectedGraphFromFile(path string) (*Graph, error) {

	result := new(Graph)

	file, err := os.Open(path)

	if err != nil {
		p.e = err
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	number := strings.Fields(scanner.Text())

	amV, err := strconv.Atoi(number[0])
	if err != nil {
		return result, err
	}

	amE, err := strconv.Atoi(number[1])
	if err != nil {
		return result, err
	}

	result.Init(amV, amE)

	for scanner.Scan() {
		edges := make([]int, countNumbers(scanner.Text()))
		numbers := strings.Fields(scanner.Text())
		for i, num := range numbers {
			edges[i], err = strconv.Atoi(num)
			if err != nil {
				p.e = err
				return result, err
			}
		}

		result.AddVertexWithEdges(edges)
	}

	if scanner.Err() != nil {
		p.e = scanner.Err()
		return result, scanner.Err()
	}

	return result, nil

}

//Err return last parser error
func (p *Parser) Err() error {
	return p.e
}
