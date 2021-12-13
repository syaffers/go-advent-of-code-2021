package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Globals.
var frontier []string

func dfs(caveMap map[string][]string, visited map[string]bool) int {
	count := 0

	// Copy visited set since Go is PBR for maps.
	visited_ := make(map[string]bool)
	for k, v := range visited {
		visited_[k] = v
	}

	if len(frontier) == 0 {
		return 0
	}
	currCave := frontier[0]
	frontier = frontier[1:]

	if currCave == "end" {
		return 1
	}

	if unicode.IsLower(rune(currCave[0])) {
		visited_[currCave] = true
	}

	for _, nextCave := range caveMap[currCave] {
		if _, isVisited := visited_[nextCave]; !isVisited {
			frontier = append(frontier, nextCave)
			count += dfs(caveMap, visited_)
		}
	}

	return count
}

func dfs2(caveMap map[string][]string, visited map[string]int) int {
	count := 0

	// Copy visited set since Go is PBR for maps.
	visited_ := make(map[string]int)
	for k, v := range visited {
		visited_[k] = v
	}

	if len(frontier) == 0 {
		return 0
	}
	currCave := frontier[0]
	frontier = frontier[1:]

	if currCave == "end" {
		return 1
	}

	if unicode.IsLower(rune(currCave[0])) {
		visited_[currCave]++
	}

	hasDoubleVisit := false
	for _, nVisits := range visited_ {
		hasDoubleVisit = hasDoubleVisit || (nVisits > 1)
	}

	for _, nextCave := range caveMap[currCave] {
		_, isVisited := visited_[nextCave]
		if hasDoubleVisit {
			if !isVisited {
				frontier = append(frontier, nextCave)
				count += dfs2(caveMap, visited_)
			}
		} else {
			if nextCave != "start" {
				frontier = append(frontier, nextCave)
				count += dfs2(caveMap, visited_)
			}
		}
	}

	return count
}

func parse(data []string) map[string][]string {
	caveMap := make(map[string][]string)
	var sourceCaves []string
	var targetCaves []string

	for _, line := range data {
		tokens := strings.Split(line, "-")

		if _, ok := caveMap[tokens[0]]; ok {
			targetCaves = caveMap[tokens[0]]
			targetCaves = append(targetCaves, tokens[1])
		} else {
			targetCaves = make([]string, 0)
			targetCaves = append(targetCaves, tokens[1])
		}

		if _, ok := caveMap[tokens[1]]; ok {
			sourceCaves = caveMap[tokens[1]]
			sourceCaves = append(sourceCaves, tokens[0])
		} else {
			sourceCaves = make([]string, 0)
			sourceCaves = append(sourceCaves, tokens[0])

		}

		caveMap[tokens[0]] = targetCaves
		caveMap[tokens[1]] = sourceCaves
	}

	return caveMap
}

func part1(caveMap map[string][]string) int {
	frontier = make([]string, 0)
	frontier = append(frontier, "start")
	visited := make(map[string]bool)
	visited["start"] = true
	return dfs(caveMap, visited)
}

func part2(caveMap map[string][]string) int {
	frontier = make([]string, 0)
	frontier = append(frontier, "start")
	visited := make(map[string]int)

	return dfs2(caveMap, visited)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}

	caveMap := parse(data)
	fmt.Println(part1(caveMap))
	caveMap = parse(data)
	fmt.Println(part2(caveMap))
}
