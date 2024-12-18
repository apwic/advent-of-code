package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	GRID_X = 71
	GRID_Y = 71
	SIZE   = 1024

	TARGET = Pos{x: GRID_X - 1, y: GRID_Y - 1}

	EMPTY = "."
	WALL  = "#"

	DIRECTIONS = []Pos{
		{x: 0, y: 1},  // RIGHT
		{x: 1, y: 0},  // DOWN
		{x: 0, y: -1}, // LEFT
		{x: -1, y: 0}, // UP
	}
)

type Pos struct {
	x int
	y int
}

type Node struct {
	pos  Pos
	cost int
}

func printMatr(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(fileName string) []Pos {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := []Pos{}
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		y, _ := strconv.Atoi(text[0])
		x, _ := strconv.Atoi(text[1])
		pos = append(pos, Pos{x: x, y: y})
	}

	return pos
}

func createGrid(pos []Pos) [][]string {
	grid := make([][]string, GRID_X)

	for i := range GRID_X {
		for range GRID_Y {
			grid[i] = append(grid[i], EMPTY)
		}
	}

	for _, p := range pos {
		grid[p.x][p.y] = WALL
	}

	return grid
}

func valid(i, j int) bool {
	return 0 <= i && i < GRID_X && 0 <= j && j < GRID_Y
}

func BFS(grid [][]string) int {
	queue := []Node{
		{pos: Pos{x: 0, y: 0}, cost: 0},
	}
	visited := map[Pos]bool{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node.pos] {
			continue
		}
		visited[node.pos] = true

		if node.pos == TARGET {
			return node.cost
		}

		for _, d := range DIRECTIONS {
			nx, ny := node.pos.x+d.x, node.pos.y+d.y
			if valid(nx, ny) && grid[nx][ny] == EMPTY {
				queue = append(queue, Node{
					pos:  Pos{x: nx, y: ny},
					cost: node.cost + 1,
				})
			}
		}
	}

	return -1
}

func checkExit(pos []Pos) Pos {
	for n := SIZE + 1; n < len(pos); n++ {
		grid := createGrid(pos[:n])
		cost := BFS(grid)

		if cost == -1 {
			return pos[n-1]
		}
	}

	return Pos{}
}

func solve(fileName string) {
	start := time.Now()
	pos := parseInput(fileName)
	grid := createGrid(pos[:SIZE])
	puzzle_1 := BFS(grid)
	puzzle_2 := checkExit(pos)

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Printf("puzzle 2: %d,%d\n", puzzle_2.y, puzzle_2.x)
}

func main() {
	solve("input.txt")
}
