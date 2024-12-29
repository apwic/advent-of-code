package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILENAME = "input.txt"
	ROW      = 6
	COL      = 5

	EMPTY = "."
)

var (
	locks = make([][]int, 0)
	keys  = make([][]int, 0)
)

func parseInput() {
	file, _ := os.Open(FILENAME)
	scanner := bufio.NewScanner(file)

	grid := make([]int, 0)
	isKey := false
	idx := 0

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")

		if len(text) == 0 {
			continue
		}

		// check for the first line of the grid
		// reset the grid
		if idx == 0 {
			isKey = text[0] == EMPTY
			grid = make([]int, 0)

			for range COL {
				grid = append(grid, -1)
			}
		}

		for i := range COL {
			if string(text[i]) != EMPTY {
				grid[i]++
			}
		}

		// end of the grid, add to corresponding map
		if idx == ROW {
			addKey(isKey, grid)
			idx = 0
			continue
		}

		idx++
	}
}

func addKey(isKey bool, grid []int) {
	if isKey {
		keys = append(keys, grid)
	} else {
		locks = append(locks, grid)
	}
}

// unused for now
func makeKey(grid []int) string {
	gridStr := make([]string, 0)

	for _, num := range grid {
		numStr := strconv.Itoa(num)
		gridStr = append(gridStr, numStr)
	}

	return strings.Join(gridStr, "")
}

func isFit(lock, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] >= ROW {
			return false
		}
	}

	return true
}

func findFit() int {
	fit := 0

	for _, lock := range locks {
		for _, key := range keys {
			if isFit(lock, key) {
				fit++
			}
		}
	}

	return fit
}

func solve() {
	puzzle_1 := findFit()
	fmt.Println("puzzle 1:", puzzle_1)
}

func main() {
	parseInput()
	solve()
}
