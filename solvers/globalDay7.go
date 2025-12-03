package solvers

import (
	"firstbulb.xyz/aoc-2025/utils"
)

func Day7Part1(path string) {
	if path == "" {
		path = "input/test/day_7.txt"
	}

	_, err := utils.ParseFileToStringArray(path)

	if err != nil {
		panic("Coucl not parse file")
	}
}

func Day7Part2(path string) {
	if path == "" {
		path = "input/test/day_7.txt"
	}

	_, err := utils.ParseFileToStringArray(path)
	if err != nil {
		panic("Coucl not parse file")
	}
}
