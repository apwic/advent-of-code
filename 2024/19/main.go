package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseInput(fileName string) ([]string, []string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	patterns := []string{}
	towels := []string{}
	i := 0

	for scanner.Scan() {
		if i == 0 {
			line := strings.Split(scanner.Text(), ",")

			for _, l := range line {
				patterns = append(patterns, strings.TrimSpace(l))
			}
		} else if i > 1 {
			towels = append(towels, strings.TrimSpace(scanner.Text()))
		}

		i++
	}

	return patterns, towels
}

func DP(towel string, patterns []string, cache *map[string]int) int {
	if _, exist := (*cache)[towel]; !exist {
		if len(towel) == 0 {
			return 1
		}

		count := 0
		for _, pattern := range patterns {
			if strings.HasPrefix(towel, pattern) {
				count += DP(towel[len(pattern):], patterns, cache)
			}
		}
		(*cache)[towel] += count
	}

	return (*cache)[towel]
}

func solve(fileName string) {
	start := time.Now()
	patterns, towels := parseInput(fileName)
	cache := make(map[string]int)
	puzzle_1, puzzle_2 := 0, 0

	for _, towel := range towels {
		count := DP(towel, patterns, &cache)
		if count > 0 {
			puzzle_1++
		}
		puzzle_2 += count
	}

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Println("puzzle 2:", puzzle_2)
}

func main() {
	solve("input.txt")
}
