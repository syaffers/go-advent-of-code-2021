package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Heightmap struct {
	area   []int
	width  int
	height int
}

// Globals.
var visited map[int]bool

func floodCount(p int, hmap Heightmap) int {
	// Base case: we reach an edge.
	if hmap.area[p] == 9 {
		return 0
	}

	// Add current position to visited set.
	visited[p] = true

	// Recursive case.
	count := 0
	if p/hmap.width > 0 {
		// Look up.
		if _, ok := visited[p-hmap.width]; !ok {
			count += floodCount(p-hmap.width, hmap)
		}
	}
	if p/hmap.width < (hmap.height - 1) {
		// Look down.
		if _, ok := visited[p+hmap.width]; !ok {
			count += floodCount(p+hmap.width, hmap)
		}
	}
	if p%hmap.width > 0 {
		// Look left.
		if _, ok := visited[p-1]; !ok {
			count += floodCount(p-1, hmap)
		}
	}
	if p%hmap.width < (hmap.width - 1) {
		// Look right.
		if _, ok := visited[p+1]; !ok {
			count += floodCount(p+1, hmap)
		}
	}

	return 1 + count
}

func parse(data []string) Heightmap {
	width, height := 0, 0
	area := make([]int, 0)

	for _, line := range data {
		width = len(line)
		height++
		for _, token := range line {
			n, _ := strconv.Atoi(string(token))
			area = append(area, n)
		}
	}

	return Heightmap{area: area, width: width, height: height}
}

func part1(hmap Heightmap) int {
	risk := 0
	for p := 0; p < len(hmap.area); p++ {
		isLowPoint := true

		if p/hmap.width > 0 {
			// Look up.
			isLowPoint = isLowPoint && (hmap.area[p] < hmap.area[p-hmap.width])
		}
		if p/hmap.width < (hmap.height - 1) {
			// Look down.
			isLowPoint = isLowPoint && (hmap.area[p] < hmap.area[p+hmap.width])
		}
		if p%hmap.width > 0 {
			// Look left.
			isLowPoint = isLowPoint && (hmap.area[p] < hmap.area[p-1])
		}
		if p%hmap.width < (hmap.width - 1) {
			// Look right.
			isLowPoint = isLowPoint && (hmap.area[p] < hmap.area[p+1])
		}

		if isLowPoint {
			risk += hmap.area[p] + 1
		}
	}
	return risk
}

func part2(hmap Heightmap) int {
	// Loop through all unvisited cells and find basins.
	basins := make([]int, 0)
	for i := 0; i < len(hmap.area); i++ {
		if _, ok := visited[i]; !ok {
			basin := floodCount(i, hmap)
			if basin > 0 {
				basins = append(basins, basin)
			}
		}
	}

	// Find top 3 basins.
	sort.Ints(basins)
	total := 1
	for _, basin := range basins[len(basins)-3:] {
		total *= basin
	}
	return total
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	visited = make(map[int]bool)
	hmap := parse(data)
	fmt.Println(part1(hmap))
	hmap = parse(data)
	fmt.Println(part2(hmap))
}
