package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
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

// Bron-Kerbosch Algorithm
func bronKerbosch(R, P, X map[string]bool, maxClique *map[string]bool) {
	if len(P) == 0 && len(X) == 0 {
		if len(R) > len(*maxClique) {
			newClique := make(map[string]bool)
			for key := range R {
				(newClique)[key] = true
			}
			(*maxClique) = newClique
		}
		return
	}

	for key := range P {
		newR := make(map[string]bool)
		for node := range R {
			newR[node] = true
		}
		newR[key] = true

		newP := make(map[string]bool)
		newX := make(map[string]bool)
		for _, neighbor := range graphs[key] {
			if P[neighbor] {
				newP[neighbor] = true
			}
			if X[neighbor] {
				newX[neighbor] = true
			}
		}

		bronKerbosch(newR, newP, newX, maxClique)

		delete(P, key)
		X[key] = true
	}
}

func findInterconnected() []string {
	R := make(map[string]bool)
	P := make(map[string]bool)
	X := make(map[string]bool)

	for node := range graphs {
		P[node] = true
	}

	maxClique := make(map[string]bool)
	bronKerbosch(R, P, X, &maxClique)

	ans := make([]string, 0, len(maxClique))
	for v := range maxClique {
		ans = append(ans, v)
	}
	slices.Sort(ans)

	return ans
}

func solve() {
	start := time.Now()
	for key := range graphs {
		if string(key[0]) != CHIEF {
			continue
		}
		findCyclic(key, key, CYCLIC_LENGTH, []string{})
	}

	puzzle_1 := len(cylicGraphs)
	puzzle_2 := strings.Join(findInterconnected(), ",")

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Println("puzzle 2:", puzzle_2)
}

func main() {
	parseInput()
	solve()
}
