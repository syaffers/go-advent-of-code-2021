package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func minMinusMax(counter map[string]int) int {
	maxCount, minCount := 0, 1<<63-1
	for _, count := range counter {
		if count > maxCount {
			maxCount = count
		}
		if count < minCount {
			minCount = count
		}
	}
	return maxCount - minCount
}

func parse(data []string) (string, map[string]string) {
	template := ""
	rules := make(map[string]string)
	state := 0

	for _, line := range data {
		if line == "" {
			state++
			continue
		}
		if state == 0 {
			template = line
		} else {
			tokens := strings.Split(line, " -> ")
			rules[tokens[0]] = tokens[1]
		}
	}

	return template, rules
}

func part1(template string, rules map[string]string) int {
	for k := 0; k < 10; k++ {
		compound := ""
		for i := 0; i < len(template)-1; i++ {
			part := template[i : i+2]
			compound += template[i:i+1] + rules[part]
			if i == len(template)-2 {
				compound += template[i+1 : i+2]
			}
		}
		template = compound
	}

	counter := make(map[string]int)
	for _, element := range template {
		counter[string(element)]++
	}

	return minMinusMax(counter)
}

func part2(template string, rules map[string]string) int {
	// Populate all results.
	results := make(map[string][2]string)
	for rule, result := range rules {
		pair := [2]string{rule[0:1] + result, result + rule[1:]}
		results[rule] = pair
	}

	// Count starting bigrams from template
	bigrams := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		part := template[i : i+2]
		bigrams[part]++
	}

	for k := 0; k < 40; k++ {

		newBigrams := make(map[string]int)
		for bigram, count := range bigrams {
			newBigrams[results[bigram][0]] += count
			newBigrams[results[bigram][1]] += count
		}
		bigrams = newBigrams
	}

	counter := make(map[string]int)
	for bigram, count := range bigrams {
		counter[bigram[0:1]] += count
		// counter[bigram[1:]] += count
	}
	counter[template[len(template)-1:]]++

	return minMinusMax(counter)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}

	template, rules := parse(data)
	fmt.Println(part1(template, rules))
	template, rules = parse(data)
	fmt.Println(part2(template, rules))
}
