package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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

// using BFS in each step either go forward or rotate
func BFS(grid [][]string, startPos Pos) int {
	visited := make(map[Pos]bool)
	visited[startPos] = true
	queue := []Move{{pos: startPos, cost: 0}}
	cost := math.MaxInt

	for len(queue) > 0 {
		currPos := queue[0].pos
		currCost := queue[0].cost
		queue = queue[1:]

		if grid[currPos.x][currPos.y] == END {
			cost = min(cost, currCost)
			continue
		}

		d := DIRECTIONS[currPos.dir]
		nx, ny := currPos.x+d.x, currPos.y+d.y

		// check for rotate, doesn't matter if encounter a wall or no
		// rotate can be performed any step
		for _, clockwise := range []bool{true, false} {
			rotatedPos := Pos{
				x:   currPos.x,
				y:   currPos.y,
				dir: rotate(currPos.dir, clockwise),
			}

			if !exist(visited, rotatedPos) {
				visited[rotatedPos] = true
				queue = append(queue, Move{pos: rotatedPos, cost: currCost + ROTATE_COST})
			}
		}

		// if next is not wall then can move forward
		if grid[nx][ny] != WALL {
			forwardPos := Pos{
				x:   nx,
				y:   ny,
				dir: currPos.dir,
			}

			if !exist(visited, forwardPos) {
				visited[forwardPos] = true
				queue = append(queue, Move{pos: forwardPos, cost: currCost + FORWARD_COST})
			}
		}
	}

	return cost
}

// using Djikstra to find the smallest cost
func Djikstra(grid [][]string, startPos Pos) int {
	visited := make(map[Pos]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Move{pos: startPos, cost: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Move)
		currPos := curr.pos
		currCost := curr.cost

		if visited[currPos] {
			continue
		}
		visited[currPos] = true

		if grid[currPos.x][currPos.y] == END {
			return currCost
		}

		// check for rotate, doesn't matter if encounter a wall or no
		// rotate can be performed any step
		for _, clockwise := range []bool{true, false} {
			rotatedPos := Pos{
				x:   currPos.x,
				y:   currPos.y,
				dir: rotate(currPos.dir, clockwise),
			}

			if !visited[rotatedPos] {
				heap.Push(pq, &Move{pos: rotatedPos, cost: currCost + ROTATE_COST})
			}
		}

		d := DIRECTIONS[currPos.dir]
		nx, ny := currPos.x+d.x, currPos.y+d.y
		if grid[nx][ny] != WALL {
			forwardPos := Pos{
				x:   nx,
				y:   ny,
				dir: currPos.dir,
			}

			if !visited[forwardPos] {
				heap.Push(pq, &Move{pos: forwardPos, cost: currCost + FORWARD_COST})
			}
		}
	}

	return -1
}

func DFS(grid [][]string, visited map[Pos]bool, targetCost int, move Move, path []Pos, result *[][]Pos) {
	currPos := move.pos
	currCost := move.cost

	if visited[currPos] {
		return
	}
	visited[currPos] = true
	path = append(path, currPos)

	if currCost > targetCost {
		visited[currPos] = false
		return
	}

	if grid[currPos.x][currPos.y] == END && currCost == targetCost {
		validPath := make([]Pos, len(path))
		copy(validPath, path)
		*result = append(*result, validPath)
	}

	for _, clockwise := range []bool{true, false} {
		rotatedPos := Pos{
			x:   currPos.x,
			y:   currPos.y,
			dir: rotate(currPos.dir, clockwise),
		}

		if !visited[rotatedPos] {
			DFS(grid, visited, targetCost, Move{pos: rotatedPos, cost: currCost + ROTATE_COST}, path, result)
		}
	}

	d := DIRECTIONS[currPos.dir]
	nx, ny := currPos.x+d.x, currPos.y+d.y
	if grid[nx][ny] != WALL {
		forwardPos := Pos{
			x:   nx,
			y:   ny,
			dir: currPos.dir,
		}

		if !visited[forwardPos] {
			DFS(grid, visited, targetCost, Move{pos: forwardPos, cost: currCost + FORWARD_COST}, path, result)
		}
	}

	visited[currPos] = false
	path = path[:len(path)-1]
}

func findPaths(grid [][]string, start Pos, targetCost int) [][]Pos {
	visited := make(map[Pos]bool)
	var results [][]Pos
	DFS(grid, visited, targetCost, Move{pos: start, cost: 0}, []Pos{}, &results)
	return results
}

func fillGrid(grid *[][]string, path []Pos) {
	for _, p := range path {
		(*grid)[p.x][p.y] = SEAT
	}
}

func countSeat(grid [][]string) int {
	m, n := len(grid), len(grid[0])
	count := 0

	for i := range m {
		for j := range n {
			if grid[i][j] == SEAT {
				count++
			}
		}
	}

	return count
}

func solve(grid [][]string) {
	startPos := findStartPos(grid)
	puzzle_1 := Djikstra(grid, startPos)
	fmt.Println("puzzle 1:", puzzle_1)

	results := findPaths(grid, startPos, puzzle_1)
	for _, path := range results {
		fillGrid(&grid, path)
	}

	puzzle_2 := countSeat(grid)
	fmt.Println("puzzle 2:", puzzle_2)
}

func main() {
	start := time.Now()
	grid := parseInput("input.txt")

	solve(grid)
	fmt.Println("time elapsed:", time.Since(start))
}
