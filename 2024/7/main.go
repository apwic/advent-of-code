package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func digits(num int) int {
	if num == 0 {
		return 1
	}

	count := 0
	for num != 0 {
		num /= 10
		count++
	}

	return count
}

func backtrack(nums []int, idx int, curr int, target int) bool {
	if idx == len(nums) {
		return curr == target
	}

	add := curr + nums[idx]
	mul := curr * nums[idx]

	return backtrack(nums, idx+1, add, target) || backtrack(nums, idx+1, mul, target)
}

func backtrack2(nums []int, idx int, curr int, target int) bool {
	if idx == len(nums) {
		return curr == target
	}

	add := curr + nums[idx]
	mul := curr * nums[idx]
	comb := curr*(int(math.Pow(10, float64(digits(nums[idx]))))) + nums[idx]

	return backtrack2(nums, idx+1, add, target) || backtrack2(nums, idx+1, mul, target) || backtrack2(nums, idx+1, comb, target)
}

func main() {
	start := time.Now()
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
	puzzle_2 := 0
	for idx := range grid {
		nums := grid[idx]
		target := targets[idx]

		if backtrack(nums, 1, nums[0], target) {
			puzzle_1 += target
		}

		if backtrack2(nums, 1, nums[0], target) {
			puzzle_2 += target
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Time elapsed :", elapsed)

	fmt.Println("puzzle 1", puzzle_1)
	fmt.Println("puzzle 2", puzzle_2)
}
