package main

import (
	"fmt"
	"math"
	"time"
)

var (
	test     = []int{}
	input    = []int{}
	PUZZLE_1 = 25
	PUZZLE_2 = 75 - PUZZLE_1
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

func rules(num int) []int {
	if num == 0 {
		return []int{1}
	}

	if digits(num)%2 == 0 {
		digit := digits(num)
		mult := int(math.Pow10(digit / 2))
		a := num / mult
		b := num % mult

		return []int{a, b}
	}

	return []int{num * 2024}
}

func process(stones *map[int]int) {
	new_stones := make(map[int]int)

	for key, val := range *stones {
		nums := rules(key)

		for _, num := range nums {
			if _, exist := new_stones[num]; exist {
				new_stones[num] += val
			} else {
				new_stones[num] = val
			}
		}
	}

	*stones = new_stones
}

func puzzle(stones *map[int]int, count int) int {
	for range count {
		process(stones)
	}

	result := 0
	for _, val := range *stones {
		result += val
	}

	return result
}

func main() {
	start := time.Now()

	arr := input
	stones := make(map[int]int)

	for _, num := range arr {
		stones[num] = 1
	}

	fmt.Println("puzzle 1: ", puzzle(&stones, PUZZLE_1))
	fmt.Println("puzzle 2: ", puzzle(&stones, PUZZLE_2))

	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
}
