package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	InfoColor = "\033[1;31m%02d \033[0m"
)

type Cell struct {
	number int
	seen   bool
}

type Board struct {
	cells      []Cell
	solvedTime int
}

func parse(data []string) ([]int, []Board) {
	// Parse the first line.
	comma := regexp.MustCompile(",")
	drawTokens := comma.Split(data[0], -1)
	draws := make([]int, len(drawTokens))
	for i, n := range drawTokens {
		draws[i], _ = strconv.Atoi(n)
	}

	// Parse all other lines.
	multiSpace := regexp.MustCompile(`\s+`)
	boards := make([]Board, 0)
	board := Board{cells: make([]Cell, 25), solvedTime: -1}
	c := 0
	for _, line := range data[1:] {
		// Skip empty lines.
		if len(line) == 0 {
			continue
		}
		// Parse every 5 lines into boards.
		tokens := multiSpace.Split(strings.Trim(line, " "), -1)
		for i := 0; i < 5; i++ {
			n, _ := strconv.Atoi(tokens[i])
			cell := Cell{number: n, seen: false}
			board.cells[5*c+i] = cell
		}
		c = (c + 1) % 5
		if c == 0 {
			boards = append(boards, board)
			board = Board{cells: make([]Cell, 25), solvedTime: -1}
		}
	}

	return draws, boards
}

func printBoard(board Board) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.cells[5*i+j].seen {
				fmt.Printf(InfoColor, board.cells[5*i+j].number)
			} else {
				fmt.Printf("%02d ", board.cells[5*i+j].number)
			}
		}
		fmt.Println()
	}
	fmt.Println("Solved at:", board.solvedTime)
}

func unmarkedSum(board Board) int {
	value := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.cells[5*i+j].seen {
				value += board.cells[5*i+j].number
			}
		}
	}
	return value
}

func isSolved(board Board) bool {
	solved := false
	// Check rows.
	for i := 0; i < 5; i++ {
		rowIsSolved := true
		for j := 0; j < 5; j++ {
			rowIsSolved = rowIsSolved && board.cells[5*i+j].seen
		}
		solved = solved || rowIsSolved
	}
	// Check columns.
	for i := 0; i < 5; i++ {
		colIsSolved := true
		for j := 0; j < 5; j++ {
			colIsSolved = colIsSolved && board.cells[i+5*j].seen
		}
		solved = solved || colIsSolved
	}

	return solved
}

func part1(draws []int, boards []Board) int {
	for _, n := range draws {
		for _, board := range boards {
			for j := 0; j < 25; j++ {
				if board.cells[j].number == n {
					board.cells[j].seen = true
					if isSolved(board) {
						return unmarkedSum(board) * n
					}
				}
			}
		}
	}
	return 0
}

func part2(draws []int, boards []Board) int {
	nBoards := len(boards)
	for k, n := range draws {
		nSolved := 0
		for i, board := range boards {
			for j := 0; j < 25; j++ {
				if board.cells[j].number == n {
					board.cells[j].seen = true
				}
			}
			if isSolved(board) {
				if board.solvedTime < 0 {
					boards[i].solvedTime = k
				}
				nSolved++
			}
		}

		if nSolved == nBoards {
			lastBoard := boards[0]
			lastSolvedTime := boards[0].solvedTime
			for _, board := range boards {
				if board.solvedTime > lastSolvedTime {
					lastBoard = board
					lastSolvedTime = board.solvedTime
				}
			}
			return unmarkedSum(lastBoard) * n
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0, 0) // This is an empty list.

	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	draws, boards := parse(data)
	fmt.Println(part1(draws, boards))
	draws, boards = parse(data)
	fmt.Println(part2(draws, boards))
}
