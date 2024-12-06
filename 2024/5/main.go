package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	rules := make(map[int]map[int]bool)
	var pages [][]int

	if_rules := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if_rules = false
			continue
		}

		if if_rules {
			rule := strings.Split(line, "|")
			a, err := strconv.Atoi(rule[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			b, err := strconv.Atoi(rule[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			if _, exist := rules[a]; !exist {
				rules[a] = make(map[int]bool)
			}

			rules[a][b] = true
		} else {
			page := strings.Split(line, ",")
			temp := make([]int, 0)
			for _, p := range page {
				num, err := strconv.Atoi(p)
				if err != nil {
					fmt.Println(err)
					return
				}

				temp = append(temp, num)
			}

			pages = append(pages, temp)
		}
	}

	puzzle_1 := 0
	puzzle_2 := 0
	for _, page := range pages {
		order := true

		// puzzle 1
		// check if any of the order of the element is inscorrect
		// by checking the map of page[j] if page[i] exist then
		// it is backward (the order is wrong)
		for i := 0; i < len(page); i++ {
			for j := i + 1; j < len(page); j++ {
				if _, exist := rules[page[j]][page[i]]; exist {
					order = false
					break
				}
			}

			if !order {
				break
			}
		}

		if order {
			puzzle_1 += page[(len(page) / 2)]
			continue
		}

		// puzzle 2
		// use bubble sort since we can check the hashmap
		for i := 0; i < len(page); i++ {
			swap := false
			for j := 0; j < len(page)-i-1; j++ {
				if _, exist := rules[page[j+1]][page[j]]; exist {
					page[j], page[j+1] = page[j+1], page[j]
					swap = true
				}
			}

			if !swap {
				break
			}
		}

		puzzle_2 += page[len(page)/2]
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)

	fmt.Println("puzzle 1: ", puzzle_1)
	fmt.Println("puzzle 2: ", puzzle_2)
}
