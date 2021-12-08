package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	signals []string
	outputs []string
}

func runemapToString(m map[rune]bool) string {
	s := ""
	for k, v := range m {
		if v {
			s += string(k)
		}
	}
	return s
}

func charDiff(s string, t string) string {
	m := make(map[rune]bool, 0)
	for _, c := range s {
		m[c] = true
	}
	for _, c := range t {
		m[c] = false
	}

	return runemapToString(m)
}

func outputToIntString(output string, segMap []string) string {
	segIdx := make(map[string]int, 0)
	segSet := []bool{false, false, false, false, false, false, false}

	for i, c := range segMap {
		segIdx[c] = i
	}

	switch len(output) {
	case 2:
		return "1"
	case 3:
		return "7"
	case 4:
		return "4"
	case 5:
		for _, c := range output {
			segSet[segIdx[string(c)]] = true
		}
		if !segSet[1] && !segSet[5] {
			return "2"
		}
		if !segSet[1] && !segSet[4] {
			return "3"
		}
		if !segSet[2] && !segSet[4] {
			return "5"
		}
	case 6:
		for _, c := range output {
			segSet[segIdx[string(c)]] = true
		}
		if !segSet[2] {
			return "6"
		}
		if !segSet[3] {
			return "0"
		}
		if !segSet[4] {
			return "9"
		}
	case 7:
		return "8"
	}
	return "-1"
}

func parse(data []string) []Entry {
	entries := make([]Entry, 0)

	for _, line := range data {
		i := 0
		isSignal := true
		entry := Entry{signals: make([]string, 10), outputs: make([]string, 4)}
		for _, token := range strings.Split(line, " ") {
			if token == "|" {
				isSignal = false
				i = 0
				continue
			}

			if isSignal {
				entry.signals[i] = token
				i++
			} else {
				entry.outputs[i] = token
				i++
			}
		}
		entries = append(entries, entry)
	}

	return entries
}

func part1(entries []Entry) int {
	counter := 0
	for _, entry := range entries {
		for _, output := range entry.outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				counter++
			}
		}
	}
	return counter
}

func part2(entries []Entry) int {
	result := 0
	for _, entry := range entries {
		numMap := make([]string, 10)
		segMap := []string{"", "", "", "", "", "", ""}
		sixSegs := make([]string, 0)
		fiveSegs := make([]string, 0)

		for _, signal := range entry.signals {
			switch len(signal) {
			case 2: // Trivial 1.
				numMap[1] = signal
			case 3: // Trivial 7.
				numMap[7] = signal
			case 4: // Trivial 4.
				numMap[4] = signal
			case 5: // Non-trivial 2, 3, or 5.
				fiveSegs = append(fiveSegs, signal)
			case 6: // Non-trivial 0, 6, or 9.
				sixSegs = append(sixSegs, signal)
			case 7: // Trivial 8.
				numMap[8] = signal
			}
		}

		// Segment 'a' is the 7 segment without the 1 segment.
		segMap[0] = charDiff(numMap[7], numMap[1])

		// Segment 'f' is the 8 segment without the 6 segment (i.e. just the 2
		// segment). The way to get this is to subtract all the 6-segmented
		// numbers from the 8 segment and if there is one segment left after
		// subtracting the remainder with the 1 segment, it must be 'f'.
		for _, s := range sixSegs {
			diff := charDiff(numMap[8], s)
			diff = charDiff(numMap[1], diff)
			if len(diff) == 1 {
				segMap[5] = diff
			}
		}

		// Segment 'c' is the 1 segment without the 'f' segment.
		segMap[2] = charDiff(numMap[1], segMap[5])

		// Segment 'e' is the 8 segment without the 5 segment and segment 'c'.
		// The way to get this is to subtract all the 5-segmented numbers from
		// the 8 segment and if there is one segment left after subtracting
		// segment 'c', it must be 'e'.
		for _, s := range fiveSegs {
			diff := charDiff(numMap[8], s)
			diff = charDiff(diff, segMap[2])
			if len(diff) == 1 {
				segMap[4] = diff
			}
		}

		// Segment 'd' is the 8 segment without any 6-segmented numbers and
		// segment 'c' and segment 'e'.
		for _, s := range sixSegs {
			diff := charDiff(numMap[8], s)
			diff = charDiff(charDiff(diff, segMap[2]), segMap[4])
			if len(diff) == 1 {
				segMap[3] = diff
			}
		}

		// Segment 'g' is the 2 segment number without 'a', 'c', 'd', 'e'. If
		// we loop through all 6-segments and subtract the segments listed
		// above, we should be left with one that is 'g'.
		for _, s := range fiveSegs {
			diff := charDiff(s, segMap[0])
			diff = charDiff(diff, segMap[2])
			diff = charDiff(diff, segMap[3])
			diff = charDiff(diff, segMap[4])
			if len(diff) == 1 {
				segMap[6] = diff
			}
		}

		// Segment 1 is the one that's missing!
		inSegMap := ""
		for _, s := range segMap {
			inSegMap += s
		}
		segMap[1] = charDiff("abcdefg", inSegMap)

		// Assemble.
		intString := ""
		for _, output := range entry.outputs {
			intString += outputToIntString(output, segMap)
		}
		r, _ := strconv.Atoi(intString)
		result += r
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	entries := parse(data)
	fmt.Println(part1(entries))
	entries = parse(data)
	fmt.Println(part2(entries))
}
