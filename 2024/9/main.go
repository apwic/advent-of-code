package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	nums := make([]int, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for _, ch := range line {
			num, err := strconv.Atoi(ch)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			nums = append(nums, num)
		}
	}

	return nums, err
}

func main() {
	start := time.Now()
	nums, err := parseInput("input.txt")
	if err != nil {
		return
	}

	disk := make([]int, 0)
	file_idx := 0

	for i, num := range nums {
		// files if index if even
		// free space if odd
		if i%2 == 0 {
			for range num {
				disk = append(disk, file_idx)
			}
			file_idx++
		} else {
			for range num {
				disk = append(disk, -1)
			}
		}
	}

	empty_space := make([]int, 0)
	filled_space := make([]int, 0)
	for i, space := range disk {
		if space == -1 {
			empty_space = append(empty_space, i)
		} else {
			filled_space = append(filled_space, i)
		}
	}

	filled_idx := len(disk) - 1
	empty_idx := 0
	for avail := len(empty_space); avail > 0; avail-- {
		if disk[filled_idx] == -1 {
			filled_idx--
			continue
		}

		disk[filled_idx], disk[empty_space[empty_idx]] = disk[empty_space[empty_idx]], disk[filled_idx]
		filled_idx--
		empty_idx++
	}

	puzzle_1 := 0
	for i, space := range disk {
		if space == -1 {
			break
		}

		puzzle_1 += i * space
	}

	elapsed := time.Since(start)
	fmt.Println("Time elapsed:", elapsed)

	fmt.Println("puzzle 1:", puzzle_1)
}
