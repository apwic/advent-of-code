package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func printMatr(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(fileName string) ([][]string, map[string][][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]string
	freq := make(map[string][][]int)
	i := 0
	j := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		temp := make([]string, 0)
		j = 0

		for _, ch := range line {
			temp = append(temp, ch)

			if ch == "." {
				j++
				continue
			}

			if _, exist := freq[ch]; exist {
				freq[ch] = append(freq[ch], []int{i, j})
			} else {
				freq[ch] = make([][]int, 0)
				freq[ch] = append(freq[ch], []int{i, j})
			}
			j++
		}

		grid = append(grid, temp)
		i++
	}

	return grid, freq, nil
}

func abs(a, b int) int {
	diff := a - b
	if diff < 0 {
		return diff * -1
	}

	return diff
}

func valid(m, n, x, y int) bool {
	return 0 <= x && x < m && 0 <= y && y < n
}

func main() {
	start := time.Now()
	grid, freq, err := parseInput("input.txt")
	if err != nil {
		return
	}

	m, n := len(grid), len(grid[0])

	for _, val := range freq {
		for i := 0; i < len(val); i++ {
			for j := i + 1; j < len(val); j++ {
				x_a, y_a := val[i][0], val[i][1]
				x_b, y_b := val[j][0], val[j][1]

				if x_a == x_b && y_a == y_b {
					continue
				}

				diff_x := x_a - x_b
				diff_y := y_a - y_b

				if valid(m, n, x_a+diff_x, y_a+diff_y) {
					grid[x_a+diff_x][y_a+diff_y] = "#"
				}

				if valid(m, n, x_b-diff_x, y_b-diff_y) {
					grid[x_b-diff_x][y_b-diff_y] = "#"
				}
			}
		}
	}

	puzzle_1 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "#" {
				puzzle_1++
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Time elapsed : ", elapsed)

	fmt.Println("puzzle 1: ", puzzle_1)
}
