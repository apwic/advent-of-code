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
	pos := make(map[int]int)

	j := 0
	for i, num := range nums {
		pos[i] = j
		j += num
		// files if index if even
		// free space if odd
		if i%2 == 0 {
			for range num {
				disk = append(disk, file_idx)
				disk2 = append(disk2, file_idx)
			}
			file_idx++
		} else {
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
	id := 1
	for j := len(nums) - 1; j >= 0; j -= 2 {
		for k := id; k <= j; k += 2 {
			// space is not enough
			if nums[k] < nums[j] {
				continue
			}

			// found the space
			// fill the empty space with file
			i := pos[k]
			for l := 0; l < nums[j]; l++ {
				disk2[i] = j / 2
				i++
			}

			// fill with empty space
			i = pos[j]
			for l := 0; l < nums[j]; l++ {
				disk2[i] = -1
				i++
			}

			nums[k] -= nums[j]
			pos[k] += nums[j]
			nums[j] = 0

			if nums[id] == -1 {
				id += 2
			}
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
