package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	var first, second []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Fields(line)

		num, err := strconv.Atoi(part[0])
		if err != nil {
			log.Println(err)
			return
		}

		first = append(first, num)

		num, err = strconv.Atoi(part[1])
		if err != nil {
			log.Println(err)
			return
		}

		second = append(second, num)
	}

	sort.Slice(first, func(i, j int) bool {
		return first[i] < first[j]
	})

	sort.Slice(second, func(i, j int) bool {
		return second[i] < second[j]
	})

	ans := 0

	for i := 0; i < len(first); i++ {
		distance := first[i] - second[i]

		if distance < 0 {
			distance *= -1
		}

		ans += distance
	}

	fmt.Println(ans)
}
