package main

import (
	"bufio"
	"fmt"
	"os"
)

func bin2bit(bins string, bitlen int) []int {
	bitarr := make([]int, bitlen)
	for i, bin := range bins {
		if bin == 49 {
			bitarr[i] = 1
		} else {
			bitarr[i] = 0
		}
	}
	return bitarr
}

func bit2int(bitarr []int, bitlen int) int {
	xr := 1 << (bitlen - 1)
	result := 0

	for _, bit := range bitarr {
		result += xr * bit
		xr >>= 1
	}

	return result
}

func o2_rating(bins []string, bitpos int, bitlen int) int {
	// Base case.
	if len(bins) == 1 {
		return bit2int(bin2bit(bins[0], bitlen), bitlen)
	}

	// Recursive case.
	filtered_bins := make([]string, 0)
	zeros := 0
	ones := 0
	var check byte = 49

	for _, bin := range bins {
		if bin[bitpos] == 48 {
			zeros++
		} else {
			ones++
		}
	}

	if zeros > ones {
		check = 48
	}

	for _, bin := range bins {
		if bin[bitpos] == check {
			filtered_bins = append(filtered_bins, bin)
		}
	}
	return o2_rating(filtered_bins, bitpos+1, bitlen)
}

func co2_rating(bins []string, bitpos int, bitlen int) int {
	// Base case.
	if len(bins) == 1 {
		return bit2int(bin2bit(bins[0], bitlen), bitlen)
	}

	// Recursive case.
	filtered_bins := make([]string, 0)
	zeros := 0
	ones := 0
	var check byte = 48

	for _, bin := range bins {
		if bin[bitpos] == 48 {
			zeros++
		} else {
			ones++
		}
	}

	if zeros > ones {
		check = 49
	}

	for _, bin := range bins {
		if bin[bitpos] == check {
			filtered_bins = append(filtered_bins, bin)
		}
	}
	return co2_rating(filtered_bins, bitpos+1, bitlen)
}

func part1(data []string) int {
	n := len(data)
	bitlen := len(data[0])
	gamma := make([]int, bitlen)
	epsilon := make([]int, bitlen)

	for i := 0; i < n; i++ {
		for j := 0; j < bitlen; j++ {
			if data[i][j] == 49 {
				gamma[j]++
			}
		}
	}

	for j := 0; j < bitlen; j++ {
		if gamma[j] > (n / 2) {
			gamma[j] = 1
		} else {
			gamma[j] = 0
		}
		epsilon[j] = 1 - gamma[j]
	}

	return bit2int(gamma, bitlen) * bit2int(epsilon, bitlen)
}

func part2(data []string) int {
	return o2_rating(data, 0, len(data[0])) * co2_rating(data, 0, len(data[0]))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}
