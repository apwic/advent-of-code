package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	EMPTY = "."
	WALL  = "#"
	START = "S"
	END   = "E"

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

type Shortcut struct {
	start Pos
	end   Offset
}

type Offset struct {
	pos      Pos
	distance int
}

func printMatrStr(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func printMatrInt(grid [][]int) {
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
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	return grid
}

func valid(grid [][]string, i, j int) bool {
	m, n := len(grid), len(grid[0])
	return 0 <= i && i < m && 0 <= j && j < n
}

func findPos(grid [][]string) (Pos, Pos) {
	m, n := len(grid), len(grid[0])
	start := Pos{}
	end := Pos{}

	for i := range m {
		for j := range n {
			if grid[i][j] == START {
				start = Pos{x: i, y: j}
			}

			if grid[i][j] == END {
				end = Pos{x: i, y: j}
			}
		}
	}

	return start, end
}

func DFS(grid [][]string, curr Pos, cost *[][]int, path *[]Pos) {
	// early exit
	if grid[curr.x][curr.y] == WALL || (*cost)[curr.x][curr.y] != 0 {
		return
	}

	// update cost
	if len(*path) == 0 {
		(*cost)[curr.x][curr.y] = 1
	} else {
		prevPos := (*path)[len(*path)-1]
		prevCost := (*cost)[prevPos.x][prevPos.y]
		(*cost)[curr.x][curr.y] = prevCost + 1
	}

	// update path
	(*path) = append((*path), curr)

	if grid[curr.x][curr.y] == END {
		return
	}

	for _, d := range DIRECTIONS {
		nx, ny := d.x+curr.x, d.y+curr.y
		DFS(grid, Pos{x: nx, y: ny}, cost, path)
	}
}

func distance(start, end Pos) int {
	xDistance := math.Abs(float64(start.x - end.x))
	yDistance := math.Abs(float64(start.y - end.y))
	return int(xDistance + yDistance)
}

func offsets(from Pos, radius int) []Offset {
	result := []Offset{}

	for y := radius * -1; y <= radius; y++ {
		for x := radius * -1; x <= radius; x++ {
			candidatePoint := Pos{from.x + x, from.y + y}
			candidate := Offset{
				candidatePoint,
				distance(from, candidatePoint),
			}

			if candidate.distance > 0 && candidate.distance <= radius {
				result = append(result, candidate)
			}
		}
	}

	return result
}

func cheat(grid [][]string, path []Pos, cost [][]int, radius int) map[int]int {
	shortcuts := make(map[Shortcut]int)
	for _, current := range path {
		step := cost[current.x][current.y]
		offsets := offsets(current, radius)
		for _, offset := range offsets {
			if !valid(grid, offset.pos.x, offset.pos.y) {
				continue
			}
			routeStep := cost[offset.pos.x][offset.pos.y]

			saving := routeStep - step - offset.distance
			if saving > 0 {
				shortcuts[Shortcut{current, offset}] = saving
			}
		}
	}

	// Transform to summary.
	result := make(map[int]int)
	for _, saving := range shortcuts {
		result[saving]++
	}

	return result
}

func solve(puzzle int, fileName string, threshold int, radius int) {
	grid := parseInput(fileName)
	start, _ := findPos(grid)
	m, n := len(grid), len(grid[0])
	cost := make([][]int, m)
	path := make([]Pos, 0)

	for i := range m {
		for range n {
			cost[i] = append(cost[i], 0)
		}
	}

	DFS(grid, start, &cost, &path)
	result := cheat(grid, path, cost, radius)

	count := 0
	for time, freq := range result {
		if time >= threshold {
			count += freq
		}
	}

	fmt.Printf("puzzle %d: %d\n", puzzle, count)
}

func main() {
	solve(1, "input.txt", 100, 2)
	solve(2, "input.txt", 100, 20)
}
