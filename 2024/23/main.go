package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	FILENAME      = "input.txt"
	CYCLIC_LENGTH = 3
	CHIEF         = "t"
)

var (
	graphs      = map[string][]string{}
	cylicGraphs = map[string]bool{}
)

func parseInput() {
	file, _ := os.Open(FILENAME)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		a, b := nodes[0], nodes[1]

		if _, exist := graphs[a]; exist {
			graphs[a] = append(graphs[a], b)
		} else {
			graphs[a] = []string{b}
		}

		if _, exist := graphs[b]; exist {
			graphs[b] = append(graphs[b], a)
		} else {
			graphs[b] = []string{a}
		}
	}
}

func findCyclic(startNode, currNode string, length int, path []string) {
	if length < 0 {
		return
	}

	if startNode == currNode && length == 0 {
		slices.Sort(path)
		cylicGraphs[strings.Join(path, "")] = true
		return
	}

	path = append(path, currNode)
	for _, neighbor := range graphs[currNode] {
		findCyclic(startNode, neighbor, length-1, path)
	}
}

func solve() {
	for key := range graphs {
		if string(key[0]) != CHIEF {
			continue
		}
		findCyclic(key, key, CYCLIC_LENGTH, []string{})
	}

	puzzle_1 := len(cylicGraphs)
	fmt.Println("puzzle 1:", puzzle_1)
}

func main() {
	parseInput()
	solve()
}
