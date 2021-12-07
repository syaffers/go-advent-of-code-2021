package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lanternfish struct {
	rate  int
	timer int
}

func mod(a int, b int) int {
	return (a%b + b) % b
}

func parse(data []string) []Lanternfish {
	fishes := make([]Lanternfish, 0)
	for _, token := range strings.Split(data[0], ",") {
		currTime, _ := strconv.Atoi(token)
		fish := Lanternfish{rate: 7, timer: currTime}
		fishes = append(fishes, fish)
	}

	return fishes
}

func part1(fishes []Lanternfish) int {
	for day := 0; day < 80; day++ {
		for i, fish := range fishes {
			fishes[i].timer = mod(fish.timer-1, fish.rate)

			if fishes[i].timer == 0 {
				fishes[i].rate = 7
			} else if fishes[i].timer == (fish.rate - 1) {
				spawn := Lanternfish{rate: fish.rate + 2, timer: fish.rate + 1}
				fishes = append(fishes, spawn)
			}
		}
	}
	return len(fishes)
}

func part2(fishes []Lanternfish) int {
	counter := make([]int, 9)
	tmp := 0

	for _, fish := range fishes {
		counter[fish.timer]++
	}

	for day := 0; day < 256; day++ {
		tmp = counter[0]
		for i := 1; i < 9; i++ {
			counter[i-1] = counter[i]
		}
		counter[6] += tmp
		counter[8] = tmp
	}

	tmp = 0
	for _, n := range counter {
		tmp += n
	}

	return tmp
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	fishes := parse(data)
	fmt.Println(part1(fishes))
	fishes = parse(data)
	fmt.Println(part2(fishes))
}
