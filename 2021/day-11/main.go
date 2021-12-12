package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N int = 10

func printArea(a [][]int) {
	for _, r := range a {
		for _, c := range r {
			fmt.Printf("%c", c+0x30)
		}
		fmt.Println()
	}
}

func sumArea(a [][]int) int {
	n := 0
	for _, r := range a {
		for _, c := range r {
			n += c
		}
	}
	return n
}

func simulate(octoArea [][]int) int {
	// Init mask.
	isExploding := make([][]int, N)
	for i := 0; i < N; i++ {
		row := make([]int, N)
		for j := 0; j < N; j++ {
			row[j] = 0
		}
		isExploding[i] = row
	}

	// Step 1.
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			octoArea[i][j]++
		}
	}

	// Step 2.
	nFlash := 0
	for {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if octoArea[i][j] > 9 && isExploding[i][j] == 0 {
					isExploding[i][j] = 1
					if i > 0 {
						octoArea[i-1][j]++ // N.
					}
					if i > 0 && j < (N-1) {
						octoArea[i-1][j+1]++ // NE.
					}
					if j < (N - 1) {
						octoArea[i][j+1]++ // E.
					}
					if i < (N-1) && j < (N-1) {
						octoArea[i+1][j+1]++ // SE.
					}
					if i < (N - 1) {
						octoArea[i+1][j]++ // S.
					}
					if i < (N-1) && j > 0 {
						octoArea[i+1][j-1]++ // SW.
					}
					if j > 0 {
						octoArea[i][j-1]++ // W.
					}
					if i > 0 && j > 0 {
						octoArea[i-1][j-1]++ // NW.
					}
				}
			}
		}
		if sumArea(isExploding) == nFlash {
			break
		}
		nFlash = sumArea(isExploding)
	}

	// Step 3.
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if octoArea[i][j] > 9 {
				octoArea[i][j] = 0
			}
		}
	}

	return nFlash
}

func parse(data []string) [][]int {
	octoArea := make([][]int, N)

	for i := 0; i < N; i++ {
		octoRow := make([]int, N)
		for j, token := range data[i] {
			n, _ := strconv.Atoi(string(token))
			octoRow[j] = n
		}
		octoArea[i] = octoRow
	}

	return octoArea
}

func part1(octoArea [][]int) int {
	total := 0
	for k := 0; k < 100; k++ {
		total += simulate(octoArea)
	}
	return total
}

func part2(octoArea [][]int) int {
	k := 0
	for {
		if simulate(octoArea) == 100 {
			break
		}
		k++
	}
	return k + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	octoArea := parse(data)
	fmt.Println(part1(octoArea))
	octoArea = parse(data)
	fmt.Println(part2(octoArea))
}
