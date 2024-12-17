package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Register struct {
	A int
	B int
	C int

	Pointer int
	Out     []int
	Program map[int]int
}

func parseInput(fileName string) Register {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d+`)
	register := Register{A: 0, B: 0, C: 0}
	programs := map[int]int{}

	line := 0
	for scanner.Scan() {
		if line < 3 {
			nums_string := re.FindAllString(scanner.Text(), -1)

			if line == 0 {
				num, _ := strconv.Atoi(nums_string[0])
				register.A = num
			} else if line == 1 {
				num, _ := strconv.Atoi(nums_string[0])
				register.B = num
			} else if line == 2 {
				num, _ := strconv.Atoi(nums_string[0])
				register.C = num
			}
		} else {
			text := scanner.Text()

			if len(text) == 0 {
				continue
			}

			nums_string := strings.Split(scanner.Text()[len("Program: "):], ",")

			for idx, num_str := range nums_string {
				num, _ := strconv.Atoi(num_str)
				programs[idx] = num
			}

			register.Program = programs
		}

		line++
	}

	return register
}

// 0
func (r *Register) adv(num int) {
	r.A >>= num
}

// 1
func (r *Register) bxl(num int) {
	r.B ^= num
}

// 2
func (r *Register) bst(num int) {
	r.B = num % 8
}

// 3
func (r *Register) jnz(num int) {
	if r.A != 0 {
		r.Pointer = num - 2
	}
}

// 4
func (r *Register) bxc() {
	r.B ^= r.C
}

// 5
func (r *Register) out(num int) {
	r.Out = append(r.Out, num%8)
}

// 6
func (r *Register) bdv(num int) {
	r.B = r.A >> num
}

// 7
func (r *Register) cdv(num int) {
	r.C = r.A >> num
}

func (r *Register) combo(num int) int {
	if 0 <= num && num <= 3 {
		return num
	} else if num == 4 {
		return r.A
	} else if num == 5 {
		return r.B
	} else if num == 6 {
		return r.C
	}

	return -1
}

func (r *Register) operate(op int, num int) {
	switch op {
	case 0:
		r.adv(r.combo(num))
	case 1:
		r.bxl(num)
	case 2:
		r.bst(r.combo(num))
	case 3:
		r.jnz(num)
	case 4:
		r.bxc()
	case 5:
		r.out(r.combo(num))
	case 6:
		r.bdv(r.combo(num))
	case 7:
		r.cdv(r.combo(num))
	}
	r.Pointer += 2
}

func (r *Register) Run() {
	for r.Pointer < len(r.Program) {

		a := r.Program[r.Pointer]
		b := r.Program[r.Pointer+1]
		r.operate(a, b)
	}
}

func (r *Register) String() string {
	var sb strings.Builder
	for _, num := range r.Out {
		sb.WriteString(strconv.Itoa(num))
		sb.WriteString(",")
	}

	return sb.String()[:len(sb.String())-1]
}

func Reverse(register Register) int {
	program := make([]int, len(register.Program))
	for i := 0; i < len(register.Program); i++ {
		program[i] = register.Program[i]
	}

	A := 0

	for n := len(program) - 1; n >= 0; n-- {
		A <<= 3
		mock := Register{
			A:       A,
			B:       register.B,
			C:       register.C,
			Program: register.Program,
			Pointer: 0,
			Out:     []int{},
		}
		mock.Run()

		for !slices.Equal(mock.Out, program[n:]) {
			A++
			mock.A = A
			mock.Out = []int{}
			mock.Pointer = 0
			mock.Run()
		}
	}

	return A
}

func Solve(fileName string) {
	start := time.Now()

	register := parseInput(fileName)
	register.Run()

	A := Reverse(register)

	fmt.Println("time elapsed:", time.Since(start))
	fmt.Println("puzzle 1:", register.String())
	fmt.Println("puzzle 2:", A)
}

func main() {
	Solve("input.txt")
}
