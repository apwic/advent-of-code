package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func backtrack(nums []int, idx int, curr int, target int) bool {
	if idx == len(nums) {
		return curr == target
	}

	add := curr + nums[idx]
	mul := curr * nums[idx]

	return backtrack(nums, idx+1, add, target) || backtrack(nums, idx+1, mul, target)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var grid [][]int
	targets := make([]int, 0)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var target int
		nums := make([]int, 0)

		for i, str_num := range line {
			if i == 0 {
				num, err := strconv.Atoi(str_num[:len(str_num)-1])
				if err != nil {
					fmt.Println(err)
					return
				}

				target = num
				continue
			}

			num, err := strconv.Atoi(str_num)
			if err != nil {
				fmt.Println(err)
				return
			}

			nums = append(nums, num)
		}

		grid = append(grid, nums)
		targets = append(targets, target)
	}

	puzzle_1 := 0
	for idx := range grid {
		nums := grid[idx]
		target := targets[idx]

		if backtrack(nums, 1, nums[0], target) {
			puzzle_1 += target
		}
	}

	fmt.Println("puzzle 1", puzzle_1)
}
