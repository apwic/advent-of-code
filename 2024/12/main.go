package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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

func createRegion(grid [][]string, visited *map[Pos]bool, i, j int) (int, int, int) {
	m, n := len(grid), len(grid[0])
	name := grid[i][j]
	queue := []Pos{{i, j}}

	(*visited)[Pos{x: i, y: j}] = true

	area, perimeter, corner := 0, 0, 0

	// diagonal directions to check corners
	diagonalDirs := []Pos{
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		x, y := pop.x, pop.y

		area++

		// check for corners
		for _, d := range diagonalDirs {
			diag := Pos{x: x + d.x, y: y + d.y}
			adj1 := Pos{x: x + d.x, y: y}
			adj2 := Pos{x: x, y: y + d.y}

			validDiag := valid(m, n, diag.x, diag.y)
			validAdj1 := valid(m, n, adj1.x, adj1.y)
			validAdj2 := valid(m, n, adj2.x, adj2.y)

			outOfBounds := !validDiag && !validAdj1 && !validAdj2

			checkAdj1 := !validAdj1 || (validAdj1 && grid[adj1.x][adj1.y] != name)

			checkAdj2 := !validAdj2 || (validAdj2 && grid[adj2.x][adj2.y] != name)

			checkDiag := validDiag && grid[diag.x][diag.y] != name && !checkAdj1 && !checkAdj2

			if outOfBounds || (checkAdj1 && checkAdj2) || checkDiag {
				corner++
			}
		}

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

	return area, perimeter, corner
}

func findRegion(grid [][]string) (int, int) {
	m, n := len(grid), len(grid[0])
	visited := make(map[Pos]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited[Pos{x: i, y: j}] = false
		}
	}

	price1 := 0
	price2 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[Pos{x: i, y: j}] {
				continue
			}
			area, param, corner := createRegion(grid, &visited, i, j)
			price1 += area * param
			price2 += area * corner
		}
	}

	return price1, price2
}

func main() {
	start := time.Now()
	grid, err := parseInput("input.txt")
	if err != nil {
		return
	}

	puzzle_1, puzzle_2 := findRegion(grid)

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Println("puzzle 2:", puzzle_2)
}
