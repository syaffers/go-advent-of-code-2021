package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(data []string) int {
	x, y := 0, 0

	for i := 0; i < len(data); i++ {
		move_tok := strings.Split(data[i], " ")
		move := move_tok[0]
		value, _ := strconv.Atoi(move_tok[1])

		switch move {
		case "forward":
			x += value
		case "up":
			y -= value
		case "down":
			y += value
		}
	}
	return x * y
}

func part2(data []string) int {
	x, y, a := 0, 0, 0

	for i := 0; i < len(data); i++ {
		move_tok := strings.Split(data[i], " ")
		move := move_tok[0]
		value, _ := strconv.Atoi(move_tok[1])

		switch move {
		case "forward":
			x += value
			y += value * a
		case "up":
			a -= value
		case "down":
			a += value
		}
	}
	return x * y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}
