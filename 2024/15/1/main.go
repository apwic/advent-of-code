package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	WALL       = "#"
	ROBOT      = "@"
	BOX        = "O"
	EMPTY      = "."
	DIRECTIONS = map[string]Pos{
		"<": {x: 0, y: -1},
		"v": {x: 1, y: 0},
		">": {x: 0, y: 1},
		"^": {x: -1, y: 0},
	}

	POINT = 100
)

type Pos struct {
	x int
	y int
}

func printMatr(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(fileName string) ([][]string, []string) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	ifGrid := true
	grid := make([][]string, 0)
	moves := make([]string, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if len(line) == 0 {
			ifGrid = false
			continue
		}

		if ifGrid {
			grid = append(grid, line)
		} else {
			for _, l := range line {
				moves = append(moves, l)
			}
		}
	}

	return grid, moves
}

func findRobotPos(grid [][]string) Pos {
	m, n := len(grid), len(grid[0])
	robotPos := Pos{}

	for i := range m {
		for j := range n {
			if grid[i][j] == ROBOT {
				robotPos.x = i
				robotPos.y = j
				return robotPos
			}
		}
	}

	return Pos{}
}

func valid(grid [][]string, i, j int) bool {
	m, n := len(grid), len(grid[0])
	return 0 <= i && i < m && 0 <= j && j < n
}

func moveBox(grid *[][]string, pos *Pos, move string) {
	d := DIRECTIONS[move]
	nx, ny := (*pos).x+d.x, (*pos).y+d.y
	count := 1

	// find the colling box to move together
	for (*grid)[nx][ny] == BOX {
		nx += d.x
		ny += d.y
		count++
	}

	// if in the end collide with wall, then can't move
	if (*grid)[nx][ny] == WALL {
		return
	}

	// move all the robot and the box
	for count > 0 {
		(*grid)[nx][ny], (*grid)[nx-d.x][ny-d.y] = (*grid)[nx-d.x][ny-d.y], (*grid)[nx][ny]
		nx -= d.x
		ny -= d.y
		count--
	}

	(*pos).x += d.x
	(*pos).y += d.y
}

func moveRobot(grid *[][]string, pos *Pos, move string) {
	d := DIRECTIONS[move]
	nx, ny := (*pos).x+d.x, (*pos).y+d.y
	valid := valid(*grid, nx, ny)
	next := (*grid)[nx][ny]

	// if found wall, cant move
	if valid && next == WALL {
		return
	}

	if valid && next == BOX {
		moveBox(grid, pos, move)
		return
	}

	(*grid)[nx][ny] = ROBOT
	(*grid)[(*pos).x][(*pos).y] = EMPTY
	(*pos).x += d.x
	(*pos).y += d.y
}

func countBox(grid [][]string) int {
	m, n := len(grid), len(grid[0])
	count := 0

	for i := range m {
		for j := range n {
			if grid[i][j] == BOX {
				count += i*POINT + j
			}
		}
	}

	return count
}

func solve(grid [][]string, moves []string) int {
	pos := findRobotPos(grid)

	for _, move := range moves {
		moveRobot(&grid, &pos, move)
	}

	return countBox(grid)
}

func main() {
	start := time.Now()
	grid, moves := parseInput("../input.txt")
	puzzle_1 := solve(grid, moves)

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", puzzle_1)
}
