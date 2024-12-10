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

func delete(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	start := time.Now()
	nums, err := parseInput("input.txt")
	if err != nil {
		return
	}

	disk := make([]int, 0)
	disk2 := make([]int, 0)
	file_idx := 0
	empty_space := make([]int, 0)
	filled_space := make([]int, 0)

	for i, num := range nums {
		// files if index if even
		// free space if odd
		if i%2 == 0 {
			filled_space = append(filled_space, num)
			for range num {
				disk = append(disk, file_idx)
				disk2 = append(disk2, file_idx)
			}
			file_idx++
		} else {
			empty_space = append(empty_space, num)
			for range num {
				disk = append(disk, -1)
				disk2 = append(disk2, -1)
			}
		}
	}

	empty_space_idx := make([]int, 0)
	filled_space_idx := make([]int, 0)
	for i, space := range disk {
		if space == -1 {
			empty_space_idx = append(empty_space_idx, i)
		} else {
			filled_space_idx = append(filled_space_idx, i)
		}
	}

	// puzzle 1
	filled_idx := len(disk) - 1
	empty_idx := 0
	for avail := len(empty_space_idx); avail > 0; avail-- {
		if disk[filled_idx] == -1 {
			filled_idx--
			continue
		}

		disk[filled_idx], disk[empty_space_idx[empty_idx]] = disk[empty_space_idx[empty_idx]], disk[filled_idx]
		filled_idx--
		empty_idx++
	}

	// puzzle 2
	// puzzle 2
	filled_idx = len(filled_space_idx) - 1
	for i := len(filled_space) - 1; i >= 0; i-- {
		// find empty space from the beginning
		empty_idx = 0
		swap := false
		for j, space := range empty_space {
			// no space, find next block
			if space < filled_space[i] {
				empty_idx += space
				continue
			}

			// empty index is ahead of the file blocks
			if empty_space_idx[empty_idx] >= filled_space_idx[filled_idx] {
				break
			}

			// there's space
			// do swap here
			swap = true
			for range filled_space[i] {
				disk2[empty_space_idx[empty_idx]], disk2[filled_space_idx[filled_idx]] = disk2[filled_space_idx[filled_idx]], disk2[empty_space_idx[empty_idx]]

				empty_space_idx = delete(empty_space_idx, empty_idx)
				filled_idx--
			}

			// reduce the empty space
			empty_space[j] -= filled_space[i]
			break
		}

		// align the filled_idx
		if !swap {
			filled_idx -= filled_space[i]
		}
	}

	puzzle_1 := 0
	puzzle_2 := 0
	for i := range disk {
		if disk[i] != -1 {
			puzzle_1 += i * disk[i]
		}

		if disk2[i] != -1 {
			puzzle_2 += i * disk2[i]
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Time elapsed:", elapsed)

	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Println("puzzle 2:", puzzle_2)
}
