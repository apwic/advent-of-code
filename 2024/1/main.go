package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var first, second []int
	first_map := map[int]int{}
	second_map := map[int]int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Fields(line)

		num, err := strconv.Atoi(part[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		first = append(first, num)
		val, exist := first_map[num]

		if exist {
			first_map[num] = val + 1
		} else {
			first_map[num] = 1
		}

		num, err = strconv.Atoi(part[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		second = append(second, num)
		val, exist = second_map[num]

		if exist {
			second_map[num] = val + 1
		} else {
			second_map[num] = 1
		}
	}

	sort.Slice(first, func(i, j int) bool {
		return first[i] < first[j]
	})

	sort.Slice(second, func(i, j int) bool {
		return second[i] < second[j]
	})

	puzzle_1 := 0

	for i := 0; i < len(first); i++ {
		distance := first[i] - second[i]

		if distance < 0 {
			distance *= -1
		}

		puzzle_1 += distance
	}

	fmt.Println("Puzzle 1: ", puzzle_1)

	puzzle_2 := 0

	for key, val := range first_map {
		puzzle_2 += key * val * second_map[key]
	}

	fmt.Println("Puzzle 2: ", puzzle_2)

	return
}
