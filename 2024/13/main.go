package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	MAX_SPIN = 10000000000000
	COST_A   = 3
	COST_B   = 1
	MAX_COST = MAX_SPIN*COST_A + MAX_SPIN*COST_B
	OFFSET   = 10000000000000
)

type Pos struct {
	x int
	y int
}

type Machine struct {
	a      Pos
	b      Pos
	target Pos
}

func parseInput(fileName string) ([]Machine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	reg, err := regexp.Compile(`[0-9]+`)

	machines := make([]Machine, 0)
	count := 0
	temp := make([]Pos, 0)
	for scanner.Scan() {
		if count == 3 {
			machine := Machine{
				a:      Pos{temp[0].x, temp[0].y},
				b:      Pos{temp[1].x, temp[1].y},
				target: Pos{temp[2].x, temp[2].y},
			}
			machines = append(machines, machine)
			temp = make([]Pos, 0)
			count = 0
			continue
		}

		num := reg.FindAllString(scanner.Text(), -1)
		a, _ := strconv.Atoi(num[0])
		b, _ := strconv.Atoi(num[1])
		temp = append(temp, Pos{x: a, y: b})
		count++
	}

	return machines, nil
}

// BRUTE FORCE
// too long for the part 2
func combination(machine Machine, offset int) int {
	machine.target.x += offset
	machine.target.y += offset
	ans := MAX_COST + 1

	for i := 0; i <= MAX_SPIN; i++ {
		if i*machine.a.x > machine.target.x {
			break
		}

		for j := 0; j <= MAX_SPIN; j++ {
			dist_x := i*machine.a.x + j*machine.b.x
			dist_y := i*machine.a.y + j*machine.b.y

			if dist_x > machine.target.x || dist_y > machine.target.y {
				continue
			}

			if dist_x == machine.target.x && dist_y == machine.target.y {
				ans = min(ans, int(i*COST_A+j*COST_B))
			}
		}
	}

	if ans > MAX_COST {
		return 0
	}

	return ans
}

// USING LINEAR ALGEBRA
// the linear system is represented as:
// a.x*a + b.x*b = target.x
// a.y*a + b.y*b = target.y
// use crammer rule for this

func solve(machine Machine, offset int) int {
	machine.target.x += offset
	machine.target.y += offset

	det := machine.a.x*machine.b.y - machine.b.x*machine.a.y

	if det == 0 {
		return 0
	}

	det_a := machine.target.x*machine.b.y - machine.target.y*machine.b.x
	det_b := machine.a.x*machine.target.y - machine.a.y*machine.target.x

	if det_a%det != 0 || det_b%det != 0 {
		return 0
	}

	a := det_a / det
	b := det_b / det

	return 3*a + b
}

func findCombination(machines []Machine) (int, int) {
	puzzle_1 := 0
	puzzle_2 := 0

	for _, machine := range machines {
		puzzle_1 += solve(machine, 0)
		puzzle_2 += solve(machine, OFFSET)
	}

	return puzzle_1, puzzle_2
}

func main() {
	start := time.Now()
	machines, err := parseInput("input.txt")
	if err != nil {
		return
	}

	puzzle_1, puzzle_2 := findCombination(machines)

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", puzzle_1)
	fmt.Println("puzzle 2:", puzzle_2)
}
