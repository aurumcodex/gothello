package main

import "fmt"

// [===== utility functions =====]

//
func getCells(moveset []Move) []int {
	var cells []int

	for _, c := range moveset {
		cells = append(cells, c.cell)
	}

	return cells
}

// [===== printing functions =====]
func getRow(x int) int {
	return (x / 8) + 1
}

func printChar(i int, s string) {
	if i%8 == 7 {
		fmt.Printf(" %v\n", s)
	} else if i%8 != 7 {
		fmt.Printf(" %v", s)
	}
}

func color(n int, full bool) string {
	switch n {
	case black:
		if !full {
			return "B"
		}
		return "Black"
	case white:
		if !full {
			return "W"
		}
		return "White"
	default:
		return ""
	}
}

func printMove(i int, list []int) {
	if sliceContains(i, list) {
		printChar(i, "+")
	} else {
		printChar(i, "-")
	}
}

func printBlack(i int) {
	printChar(i, "B")
}

func printWhite(i int) {
	printChar(i, "W")
}

// [===== general purpose functions =====]
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func contains(i int, list [8]int) bool {
	for _, j := range list {
		if j == i {
			return true
		}
	}
	return false
}

func sliceContains(i int, list []int) bool {
	for _, j := range list {
		if j == i {
			return true
		}
	}
	return false
}
