package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	// x and y reversed for easier debugging
	GRID_X = 103
	GRID_Y = 101
	TIME   = 100
)

type Pos struct {
	x int
	y int
}

type Robot struct {
	pos   Pos
	speed Pos
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func printMatr(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(fileName string) ([]Robot, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	re := regexp.MustCompile(`-?\d+`)
	scanner := bufio.NewScanner(file)
	robots := make([]Robot, 0)

	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		nums := make([]int, 0)

		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			nums = append(nums, num)
		}

		robots = append(robots, Robot{
			// x and y are reversed
			pos:   Pos{x: nums[1], y: nums[0]},
			speed: Pos{x: nums[3], y: nums[2]},
		})
	}

	return robots, nil
}

func place(robot Robot) (int, int) {
	x := mod(robot.pos.x+TIME*robot.speed.x, GRID_X)
	y := mod(robot.pos.y+TIME*robot.speed.y, GRID_Y)

	return x, y
}

func placeRobots(robots []Robot, grid *[][]int) {
	for _, robot := range robots {
		x, y := place(robot)
		(*grid)[x][y] += 1
	}
}

func findQuadrant(x, y int) int {
	x_half := GRID_X / 2
	y_half := GRID_Y / 2

	if x == x_half || y == y_half {
		return -1
	}

	if x < x_half && y < y_half {
		return 0
	} else if x > x_half && y < y_half {
		return 1
	} else if x < x_half && y > y_half {
		return 2
	} else {
		return 3
	}
}

func solve(robots []Robot, grid *[][]int) int {
	quadrant := []int{0, 0, 0, 0}

	placeRobots(robots, grid)

	for x := range GRID_X {
		for y := range GRID_Y {
			if (*grid)[x][y] == 0 {
				continue
			}

			q := findQuadrant(x, y)
			if q == -1 {
				continue
			}

			quadrant[q] += (*grid)[x][y]
		}
	}

	ans := 1
	for _, q := range quadrant {
		ans *= q
	}

	return ans
}

func main() {
	start := time.Now()
	robots, err := parseInput("input.txt")
	if err != nil {
		return
	}

	grid := make([][]int, 0)
	for range GRID_X {
		temp := make([]int, 0)
		for range GRID_Y {
			temp = append(temp, 0)
		}

		grid = append(grid, temp)
	}

	puzzle_1 := solve(robots, &grid)
	fmt.Println("time elapsed:", time.Since(start))

	fmt.Println("puzzle 1: ", puzzle_1)
}
