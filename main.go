package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	GRID_ROWS = 40
	GRID_COLS = 40
)

// term escapes
const (
	TERM_RESET_CURSOR = "\033[H\033[2J"

	// colors
	COLOR_RESET   = "\033[0m"
	COLOR_RED     = "\033[31m"
	COLOR_GREEN   = "\033[32m"
	COLOR_YELLOW  = "\033[33m"
	COLOR_BLUE    = "\033[34m"
	COLOR_MAGENTA = "\033[35m"
	COLOR_CYAN    = "\033[36m"
	COLOR_GRAY    = "\033[37m"
	COLOR_WHITE   = "\033[97m"
)

func main() {
	iteration := 0
	grid := [GRID_ROWS][GRID_COLS]int{}

	// # # #
	// # O #
	// # # #

	// Rules
	//
	// Underpopulation: fewer than 2 neighbors dies
	// Survive: 2 or 3 neighbors
	// Overpopulation: more than 3 neighbors
	// Reproduction: Any dead cell with exactly 3 live neighbors

	// init grid
	gliderGun(&grid)
	// randomPattern(&grid)
	for {
		printOut(grid)

		iteration++
		fmt.Printf("%sIteration: %d%s", COLOR_GREEN, iteration, COLOR_RESET)
		grid = nextGeneration(grid)
		time.Sleep(time.Millisecond * 200)
	}
}

func printOut(grid [GRID_ROWS][GRID_COLS]int) {
	var strB strings.Builder
	strB.WriteString(TERM_RESET_CURSOR)

	// top borders
	for q := 0; q < GRID_COLS; q++ {
		strB.WriteRune('_')
	}
	strB.WriteString("\n")

	for i := 0; i < GRID_ROWS; i++ {
		for j := 0; j < GRID_COLS; j++ {

			// left borders
			if j == 0 {
				strB.WriteRune('|')
			}

			if grid[i][j] == 1 {
				strB.WriteString(COLOR_WHITE)
				strB.WriteRune('O')
				strB.WriteString(COLOR_RESET)
			} else {
				strB.WriteString(COLOR_RED)
				strB.WriteRune('+')
				strB.WriteString(COLOR_RESET)
			}

			if j == GRID_COLS-1 {
				strB.WriteRune('|')
			}
		}

		strB.WriteString("\n")
	}

	// bottom borders
	for q := 0; q <= GRID_COLS; q++ {
		strB.WriteRune('-')
	}

	fmt.Println(strB.String())
}

func countLiveNeighbors(grid [GRID_ROWS][GRID_COLS]int, row, col int) int {
	// # # #
	// # O #
	// # # #
	//
	// { i  j } <- index
	// {-1, -1} {-1, 0} {-1, 1}
	// { 0, -1} { 0, 0} { 0, 1}
	// { 1,  1} { 1, 0} { 1, 1}
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			// if current cell
			if i == 0 && j == 0 {
				continue
			}

			// count alive neighbors
			// if not out of bounds
			if (row+i >= 0 && col+j >= 0) && (GRID_ROWS > row+i && GRID_COLS > col+j) {
				count += grid[row+i][col+j]
			}
		}
	}

	return count
}

func nextGeneration(grid [GRID_ROWS][GRID_COLS]int) [GRID_ROWS][GRID_COLS]int {
	gen := [GRID_ROWS][GRID_COLS]int{}

	for i := 0; i < GRID_ROWS; i++ {
		for j := 0; j < GRID_COLS; j++ {
			liveNeighbors := countLiveNeighbors(grid, i, j)

			isAlive := grid[i][j] == 1
			isOverPopulation := liveNeighbors > 3
			isUnderPopulation := liveNeighbors < 2
			if isAlive {
				if isOverPopulation || isUnderPopulation {
					gen[i][j] = 0
				} else {
					gen[i][j] = 1
				}
			} else {
				// if 3 neighbors alive then reproduce
				if liveNeighbors == 3 {
					gen[i][j] = 1
				}
			}
		}
	}

	return gen
}

func gliderGun(grid *[GRID_ROWS][GRID_COLS]int) {
	gliderGun := [36][2]int{
		{1, 10},
		{1, 11},
		{2, 10},
		{2, 11},
		{11, 10},
		{11, 11},
		{11, 12},
		{12, 9},
		{12, 13},
		{13, 8},
		{13, 14},
		{14, 8},
		{14, 14},
		{15, 11},
		{16, 9},
		{16, 13},
		{17, 10},
		{17, 11},
		{17, 12},
		{18, 11},
		{21, 8},
		{21, 9},
		{21, 10},
		{22, 8},
		{22, 9},
		{22, 10},
		{23, 7},
		{23, 11},
		{25, 6},
		{25, 7},
		{25, 11},
		{25, 12},
		{35, 8},
		{35, 9},
		{36, 8},
		{36, 9},
	}

	for _, cell := range gliderGun {
		grid[cell[0]][cell[1]] = 1
	}
}

// idk this pattern is, but its good
func randomPattern(grid *[GRID_ROWS][GRID_COLS]int) {
	pattern := [36][2]int{
		{1, 15},
		{1, 16},
		{2, 15},
		{2, 16},
		{11, 15},
		{11, 16},
		{11, 17},
		{12, 14},
		{12, 18},
		{13, 13},
		{13, 19},
		{14, 13},
		{14, 19},
		{15, 16},
		{16, 15},
		{16, 18},
		{17, 15},
		{17, 16},
		{17, 17},
		{18, 16},
		{21, 13},
		{21, 14},
		{21, 15},
		{22, 13},
		{22, 14},
		{22, 15},
		{23, 12},
		{23, 16},
		{25, 11},
		{25, 12},
		{25, 16},
		{25, 17},
		{35, 13},
		{35, 14},
		{36, 13},
		{36, 14},
	}

	for _, cell := range pattern {
		grid[cell[0]][cell[1]] = 1
	}
}
