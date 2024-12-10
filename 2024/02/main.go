package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func pop(s []int, i int) []int {
	r := make([]int, 0)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}

func increasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if !(1 <= diff && diff <= 3) {
			return false
		}
	}

	return true
}

func decreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if !(1 <= diff && diff <= 3) {
			return false
		}
	}

	return true
}

func safe_removal(report []int) bool {
	for i := 0; i < len(report); i++ {
		removed := pop(report, i)

		if increasing(removed) || decreasing(removed) {
			return true
		}
	}

	return false
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	puzzle_1 := 0
	puzzle_2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Fields(line)

		var reports []int

		for _, str_num := range part {
			num, err := strconv.Atoi(str_num)
			if err != nil {
				fmt.Println(err)
				return
			}

			reports = append(reports, num)
		}

		if increasing(reports) || decreasing(reports) {
			puzzle_1++
		} else if safe_removal(reports) {
			puzzle_2++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)

	fmt.Println("Puzzle 1: ", puzzle_1)
	fmt.Println("Puzzle 2: ", puzzle_1+puzzle_2)
}
