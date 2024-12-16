package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	START = "S"
	END   = "E"
	WALL  = "#"
	SEAT  = "O"

	DIRECTIONS = map[string]Pos{
		">": {x: 0, y: 1},
		"v": {x: 1, y: 0},
		"<": {x: 0, y: -1},
		"^": {x: -1, y: 0},
	}

	FORWARD_COST = 1
	ROTATE_COST  = 1000
)

type Point struct {
	x int
	y int
}

type Pos struct {
	x   int
	y   int
	dir string
}

func printMatr(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(fileName string) [][]string {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
	}

	return grid
}

func findStartPos(grid [][]string) Pos {
	m, n := len(grid), len(grid[0])

	for i := range m {
		for j := range n {
			if grid[i][j] == START {
				return Pos{
					x:   i,
					y:   j,
					dir: ">",
				}
			}
		}
	}

	return Pos{}
}

func rotate(dir string, clockwise bool) string {
	directions := []string{">", "v", "<", "^"}
	for i, d := range directions {
		if d == dir {
			if clockwise {
				return directions[(i+1)%len(directions)]
			}
			return directions[(i+len(directions)-1)%len(directions)]
		}
	}
	return ""
}

func exist(visited map[Pos]bool, pos Pos) bool {
	if _, exist := visited[pos]; exist {
		return true
	}

	return false
}

// using Dijkstra to find the smallest cost
func Dijkstra(grid [][]string, startPos Pos) (int, int) {
	visited := make(map[Pos]int)
	prev := make(map[Pos][]Pos)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Move{pos: startPos, cost: 0, prev: Pos{x: -1, y: -1}})

	var lastPos Pos
	var score int

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Move)

		if _, exist := visited[curr.pos]; exist {
			if curr.cost == visited[curr.pos] {
				prev[curr.pos] = append(prev[curr.pos], curr.prev)
			}
			continue
		}

		visited[curr.pos] = curr.cost

		if curr.prev.x != -1 && curr.prev.y != -1 {
			prev[curr.pos] = []Pos{curr.prev}
		}

		if grid[curr.pos.x][curr.pos.y] == END {
			score = curr.cost
			lastPos = curr.pos
			break
		}

		// check for rotate, doesn't matter if encounter a wall or no
		// rotate can be performed any step
		for _, clockwise := range []bool{true, false} {
			rotatedPos := Pos{
				x:   curr.pos.x,
				y:   curr.pos.y,
				dir: rotate(curr.pos.dir, clockwise),
			}

			heap.Push(pq, &Move{
				pos:  rotatedPos,
				cost: curr.cost + ROTATE_COST,
				prev: curr.pos,
			})
		}

		d := DIRECTIONS[curr.pos.dir]
		nx, ny := curr.pos.x+d.x, curr.pos.y+d.y
		if grid[nx][ny] != WALL {
			forwardPos := Pos{
				x:   nx,
				y:   ny,
				dir: curr.pos.dir,
			}

			heap.Push(pq, &Move{
				pos:  forwardPos,
				cost: curr.cost + FORWARD_COST,
				prev: curr.pos,
			})
		}
	}

	points := map[Point]bool{}
	nodes := []Pos{lastPos}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		points[Point{x: node.x, y: node.y}] = true
		nodes = append(nodes, prev[node]...)
	}

	return score, len(points)
}

func solve(grid [][]string) {
	startPos := findStartPos(grid)
	puzzle_1, puzzle_2 := Dijkstra(grid, startPos)
	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Println("puzzle 2:", puzzle_2)
}

func main() {
	start := time.Now()
	grid := parseInput("input.txt")

	solve(grid)
	fmt.Println("time elapsed:", time.Since(start))
}
