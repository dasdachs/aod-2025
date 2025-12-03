package solvers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"firstbulb.xyz/aoc-2025/utils"
)

func Day3Part1(path string) {
	if path == "" {
		path = "input/test/day_3.txt"
	}

	input, err := utils.ParseFileToStringArray(path)

	if err != nil {
		panic("Coucl not parse file")
	}

	var joltage uint64

	for _, line := range input {
		if line == "" {
			continue
		}
		bankJoltage := caclulatePowerBank(line, 2)

		joltage += bankJoltage
	}

	fmt.Printf("The joltage is %d", joltage)
}

func Day3Part2(path string) {
	if path == "" {
		path = "input/test/day_3.txt"
	}

	input, err := utils.ParseFileToStringArray(path)
	if err != nil {
		panic("Coucl not parse file")
	}
	var joltage uint64

	for _, line := range input {
		if line == "" {
			continue
		}
		bankJoltage := caclulatePowerBank(line, 12)

		joltage += bankJoltage
	}

	fmt.Printf("The joltage is %d", joltage)
}

func caclulatePowerBank(numberString string, length int) uint64 {
	searchString := numberString
	var out strings.Builder

	for {
		if out.Len() == length {
			break
		}

		reminingLength := length - out.Len()
		var num string

		if out.Len() == length-1 {
			number, _ := findLargerstNumber(searchString)
			num = number
		} else {
			number, idx := findLargerstNumber(searchString[:len(searchString)-reminingLength+1])
			searchString = searchString[idx+1:]
			num = number
		}

		out.WriteString(num)
	}

	res, _ := strconv.ParseUint(out.String(), 10, 64)

	return res
}

func findLargerstNumber(line string) (string, int) {
	lineSlice := strings.Split(line, "")

	sort.Slice(lineSlice, func(i, j int) bool {
		return lineSlice[i] > lineSlice[j]
	})

	largerstNumber := string(lineSlice[0])

	var idx int

	for i, n := range line {
		if string(n) == largerstNumber {
			idx = i
			break
		}
	}

	return largerstNumber, idx
}
