package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	OBSTACLE = "#"
	EMPTY    = "."
	VISITED  = "X"
	GUARD    = "^"
)

type Vector struct {
	x        int
	y        int
	rotation int
}

func printMatr(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func turn(rotation int) (int, int) {
	rotation %= 360

	if rotation == 0 {
		return -1, 0
	} else if rotation == 90 {
		return 0, 1
	} else if rotation == 180 {
		return 1, 0
	} else {
		return 0, -1
	}
}

func valid(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}

func path(grid *[][]string, rotation int, i int, j int) {
	m, n := len(*grid), len((*grid)[0])
	dx, dy := turn(rotation)
	i, j = i+dx, j+dy

	for valid(m, n, i, j) {
		curr := (*grid)[i][j]
		if curr == EMPTY || curr == VISITED {
			(*grid)[i][j] = VISITED
		}

		if curr == OBSTACLE {
			i, j = i-dx, j-dy
			rotation += 90
			dx, dy = turn(rotation)
		}

		i, j = i+dx, j+dy
	}
}

func path2(grid [][]string, rotation int, i int, j int, obs_i int, obs_j int) bool {
	gridCopy := make([][]string, len(grid))
	for idx := range grid {
		gridCopy[idx] = append([]string(nil), grid[idx]...)
	}

	visited := make(map[Vector]bool)
	m, n := len(gridCopy), len((gridCopy)[0])
	dx, dy := turn(rotation)
	x, y := i+dx, j+dy

	// add obstacle here
	gridCopy[obs_i][obs_j] = OBSTACLE

	for valid(m, n, x, y) {
		curr := (gridCopy)[x][y]
		vector := Vector{
			x:        x,
			y:        y,
			rotation: rotation,
		}

		if visited[vector] {
			return true
		}

		if curr == EMPTY || curr == VISITED {
			visited[vector] = true
			(gridCopy)[x][y] = VISITED
		}

		if curr == OBSTACLE {
			x, y = x-dx, y-dy
			rotation = (rotation + 90) % 360
			dx, dy = turn(rotation)
		}

		x, y = x+dx, y+dy
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]string
	start_i, start_j := 0, 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		temp := make([]string, 0)

		for _, ch := range line {
			if ch == EMPTY {
				temp = append(temp, EMPTY)
			} else if ch == OBSTACLE {
				temp = append(temp, OBSTACLE)
			} else if ch == GUARD {
				start_i = len(grid)
				start_j = len(temp)
				temp = append(temp, VISITED)
			}
		}

		grid = append(grid, temp)
	}

	gridCopy := make([][]string, len(grid))
	for idx := range grid {
		gridCopy[idx] = append([]string(nil), grid[idx]...)
	}

	puzzle_1 := 0
	puzzle_2 := 0
	path(&gridCopy, 0, start_i, start_j)
	for obs_i, row := range gridCopy {
		for obs_j, curr := range row {
			if curr != VISITED {
				continue
			}

			puzzle_1++

			if path2(grid, 0, start_i, start_j, obs_i, obs_j) {
				puzzle_2++
			}
		}
	}

	fmt.Println("puzzle 1: ", puzzle_1)
	fmt.Println("puzzle 2: ", puzzle_2)
}
