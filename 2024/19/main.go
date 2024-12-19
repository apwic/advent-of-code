package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput(fileName string) (map[string]bool, []string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	patterns := map[string]bool{}
	towels := []string{}
	i := 0

	for scanner.Scan() {
		if i == 0 {
			line := strings.Split(scanner.Text(), ",")

			for _, l := range line {
				pattern := strings.TrimSpace(l)
				patterns[pattern] = true
			}
		} else if i > 1 {
			towels = append(towels, strings.TrimSpace(scanner.Text()))
		}

		i++
	}

	return patterns, towels
}

func DP(towel string, patterns map[string]bool) bool {
	n := len(towel)
	dp := make([]bool, n)

	for i := range n {
		for pattern := range patterns {
			n_p := len(pattern)

			// out of bounds
			if i < n_p-1 {
				continue
			}

			if i == n_p-1 || dp[i-n_p] {
				if towel[i-n_p+1:i+1] == pattern {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[n-1]
}

func solve(fileName string) {
	patterns, towels := parseInput(fileName)
	count := 0

	for _, towel := range towels {
		if DP(towel, patterns) {
			count++
		}
	}

	fmt.Println("puzzle 1:", count)
}

func main() {
	solve("input.txt")
}
