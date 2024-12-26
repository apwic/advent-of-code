package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	FILENAME = "input.txt"
	DEPTH    = 2000

	PRUNE = 1 << 24
)

var (
	secrets = []int{}
	cache   = map[int]int{}
	// map from sequence, to secret, to the val
	sequences = map[string]map[int]int{}
)

func parseInput(fileName string) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		secrets = append(secrets, num)
	}
}

func arrIntToArrString(nums []int) []string {
	nums_str := make([]string, len(nums))

	for i, num := range nums {
		nums_str[i] = strconv.Itoa(num)
	}

	return nums_str
}

func lastDigit(num int) int {
	str_num := strconv.Itoa(num)
	num, _ = strconv.Atoi(string(str_num[len(str_num)-1]))
	return num
}

func secretMult(secret int, num int) int {
	return ((secret * num) ^ secret) % PRUNE
}

func secretDiv(secret int, num int) int {
	return ((secret / num) ^ secret) % PRUNE
}

func addToCache(sequence *[]int, secret, digit int) {
	key := fmt.Sprintf("%d,%d,%d,%d", (*sequence)[0], (*sequence)[1], (*sequence)[2], (*sequence)[3])

	if _, exist := sequences[key]; exist {
		if _, exist := sequences[key][secret]; !exist {
			sequences[key][secret] = digit
		}
	} else {
		sequences[key] = map[int]int{secret: digit}
	}
}

func produceSecret(original, secret, depth, prevDigit int, sequence *[]int) int {
	if depth == DEPTH {
		return secret
	}

	var next int

	if val, exist := cache[secret]; exist {
		next = val
	} else {

		next = secretMult(secret, 1<<6)
		next = secretDiv(next, 1<<5)
		next = secretMult(next, 1<<11)
	}

	lastDigit := lastDigit(next)
	(*sequence) = append((*sequence), lastDigit-prevDigit)
	if depth >= 4 {
		(*sequence) = (*sequence)[1:]

		if depth < DEPTH {
			addToCache(sequence, original, lastDigit)
		}
	}

	cache[secret] = next

	return produceSecret(original, next, depth+1, lastDigit, sequence)
}

func solve() {
	start := time.Now()

	puzzle_1 := 0
	for _, secret := range secrets {
		sequence := []int{}
		ans := produceSecret(secret, secret, 0, 0, &sequence)

		puzzle_1 += ans
	}

	puzzle_2 := 0
	for key := range sequences {
		curr := 0
		for _, val := range sequences[key] {
			curr += val
		}
		puzzle_2 = max(puzzle_2, curr)
	}

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Println("puzzle 2:", puzzle_2)
}

func main() {
	parseInput(FILENAME)
	solve()
}
