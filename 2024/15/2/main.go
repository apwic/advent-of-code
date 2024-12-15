package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	WALL  = "#"
	ROBOT = "@"
	BOX   = "O"
	EMPTY = "."

	BOX_START = "["
	BOX_END   = "]"

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

type Box struct {
	start Pos
	end   Pos
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
			temp := make([]string, 0)
			for _, l := range line {
				if l == WALL {
					temp = append(temp, WALL)
					temp = append(temp, WALL)
				}

				if l == EMPTY {
					temp = append(temp, EMPTY)
					temp = append(temp, EMPTY)
				}

				if l == BOX {
					temp = append(temp, BOX_START)
					temp = append(temp, BOX_END)
				}

				if l == ROBOT {
					temp = append(temp, ROBOT)
					temp = append(temp, EMPTY)
				}
			}
			grid = append(grid, temp)
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

func boxPair(grid [][]string, x, y int) Box {
	if grid[x][y] == BOX_START {
		return Box{
			start: Pos{x: x, y: y},
			end:   Pos{x: x, y: y + 1},
		}
	}

	return Box{
		start: Pos{x: x, y: y - 1},
		end:   Pos{x: x, y: y},
	}
}

func nextBoxVert(grid [][]string, box Box, move string) []Box {
	d := DIRECTIONS[move]
	start_dx, start_dy := box.start.x+d.x, box.start.y+d.y
	end_dx, end_dy := box.end.x+d.x, box.end.y+d.y
	boxes := []Box{}

	// found wall
	if grid[start_dx][start_dy] == WALL || grid[end_dx][end_dy] == WALL {
		return nil
	}

	// box align
	if grid[start_dx][start_dy] == BOX_START &&
		grid[end_dx][end_dy] == BOX_END {
		boxes = append(boxes, boxPair(grid, start_dx, start_dy))
	}

	// misalign
	if grid[start_dx][start_dy] == BOX_END {
		boxes = append(boxes, boxPair(grid, start_dx, start_dy))
	}

	if grid[end_dx][end_dy] == BOX_START {
		boxes = append(boxes, boxPair(grid, end_dx, end_dy))
	}

	return boxes
}

func findBoxVert(grid [][]string, pos Pos, move string) [][]Box {
	d := DIRECTIONS[move]
	nx, ny := pos.x+d.x, pos.y+d.y
	boxes := make([][]Box, 0)
	boxes = append(boxes, []Box{boxPair(grid, nx, ny)})

	level := 0
	for len(boxes[level]) > 0 {
		temp := []Box{}
		for _, box := range boxes[level] {
			nextBoxes := nextBoxVert(grid, box, move)

			// found wall
			if nextBoxes == nil {
				return nil
			}

			temp = append(temp, nextBoxes...)
		}

		if len(temp) == 0 {
			break
		}

		level += 1
		boxes = append(boxes, temp)
	}

	return boxes
}

func moveBoxVert(grid *[][]string, pos *Pos, move string) {
	d := DIRECTIONS[move]
	boxes := findBoxVert(*grid, *pos, move)

	// found wall
	if boxes == nil {
		return
	}

	level := len(boxes) - 1
	for level >= 0 {
		for _, box := range boxes[level] {
			(*grid)[box.start.x+d.x][box.start.y+d.y] = BOX_START
			(*grid)[box.start.x][box.start.y] = EMPTY
			(*grid)[box.end.x+d.x][box.end.y+d.y] = BOX_END
			(*grid)[box.end.x][box.end.y] = EMPTY
		}
		level--
	}

	(*grid)[pos.x+d.x][pos.y+d.y] = ROBOT
	(*grid)[pos.x][pos.y] = EMPTY
	(*pos).x += d.x
	(*pos).y += d.y
}

func findBoxHor(grid [][]string, pos Pos, move string) []Box {
	d := DIRECTIONS[move]
	nx, ny := pos.x+d.x, pos.y+d.y
	boxes := []Box{}

	for grid[nx][ny] == BOX_START || grid[nx][ny] == BOX_END {
		boxes = append(boxes, boxPair(grid, nx, ny))
		nx += 2 * d.x
		ny += 2 * d.y
	}

	if grid[nx][ny] == WALL {
		return nil
	}

	return boxes
}

func moveBoxHor(grid *[][]string, pos *Pos, move string) {
	d := DIRECTIONS[move]
	boxes := findBoxHor(*grid, *pos, move)

	// found wall
	if boxes == nil {
		return
	}

	// move all the robot and the box
	count := len(boxes)*2 + 1
	nx, ny := (*pos).x+count*d.x, (*pos).y+count*d.y
	for range count {
		(*grid)[nx][ny], (*grid)[nx-d.x][ny-d.y] = (*grid)[nx-d.x][ny-d.y], (*grid)[nx][ny]
		nx -= d.x
		ny -= d.y
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

	if valid && (next == BOX_START || next == BOX_END) {
		if move == "^" || move == "v" {
			moveBoxVert(grid, pos, move)
		} else {
			moveBoxHor(grid, pos, move)
		}
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
			if grid[i][j] == BOX_START {
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

	puzzle_2 := solve(grid, moves)

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 2:", puzzle_2)
}
