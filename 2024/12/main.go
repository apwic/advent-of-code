package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

var directions = []Pos{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func printMatr(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	grid := make([][]string, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
	}

	return grid, nil
}

func valid(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}

func createRegion(grid [][]string, visited *map[Pos]bool, i, j int) (int, int) {
	m, n := len(grid), len(grid[0])
	name := grid[i][j]
	queue := []Pos{{i, j}}

	(*visited)[Pos{x: i, y: j}] = true

	area, perimeter := 0, 0

	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		x, y := pop.x, pop.y

		area++

		for _, d := range directions {
			next := Pos{x: x + d.x, y: y + d.y}

			if valid(m, n, next.x, next.y) {
				if grid[next.x][next.y] != name {
					// increment perimeter if neighbor cell is not part of the region
					perimeter++
				} else if !(*visited)[next] {
					queue = append(queue, next)
					(*visited)[next] = true
				}
			} else {
				// increment perimeter for out-of-bounds edges
				perimeter++
			}
		}
	}

	return area, perimeter
}

func findRegion(grid [][]string) int {
	m, n := len(grid), len(grid[0])
	visited := make(map[Pos]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited[Pos{x: i, y: j}] = false
		}
	}

	price := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[Pos{x: i, y: j}] {
				continue
			}
			area, param := createRegion(grid, &visited, i, j)
			price += area * param
		}
	}

	return price
}

func main() {
	grid, err := parseInput("input.txt")
	if err != nil {
		return
	}

	puzzle_1 := findRegion(grid)
	fmt.Println("puzzle 1:", puzzle_1)
}
