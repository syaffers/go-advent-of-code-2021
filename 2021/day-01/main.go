package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(data []string) int {
	c, m := -1, -1

	for i := 0; i < len(data); i++ {
		n, _ := strconv.Atoi(data[i])
		if n > m {
			c++
		}
		m = n
	}

	return c
}

func part2(data []string) int {
	c, m := -1, -1

	for i := 0; i < len(data)-2; i++ {
		x, _ := strconv.Atoi(data[i])
		y, _ := strconv.Atoi(data[i+1])
		z, _ := strconv.Atoi(data[i+2])
		if (x + y + z) > m {
			c++
		}
		m = x + y + z
	}

	return c
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
