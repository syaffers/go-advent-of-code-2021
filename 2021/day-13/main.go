package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Axis struct {
	x   bool
	loc int
}

type Point struct {
	x int
	y int
}

func fold(points map[Point]bool, axis Axis) {
	if axis.x {
		for p := range points {
			if p.x > axis.loc {
				xNew := p.x - 2*(p.x-axis.loc)
				delete(points, p)
				points[Point{y: p.y, x: xNew}] = true
			}
		}
	} else {
		for p := range points {
			if p.y > axis.loc {
				yNew := p.y - 2*(p.y-axis.loc)
				delete(points, p)
				points[Point{y: yNew, x: p.x}] = true
			}
		}
	}
}

func printPoints(points map[Point]bool) string {
	output := ""
	hMax, wMax := 0, 0

	for p := range points {
		if p.x > wMax {
			wMax = p.x
		}
		if p.y > hMax {
			hMax = p.y
		}
	}

	for i := 0; i <= hMax; i++ {
		for j := 0; j <= wMax; j++ {
			if _, ok := points[Point{x: j, y: i}]; ok {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

func parse(data []string) (map[Point]bool, []Axis) {
	points := make(map[Point]bool, 0)
	axes := make([]Axis, 0)
	state := 0

	for _, line := range data {
		if line == "" {
			state++
			continue
		}
		if state == 0 {
			tokens := strings.Split(line, ",")

			point := Point{x: 0, y: 0}
			point.x, _ = strconv.Atoi(tokens[0])
			point.y, _ = strconv.Atoi(tokens[1])

			points[point] = true
		} else {
			tokens := strings.Split(line, "=")
			ins := Axis{x: true, loc: 0}
			ins.x = rune(tokens[0][len(tokens[0])-1]) == 'x'
			ins.loc, _ = strconv.Atoi(tokens[1])
			axes = append(axes, ins)
		}
	}

	return points, axes
}

func part1(points map[Point]bool, axes []Axis) int {
	fold(points, axes[0])
	return len(points)
}

func part2(points map[Point]bool, axes []Axis) string {
	for _, axis := range axes {
		fold(points, axis)
	}
	return printPoints(points)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}

	points, axes := parse(data)
	fmt.Println(part1(points, axes))
	points, axes = parse(data)
	fmt.Println(part2(points, axes))
}
