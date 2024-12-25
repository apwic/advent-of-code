package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	FILENAME = "input.txt"
	DEPTH    = 2000

	PRUNE = 1 << 24
)

var (
	secrets = []int{}
	cache   = map[int]int{}
)

func parseInput(fileName string) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		secrets = append(secrets, num)
	}
}

func secretMult(secret int, num int) int {
	return ((secret * num) ^ secret) % PRUNE
}

func secretDiv(secret int, num int) int {
	return ((secret / num) ^ secret) % PRUNE
}

func produceSecret(secret int, depth int) int {
	if depth == 0 {
		return secret
	}

	if val, exist := cache[secret]; exist {
		return produceSecret(val, depth-1)
	}

	next := secretMult(secret, 1<<6)
	next = secretDiv(next, 1<<5)
	next = secretMult(next, 1<<11)

	cache[secret] = next

	return produceSecret(next, depth-1)
}

func solve() {
	puzzle_1 := 0
	for _, secret := range secrets {
		ans := produceSecret(secret, DEPTH)
		puzzle_1 += ans
	}

	fmt.Println("puzzle 1:", puzzle_1)
}

func main() {
	parseInput(FILENAME)
	solve()
}
