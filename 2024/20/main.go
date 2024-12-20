package main

import (
	"bufio"
	"fmt"
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

	THRESHOLD = 100
	MAX_CHEAT = 2
)

type Pos struct {
	x int
	y int
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

func bypass(grid [][]string, cost [][]int, pos Pos, amount int, prevCost int) int {
	count := 0

	for _, d := range DIRECTIONS {
		nx, ny := pos.x+d.x, pos.y+d.y

		if !valid(grid, nx, ny) {
			continue
		}

		if amount < MAX_CHEAT {
			count += bypass(grid, cost, Pos{x: nx, y: ny}, amount+1, prevCost)
			continue
		}

		if grid[nx][ny] == WALL {
			continue
		}

		if grid[nx][ny] == EMPTY || grid[nx][ny] == END {
			newCost := prevCost + amount
			currCost := cost[nx][ny]

			if currCost-newCost >= THRESHOLD {
				count++
			}
		}
	}

	return count
}

func cheat(grid [][]string, path []Pos, cost [][]int) int {
	count := 0
	for _, pos := range path {
		count += bypass(grid, cost, pos, 1, cost[pos.x][pos.y])
	}

	return count
}

func solve(fileName string) {
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
	puzzle_1 := cheat(grid, path, cost)
	fmt.Println("puzzle 1:", puzzle_1)
}

func main() {
	solve("input.txt")
}
