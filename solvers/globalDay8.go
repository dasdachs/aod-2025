package solvers

import (
	"firstbulb.xyz/aoc-2025/utils"
)

func Day8Part1(path string) {
	if path == "" {
		path = "input/test/day_8.txt"
	}

	_, err := utils.ParseFileToStringArray(path)

	if err != nil {
		panic("Coucl not parse file")
	}
}

func Day8Part2(path string) {
	if path == "" {
		path = "input/test/day_8.txt"
	}

	_, err := utils.ParseFileToStringArray(path)
	if err != nil {
		panic("Coucl not parse file")
	}
}
