package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	connections []int
}

type Graph struct {
	nodes map[int]Node
}

const (
	START  = 83
	lowerA = 97
	END    = 69
	lowerZ = 122
)

func test(a rune, b rune) bool {
	switch a {
	case START:
		a = lowerA
	case END:
		a = lowerZ
	}
	switch b {
	case START:
		b = lowerA
	case END:
		b = lowerZ
	}
	return a+1 == b || a >= b
}

func parse(in string) (g Graph, end int, sources []int) {
	g = Graph{make(map[int]Node)}

	content := strings.Split(in, "\n")

	for i, line := range content {
		for j := 0; j < len(line); j++ {
			// populate ends
			char := []rune(content[i])[j]
			if char == lowerA { //S
				sources = append(sources, j+i*len(line))
			} else if char == END { //E
				end = j + i*len(line)
			}

			var neighbors []int
			if j > 0 && test(char, rune(line[j-1])) {
				neighbors = append(neighbors, j-1+(i*len(line)))
			}
			if j < len(line)-1 && test(char, rune(line[j+1])) {
				neighbors = append(neighbors, j+1+(i*len(line)))
			}
			if i > 0 && test(char, rune(content[i-1][j])) {
				neighbors = append(neighbors, j+((i-1)*len(line)))
			}
			if i < len(content)-1 && test(char, rune(content[i+1][j])) {
				neighbors = append(neighbors, j+((i+1)*len(line)))
			}
			g.nodes[j+i*len(line)] = Node{
				connections: neighbors,
			}
		}
	}
	return g, end, sources
}

func dijkstra(g Graph, current int) map[int]int {
	completed := make(map[int]int)
	nodeDists := make(map[int]int)
	for i := 1; i < len(g.nodes); i++ {
		nodeDists[i] = math.MaxInt16
	}
	nodeDists[current] = 0

	for len(nodeDists) > 0 {
		for _, connection := range g.nodes[current].connections {
			newDist := nodeDists[current] + 1
			if newDist < nodeDists[connection] {
				nodeDists[connection] = newDist
			}
		}

		completed[current] = nodeDists[current]
		delete(nodeDists, current)

		current = func(nodeDists map[int]int) int {
			var out int
			lowest := math.MaxInt32
			for k, v := range nodeDists {
				if v < lowest {
					lowest = v
					out = k
				}
			}
			return out
		}(nodeDists)
	}
	return completed

}

func main() {
	in, _ := os.ReadFile("input")
	graph, end, sources := parse(string(in))

	lowest := math.MaxInt32
	for _, v := range sources {
		completed := dijkstra(graph, v)
		if completed[end] < lowest {
			lowest = completed[end]
		}
	}
	fmt.Println(lowest)
}
