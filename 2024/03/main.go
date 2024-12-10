package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mul_regex, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	num_regex, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		fmt.Println(err)
		return
	}

	dont_regex, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	do_regex, err := regexp.Compile(`do\(\)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	puzzle_1 := 0
	puzzle_2 := 0
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()

		mul := mul_regex.FindAllStringIndex(line, -1)
		do := do_regex.FindAllStringIndex(line, -1)
		dont := dont_regex.FindAllStringIndex(line, -1)

		all := append(append(mul, do...), dont...)
		sort.Slice(all, func(i, j int) bool {
			return all[i][0] < all[j][0]
		})

		for _, match := range all {
			exp := line[match[0]:match[1]]

			if exp == "do()" {
				enabled = true
			} else if exp == "don't()" {
				enabled = false
			} else if mul_regex.MatchString(exp) {
				nums := num_regex.FindAllString(exp, -1)

				a, err := strconv.Atoi(nums[0])
				if err != nil {
					fmt.Println(err)
					return
				}

				b, err := strconv.Atoi(nums[1])
				if err != nil {
					fmt.Println(err)
					return
				}

				res := a * b
				puzzle_1 += res
				if enabled {
					puzzle_2 += res
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)

	fmt.Println("Puzzle 1: ", puzzle_1)
	fmt.Println("Puzzle 2: ", puzzle_2)
}
