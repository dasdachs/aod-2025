package solvers

import (
	"fmt"
	"strings"

	"firstbulb.xyz/aoc-2025/utils"
)

func Day4Part1(path string) {
	if path == "" {
		path = "input/test/day_4.txt"
	}

	input, err := utils.ParseFileToStringArray(path)

	if err != nil {
		panic("Coucl not parse file")
	}

	grid := generateGrid(input)

	paper, _ := caclulatePaper(grid)

	fmt.Printf("Number of rols: %d", paper)
}

func Day4Part2(path string) {
	if path == "" {
		path = "input/test/day_4.txt"
	}

	input, err := utils.ParseFileToStringArray(path)
	if err != nil {
		panic("Coucd not parse file")
	}
	grid := generateGrid(input)

	var total int

	for {
		paper, coordinates := caclulatePaper(grid)

		if paper == 0 {
			break
		}

		total += paper

		for _, coordinate := range coordinates {
			r := coordinate[0]
			c := coordinate[1]
			grid[r][c] = "."
		}
	}

	fmt.Printf("Number of rols: %d", total)
}

func caclulatePaper(grid [][]string) (int, [][]int) {
	var paper int
	coordinates := make([][]int, 0)

	for rowIdx, row := range grid {
		for colIdx, col := range row {
			subGrid := getPointData(rowIdx, colIdx, grid)

			if col == "@" {
				var total int

				data := make([]string, 0, len(subGrid.above)+len(subGrid.current)+len(subGrid.under))
				data = append(data, subGrid.above...)
				data = append(data, subGrid.current...)
				data = append(data, subGrid.under...)

				for _, c := range data {
					if c == "@" {
						total++
					}
				}

				if total < 5 {
					coordinates = append(coordinates, []int{rowIdx, colIdx})
					paper++
				}
			}
		}
	}
	return paper, coordinates
}

func generateGrid(input []string) [][]string {
	grid := make([][]string, 0)

	for _, row := range input {
		if row == "" {
			continue
		}

		r := strings.Split(row, "")
		grid = append(grid, r)
	}
	return grid
}

// 3x3 grid
func getPointData(row, column int, grid [][]string) struct {
	above   []string
	current []string
	under   []string
} {
	startIdx := column
	if column > 0 {
		startIdx -= 1
	}

	endIdx := column + 1
	if column < len(grid[row])-1 {
		endIdx += 1
	}

	var subGrid struct {
		above   []string
		current []string
		under   []string
	}

	subGrid.current = grid[row][startIdx:endIdx]

	if row != 0 {
		subGrid.above = grid[row-1][startIdx:endIdx]
	}
	if row < len(grid)-1 {
		subGrid.under = grid[row+1][startIdx:endIdx]
	}

	return subGrid
}
