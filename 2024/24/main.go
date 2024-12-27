package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILENAME = "input.txt"

	AND = "AND"
	OR  = "OR"
	XOR = "XOR"
	Z   = 'z'
)

var (
	wires = make(map[string]int)
	gates = make([][]string, 0)
)

func parseInput() {
	file, _ := os.Open(FILENAME)
	scanner := bufio.NewScanner(file)

	isWire := true

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			isWire = false
			continue
		}

		if isWire {
			wireText := strings.Split(text, " ")
			wireVal, _ := strconv.Atoi(wireText[1])
			wires[string(wireText[0][:3])] = wireVal
		} else {
			gateText := strings.Split(text, " ")
			gate := gateText[:3]
			gate = append(gate, gateText[4:]...)
			gates = append(gates, gate)
		}
	}
}

func wireExist(wire string) bool {
	_, exist := wires[wire]
	return exist
}

// process all the gate with queue system
// fill the wires map
// return maximum bit for the z value
func logicGate() int {
	idx := 0
	queue := make([][]string, 0)
	queue = append(queue, gates[idx])
	maxBit := 0

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		a, op, b, target := q[0], q[1], q[2], q[3]

		// add next gate to queue first
		idx++
		if idx < len(gates) {
			queue = append(queue, gates[idx])
		}

		if wireExist(a) && wireExist(b) {
			// process the gate
			switch op {
			case AND:
				wires[target] = wires[a] & wires[b]
			case OR:
				wires[target] = wires[a] | wires[b]
			case XOR:
				wires[target] = wires[a] ^ wires[b]
			}

			if target[0] == Z {
				bit, _ := strconv.Atoi(string(target[1:]))
				maxBit = max(maxBit, bit)
			}
		} else {
			// place the unknown wire to the back of the queue
			queue = append(queue, q)
		}
	}

	return maxBit
}

func zValue(maxBit int) int {
	val := 0
	for i := range maxBit + 1 {
		z := fmt.Sprintf("z%02d", i)
		val += wires[z] << i
	}

	return val
}

func solve() {
	maxBit := logicGate()
	puzzle_1 := zValue(maxBit)

	fmt.Println("puzzle 1:", puzzle_1)
}

func main() {
	parseInput()
	solve()
}
