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

	counter := 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		edges := make([]int, len(numbers))
		for i, num := range numbers {
			edges[i], err = strconv.Atoi(num)
			if err != nil {
				p.e = err
				return result, err
			}
		}

		result.AddEdgesToVertex(counter, edges)
		counter++
	}

	if scanner.Err() != nil {
		p.e = scanner.Err()
		return result, scanner.Err()
	}

	return result, nil

}

//ParseUnweightedUndirectedGraphFromFile is parsing unweighted undirected graph from file with struct
func (p *Parser) ParseUnweightedUndirectedGraphFromFileToFastGraph(path string) (*FastGraph, error) {

	result := new(FastGraph)

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

	counter := 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		edges := make([]int, len(numbers))
		for i, num := range numbers {
			edges[i], err = strconv.Atoi(num)
			if err != nil {
				p.e = err
				return result, err
			}
		}

		result.AddEdgesToVertex(counter, edges)
		counter++
	}

	if scanner.Err() != nil {
		p.e = scanner.Err()
		return result, scanner.Err()
	}

	return result, nil

}

//ParseUnweightedUndirectedGraphCRSFromFile parse grpah from file to CRS form
func (p *Parser) ParseUnweightedUndirectedGraphCRSFromFile(path string) (*GraphCRS, error) {
	g, err := p.ParseUnweightedUndirectedGraphFromFile(path)
	if err != nil {
		return nil, err
	}
	return g.ConvertToCRS(), nil
}

//Err return last parser error
func (p *Parser) Err() error {
	return p.e
}
