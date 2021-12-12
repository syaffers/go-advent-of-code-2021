package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Globals.
var pair map[string]string

func isOpener(sym string) bool {
	return sym == "(" || sym == "[" || sym == "{" || sym == "<"
}

func checkValidity(line string) (bool, string, []string) {
	sym := ""
	exit := false
	stack := make([]string, 0)
	for _, r := range line {
		chunk := ""
		sym = string(r)
		if isOpener(sym) {
			stack = append(stack, sym)
		} else {
			for i := len(stack) - 1; i >= 0; i-- {
				if isOpener(stack[i]) {
					if stack[i] == pair[sym] {
						chunk = stack[i] + chunk
						stack = stack[0:i]
						stack = append(stack, chunk+sym)
						break
					} else {
						exit = true
					}
				} else {
					chunk = stack[i] + chunk
				}
			}
			if exit {
				break
			}
		}
	}
	return !exit, sym, stack
}

func parse(data []string) []string {
	return data
}

func part1(lines []string) int {
	total := 0
	score := make(map[string]int)

	score[")"] = 3
	score["]"] = 57
	score["}"] = 1197
	score[">"] = 25137

	for _, line := range lines {
		isValid, sym, _ := checkValidity(line)
		if !isValid {
			total += score[sym]
		}
	}
	return total
}

func part2(lines []string) int {
	// total := 0
	lineScores := make([]int, 0)
	score := make(map[string]int)

	score["("] = 1
	score["["] = 2
	score["{"] = 3
	score["<"] = 4

	for _, line := range lines {
		lineScore := 0
		isValid, _, stack := checkValidity(line)
		if isValid {
			for i := len(stack) - 1; i >= 0; i-- {
				if len(stack[i]) < 2 {
					lineScore *= 5
					lineScore += score[stack[i]]
				}
			}
			lineScores = append(lineScores, lineScore)
		}
	}
	sort.Ints(lineScores)
	return lineScores[len(lineScores)/2]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}

	pair = make(map[string]string)
	pair[")"] = "("
	pair["]"] = "["
	pair["}"] = "{"
	pair[">"] = "<"

	lines := parse(data)
	fmt.Println(part1(lines))
	lines = parse(data)
	fmt.Println(part2(lines))
}
