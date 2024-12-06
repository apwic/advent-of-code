package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	OBSTACLE = 9
	EMPTY    = 0
	VISITED  = 1
)

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

func path(grid *[][]int, rotation int, i int, j int) {
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var grid [][]int
	start_i, start_j := 0, 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		temp := make([]int, 0)

		for _, ch := range line {
			if ch == "." {
				temp = append(temp, EMPTY)
			} else if ch == "#" {
				temp = append(temp, OBSTACLE)
			} else if ch == "^" {
				start_i = len(grid)
				start_j = len(temp)
				temp = append(temp, VISITED)
			}
		}

		grid = append(grid, temp)
	}

	puzzle_1 := 0
	path(&grid, 0, start_i, start_j)
	for _, row := range grid {
		fmt.Println(row)
		for _, curr := range row {
			if curr == VISITED {
				puzzle_1++
			}
		}
	}

	fmt.Println("puzzle 1: ", puzzle_1)
}
