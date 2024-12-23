package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Move struct {
	dir    string
	target string
}

const (
	A_KEY = "A"

	UP    = "^"
	DOWN  = "v"
	LEFT  = "<"
	RIGHT = ">"

	ROBOTS_1 = 1
	ROBOTS_2 = 24
)

var NUM_SEQ = map[string]map[string][]string{
	A_KEY: {
		A_KEY: {},
		"0":   {LEFT},
		"1":   {UP, LEFT, LEFT},
		"2":   {LEFT, UP},
		"3":   {UP},
		"4":   {UP, UP, LEFT, LEFT},
		"5":   {LEFT, UP, UP},
		"6":   {UP, UP},
		"7":   {UP, UP, UP, LEFT, LEFT},
		"8":   {LEFT, UP, UP, UP},
		"9":   {UP, UP, UP},
	},
	"0": {
		A_KEY: {RIGHT},
		"0":   {},
		"1":   {UP, LEFT},
		"2":   {UP},
		"3":   {UP, RIGHT},
		"4":   {UP, UP, LEFT},
		"5":   {UP, UP},
		"6":   {UP, UP, RIGHT},
		"7":   {UP, UP, UP, LEFT},
		"8":   {UP, UP, UP},
		"9":   {UP, UP, UP, RIGHT},
	},
	"1": {
		A_KEY: {RIGHT, RIGHT, DOWN},
		"0":   {RIGHT, DOWN},
		"1":   {},
		"2":   {RIGHT},
		"3":   {RIGHT, RIGHT},
		"4":   {UP},
		"5":   {UP, RIGHT},
		"6":   {UP, RIGHT, RIGHT},
		"7":   {UP, UP},
		"8":   {UP, UP, RIGHT},
		"9":   {UP, UP, RIGHT, RIGHT},
	},
	"2": {
		A_KEY: {DOWN, RIGHT},
		"0":   {DOWN},
		"1":   {LEFT},
		"2":   {},
		"3":   {RIGHT},
		"4":   {LEFT, UP},
		"5":   {UP},
		"6":   {UP, RIGHT},
		"7":   {LEFT, UP, UP},
		"8":   {UP, UP},
		"9":   {UP, UP, RIGHT},
	},
	"3": {
		A_KEY: {DOWN},
		"0":   {LEFT, DOWN},
		"1":   {LEFT, LEFT},
		"2":   {LEFT},
		"3":   {},
		"4":   {LEFT, LEFT, UP},
		"5":   {UP, LEFT},
		"6":   {UP},
		"7":   {LEFT, LEFT, UP, UP},
		"8":   {LEFT, UP, UP},
		"9":   {UP, UP},
	},
	"4": {
		A_KEY: {RIGHT, RIGHT, DOWN, DOWN},
		"0":   {RIGHT, DOWN, DOWN},
		"1":   {DOWN},
		"2":   {DOWN, RIGHT},
		"3":   {DOWN, RIGHT, RIGHT},
		"4":   {},
		"5":   {RIGHT},
		"6":   {RIGHT, RIGHT},
		"7":   {UP},
		"8":   {RIGHT, UP},
		"9":   {RIGHT, RIGHT, UP},
	},
	"5": {
		A_KEY: {DOWN, DOWN, RIGHT},
		"0":   {DOWN, DOWN},
		"1":   {LEFT, DOWN},
		"2":   {DOWN},
		"3":   {DOWN, RIGHT},
		"4":   {LEFT},
		"5":   {},
		"6":   {RIGHT},
		"7":   {LEFT, UP},
		"8":   {UP},
		"9":   {UP, RIGHT},
	},
	"6": {
		A_KEY: {DOWN, DOWN},
		"0":   {LEFT, DOWN, DOWN},
		"1":   {LEFT, LEFT, DOWN},
		"2":   {LEFT, DOWN},
		"3":   {DOWN},
		"4":   {LEFT, LEFT},
		"5":   {LEFT},
		"6":   {},
		"7":   {LEFT, LEFT, UP},
		"8":   {LEFT, UP},
		"9":   {UP},
	},
	"7": {
		A_KEY: {RIGHT, RIGHT, DOWN, DOWN, DOWN},
		"0":   {RIGHT, DOWN, DOWN, DOWN},
		"1":   {DOWN, DOWN, DOWN},
		"2":   {DOWN, DOWN, RIGHT},
		"3":   {DOWN, DOWN, RIGHT, RIGHT},
		"4":   {DOWN},
		"5":   {DOWN, RIGHT},
		"6":   {DOWN, RIGHT, RIGHT},
		"7":   {},
		"8":   {RIGHT},
		"9":   {RIGHT, RIGHT},
	},
	"8": {
		A_KEY: {DOWN, DOWN, DOWN, RIGHT},
		"0":   {DOWN, DOWN, DOWN},
		"1":   {LEFT, DOWN, DOWN},
		"2":   {DOWN, DOWN},
		"3":   {DOWN, DOWN, RIGHT},
		"4":   {LEFT, DOWN},
		"5":   {DOWN},
		"6":   {DOWN, RIGHT},
		"7":   {LEFT},
		"8":   {},
		"9":   {RIGHT},
	},
	"9": {
		A_KEY: {DOWN, DOWN, DOWN},
		"0":   {LEFT, DOWN, DOWN, DOWN},
		"1":   {LEFT, LEFT, DOWN, DOWN},
		"2":   {LEFT, DOWN, DOWN},
		"3":   {DOWN, DOWN},
		"4":   {LEFT, LEFT, DOWN},
		"5":   {LEFT, DOWN},
		"6":   {DOWN},
		"7":   {LEFT, LEFT},
		"8":   {LEFT},
		"9":   {},
	},
}

var DIR_SEQ = map[string]map[string][]string{
	A_KEY: {
		A_KEY: {},
		UP:    {LEFT},
		LEFT:  {DOWN, LEFT, LEFT},
		RIGHT: {DOWN},
		DOWN:  {LEFT, DOWN},
	},
	UP: {
		A_KEY: {RIGHT},
		UP:    {},
		LEFT:  {DOWN, LEFT},
		RIGHT: {DOWN, RIGHT},
		DOWN:  {DOWN},
	},
	DOWN: {
		A_KEY: {UP, RIGHT},
		UP:    {UP},
		LEFT:  {LEFT},
		RIGHT: {RIGHT},
		DOWN:  {},
	},
	LEFT: {
		A_KEY: {RIGHT, RIGHT, UP},
		UP:    {RIGHT, UP},
		LEFT:  {},
		RIGHT: {RIGHT, RIGHT},
		DOWN:  {RIGHT},
	},
	RIGHT: {
		A_KEY: {UP},
		UP:    {LEFT, UP},
		LEFT:  {LEFT, LEFT},
		RIGHT: {},
		DOWN:  {LEFT},
	},
}

func parseInput(fileName string) []string {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	codes := []string{}

	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	return codes
}

func press(keypad map[string]map[string][]string, target string) string {
	path := []string{}
	currPad := A_KEY

	for _, targetPad := range target[:] {
		strTargetPad := string(targetPad)
		path = append(path, strings.Join(keypad[currPad][strTargetPad], ""))
		path = append(path, A_KEY)
		currPad = strTargetPad
	}

	return strings.Join(path, "")
}

func countComplexity(code string, length int) int {
	re := regexp.MustCompile(`\d+`)
	find := re.FindAllString(code, -1)
	num, _ := strconv.Atoi(find[0])

	return num * length
}

func solve(fileName string) {
	codes := parseInput(fileName)
	puzzle_1 := 0

	for _, code := range codes {
		numeric := press(NUM_SEQ, code)
		directional := press(DIR_SEQ, numeric)

		for range ROBOTS_1 {
			directional = press(DIR_SEQ, directional)
		}

		puzzle_1 += countComplexity(code, len(directional))
	}

	fmt.Println("puzzle 1:", puzzle_1)
}

func main() {
	solve("input.txt")
}
