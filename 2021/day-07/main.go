package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func parse(data []string) []int {
	crabs := make([]int, 0)
	for _, token := range strings.Split(data[0], ",") {
		currPos, _ := strconv.Atoi(token)
		crabs = append(crabs, currPos)
	}

	return crabs
}

func part1(crabs []int) int {
	minPos, maxPos, minFuel := 1<<31, 0, 1<<31

	for _, crab := range crabs {
		if crab < minPos {
			minPos = crab
		}
		if crab > maxPos {
			maxPos = crab
		}
	}

	for pos := minPos; pos <= maxPos; pos++ {
		fuel := 0
		for _, crab := range crabs {
			fuel += abs(crab - pos)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func part2(crabs []int) int {
	minPos, maxPos, minFuel, n := 1<<31, 0, 1<<31, 0

	for _, crab := range crabs {
		if crab < minPos {
			minPos = crab
		}
		if crab > maxPos {
			maxPos = crab
		}
	}

	for pos := minPos; pos <= maxPos; pos++ {
		fuel := 0
		for _, crab := range crabs {
			n = abs(crab - pos)
			fuel += ((n + 1) * n) / 2
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	crabs := parse(data)
	fmt.Println(part1(crabs))
	crabs = parse(data)
	fmt.Println(part2(crabs))
}
