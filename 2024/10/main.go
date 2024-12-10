package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

func printMatr(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	grid := make([][]int, 0)

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "")
		temp := make([]int, 0)

		for _, str_num := range nums {
			num, _ := strconv.Atoi(str_num)
			temp = append(temp, num)
		}

		grid = append(grid, temp)
	}

	return grid, nil
}

func valid(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}

func BFS(grid [][]int, start Pos) int {
	directions := []Pos{
		{x: -1, y: 0},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
	}

	m, n := len(grid), len(grid[0])
	queue := []Pos{start}
	visited := make(map[Pos]bool)
	point := 0

	for len(queue) > 0 {
		pos := queue[0]
		curr := grid[pos.x][pos.y]

		if curr == 9 {
			if _, exist := visited[pos]; !exist {
				visited[pos] = true
				point++
			}
		}

		queue = queue[1:]

		for _, d := range directions {
			x := pos.x + d.x
			y := pos.y + d.y

			if valid(m, n, x, y) && curr+1 == grid[x][y] {
				queue = append(queue, Pos{x: x, y: y})
			}
		}
	}

	return point
}

func main() {
	grid, err := parseInput("input.txt")
	if err != nil {
		return
	}
	printMatr(grid)

	m, n := len(grid), len(grid[0])
	startPos := make([]Pos, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				startPos = append(startPos, Pos{x: i, y: j})
			}
		}
	}

	puzzle_1 := 0
	for _, start := range startPos {
		puzzle_1 += BFS(grid, start)
	}

	fmt.Println("puzzle 1:", puzzle_1)
}
