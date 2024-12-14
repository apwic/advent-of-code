package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var (
	// x and y reversed for easier debugging
	GRID_X     = 103
	GRID_Y     = 101
	TIME       = 100
	START_TIME = 0
	END_TIME   = GRID_X * GRID_Y
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

func printMatrAny(grid [][]any) {
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

func place(robot Robot, time int) (int, int) {
	x := mod(robot.pos.x+time*robot.speed.x, GRID_X)
	y := mod(robot.pos.y+time*robot.speed.y, GRID_Y)

	return x, y
}

func placeRobots(robots []Robot, grid *[][]int) {
	for _, robot := range robots {
		x, y := place(robot, TIME)
		(*grid)[x][y] += 1
	}
}

func placeRobots2(robots []Robot) {
	outputDir := "grid_images"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	initializeGrid := func() [][]string {
		grid := make([][]string, GRID_X)
		for i := range grid {
			grid[i] = make([]string, GRID_Y)
			for j := range grid[i] {
				grid[i][j] = " "
			}
		}
		return grid
	}

	for time := START_TIME; time <= END_TIME; time++ {
		grid := initializeGrid()
		for _, robot := range robots {
			x, y := place(robot, time)
			grid[x][y] = "X"
		}

		// Create an image from the grid
		img := createImageFromGrid(grid)

		// Save the image to the directory
		fileName := filepath.Join(outputDir, fmt.Sprintf("%03d.png", time))
		err := saveImage(fileName, img)
		if err != nil {
			fmt.Println("Error saving image:", err)
			return
		}

		fmt.Println("Image saved:", fileName)
	}
}

func createImageFromGrid(grid [][]string) *image.RGBA {
	imgWidth := len(grid[0])
	imgHeight := len(grid)
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Define colors
	robotColor := color.RGBA{255, 255, 255, 255}
	emptyColor := color.RGBA{0, 0, 0, 255}

	for y, row := range grid {
		for x, cell := range row {
			if cell == "X" {
				img.Set(x, y, robotColor)
			} else {
				img.Set(x, y, emptyColor)
			}
		}
	}
	return img
}

func saveImage(fileName string, img *image.RGBA) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
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

	placeRobots2(robots)
}
