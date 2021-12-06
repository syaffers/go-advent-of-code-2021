package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	InfoColor = "\033[1;31m%02d \033[0m"
)

type PointPair struct {
	source Point
	target Point
}

type Point struct {
	x int
	y int
}

func parse(data []string) []PointPair {
	pairs := make([]PointPair, 0)
	for _, line := range data {
		tokens := strings.Split(line, " -> ")
		source := strings.Split(tokens[0], ",")
		target := strings.Split(tokens[1], ",")
		sx, _ := strconv.Atoi(source[0])
		sy, _ := strconv.Atoi(source[1])
		tx, _ := strconv.Atoi(target[0])
		ty, _ := strconv.Atoi(target[1])
		s := Point{x: sx, y: sy}
		t := Point{x: tx, y: ty}
		pair := PointPair{source: s, target: t}
		pairs = append(pairs, pair)
	}

	return pairs
}

func part1(pairs []PointPair) int {
	c, d := 0, 0
	p := Point{x: 0, y: 0}
	vents := make(map[Point]int)

	for _, pair := range pairs {
		s, t := pair.source, pair.target

		if (s.x != t.x) && (s.y != t.y) {
			continue
		}

		if s.x > t.x || s.y > t.y {
			d = -1
		} else if s.x < t.x || s.y < t.y {
			d = 1
		}

		p.x, p.y = s.x, s.y
		for i := s.x; i != t.x; i += d {
			vents[p]++
			p.x += d
		}
		p.x, p.y = s.x, s.y
		for i := s.y; i != t.y; i += d {
			vents[p]++
			p.y += d
		}
		vents[t]++
	}

	for _, v := range vents {
		if v >= 2 {
			c++
		}
	}
	return c
}

func part2(pairs []PointPair) int {
	c := 0
	p := Point{x: 0, y: 0}
	vents := make(map[Point]int)

	for _, pair := range pairs {
		dx, dy := 0, 0
		s, t := pair.source, pair.target

		if s.x > t.x {
			dx = -1
		}
		if s.y > t.y {
			dy = -1
		}
		if s.x < t.x {
			dx = 1
		}
		if s.y < t.y {
			dy = 1
		}

		p.x, p.y = s.x, s.y
		if dy == 0 {
			for i := s.x; i != t.x; i += dx {
				vents[p]++
				p.x += dx
			}
		} else if dx == 0 {
			for i := s.y; i != t.y; i += dy {
				vents[p]++
				p.y += dy
			}
		} else {
			i, j := s.x, s.y
			for {
				vents[p]++
				p.x += dx
				p.y += dy
				i += dx
				j += dy
				if i == t.x && j == t.y {
					break
				}
			}
		}
		vents[t]++
	}

	for _, v := range vents {
		if v >= 2 {
			c++
		}
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
	pairs := parse(data)
	fmt.Println(part1(pairs))
	pairs = parse(data)
	fmt.Println(part2(pairs))
}
