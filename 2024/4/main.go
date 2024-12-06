package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func valid(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}

func check(grid [][]string, i int, j int, dx int, dy int) bool {
	target := "XMAS"
	if !valid(len(grid), len(grid[0]), i+(len(target)-1)*dx, j+(len(target)-1)*dy) {
		return false
	}

	var b strings.Builder
	for k := 0; k < len(target); k++ {
		b.WriteString(grid[i][j])
		i += dx
		j += dy
	}

	return b.String() == target
}

func check2(grid [][]string, i int, j int) bool {
	if !valid(len(grid), len(grid[0]), i+2, j+2) {
		return false
	}

	var b1, b2 strings.Builder

	for k := 0; k < 3; k++ {
		b1.WriteString(grid[i+k][j+k])
		b2.WriteString(grid[i+k][j+2-k])
	}

	return (b1.String() == "MAS" || b1.String() == "SAM") && (b2.String() == "MAS" || b2.String() == "SAM")
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Split(line, "")
		grid = append(grid, make([]string, 0))
		idx := len(grid) - 1

		for _, ch := range part {
			grid[idx] = append(grid[idx], ch)
		}
	}

	puzzle_1 := 0
	puzzle_2 := 0
	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]

				if check(grid, i, j, dx, dy) {
					puzzle_1++
				}
			}

			if check2(grid, i, j) {
				puzzle_2++
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)

	fmt.Println("puzzle 1: ", puzzle_1)
	fmt.Println("puzzle 2: ", puzzle_2)
}
