package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
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

	puzzle_1 := 0

	for scanner.Scan() {
		line := scanner.Text()

		mul := mul_regex.FindAllString(line, -1)

		for _, exp := range mul {
			num_match := num_regex.FindAllString(exp, -1)
			a, err := strconv.Atoi(num_match[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			b, err := strconv.Atoi(num_match[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			puzzle_1 += a * b
		}
	}

	fmt.Println("Puzzle 1", puzzle_1)
}
