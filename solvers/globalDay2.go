package solvers

import (
	"fmt"
	"strconv"
	"strings"

	"firstbulb.xyz/aoc-2025/utils"
)

func Day2Part1(path string) {
	if path == "" {
		path = "input/test/day_2.txt"
	}

	input, err := utils.ParseFileToStringArray(path)

	if err != nil {
		panic("Coucl not parse file")
	}

	var res uint64

	for _, codeRange := range input {
		if codeRange == "" {
			continue
		}

		startAndStopCode := strings.Split(codeRange, "-")

		start, err := strconv.ParseUint(startAndStopCode[0], 10, 64)
		if err != nil {
			panic("Coucl not parse file")
		}

		end, err := strconv.ParseUint(startAndStopCode[1], 10, 64)

		if err != nil {
			panic("Coucl not parse file")
		}

		for i := start; i <= end; i++ {
			isValid := validateId(strconv.FormatUint(i, 10))

			if !isValid {
				res = res + i
			}
		}
	}

	fmt.Println("The answer is ", res)
}

func Day2Part2(path string) {
	if path == "" {
		path = "input/test/day_2.txt"
	}

	input, err := utils.ParseFileToStringArray(path)
	if err != nil {
		panic("Coucl not parse file")
	}

	var res uint64

	for _, codeRange := range input {
		if codeRange == "" {
			continue
		}

		startAndStopCode := strings.Split(codeRange, "-")

		start, err := strconv.ParseUint(startAndStopCode[0], 10, 64)
		if err != nil {
			panic("Coucl not parse file")
		}

		end, err := strconv.ParseUint(startAndStopCode[1], 10, 64)

		if err != nil {
			panic("Coucl not parse file")
		}

		for i := start; i <= end; i++ {
			isValid := validateId2(strconv.FormatUint(i, 10))

			if !isValid {
				res = res + i
			}
		}
	}

	fmt.Println("The answer is ", res)
}

func validateId(id string) bool {
	// No leading zeroes
	if id[0] == '0' {
		return true
	}

	if len(id)%2 != 0 {
		return true
	}

	return id[:len(id)/2] != id[len(id)/2:]
}

func validateId2(id string) bool {
	for i := 0; i < len(id)/2; i++ {
		validPattern := checkPattern(id[0:i+1], id)
		if !validPattern {
			fmt.Println(id)
			return false
		}
	}

	return true
}

func checkPattern(pattern, str string) bool {
	if len(str)%len(pattern) != 0 {
		return true
	}

	parts := make([]string, 0)

	for i := len(pattern); i < len(str); i += len(pattern) {
		parts = append(parts, str[i:i+len(pattern)])
	}

	for _, part := range parts {
		if pattern != part {
			return true
		}
	}

	return false
}
