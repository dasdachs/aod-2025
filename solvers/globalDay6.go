package solvers

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"firstbulb.xyz/aoc-2025/utils"
)

func Day6Part1(path string) {
	if path == "" {
		path = "input/test/day_6.txt"
	}

	input, err := utils.ParseFileToStringArray(path)

	if err != nil {
		panic("Coucl not parse file")
	}

	processedInput := make([][]string, 0)

	for _, line := range input {
		if line != "" {
			pl := procesLines(line)
			processedInput = append(processedInput, pl)
		}
	}

	var total uint64

	for i := range processedInput[0] {
		res := calculateLine(processedInput, i)
		total += res
	}

	fmt.Println("Sum is ", total)
}

func Day6Part2(path string) {
	if path == "" {
		path = "input/test/day_6.txt"
	}

	input, err := utils.ParseFileToStringArray(path)
	if err != nil {
		panic("Coucl not parse file")
	}

	processed := processRtlLines(input[:len(input)-2])
	op := procesLines(input[len(input)-2])

	var total uint64

	for i := 0; i < len(processed[0]); i++ {
		line := make([]string, 0)
		for _, p := range processed {
			line = append(line, p[i])
		}

		res := calculateRtlLine(line, op[i])
		total += res
	}

	fmt.Println(total)
}

func procesLines(line string) []string {
	r := regexp.MustCompile(`[0-9\+\*]+`)

	res := r.FindAllString(line, -1)

	return res
}

func calculateLine(data [][]string, idx int) uint64 {
	op := data[len(data)-1][idx]

	res, _ := strconv.ParseUint(data[0][idx], 10, 64)

	for _, i := range data[1 : len(data)-1] {
		num, _ := strconv.ParseUint(i[idx], 10, 64)
		switch op {
		case "+":
			res += num
		case "-":
			res -= num
		case "*":
			res *= num
			// case ":":
			// 	res \= num
		}
	}

	return res
}

func getLongestNumber(s string) int {
	r := strings.Split(s, "-")

	sort.Slice(r, func(i, j int) bool {
		return len(r[i]) > len(r[j])
	})

	res := len(r[0])
	return res
}

func processRtlLines(data []string) [][]string {
	first := data[0]

	breaks := []int{0}

	for i, c := range first {
		if string(c) == "-" {
			isBreak := true
			for _, l := range data[1:] {
				if string(l[i]) != "-" {
					isBreak = false
				}
			}
			if isBreak {
				breaks = append(breaks, i)
			}
		}
	}

	res := make([][]string, 0)

	for _, line := range data {
		parsed := make([]string, 0)
		for i, idx := range breaks {
			if idx == 0 {
				parsed = append(parsed, line[idx:breaks[i+1]])
			} else if i == len(breaks)-1 {
				parsed = append(parsed, line[idx+1:])
			} else {
				parsed = append(parsed, line[idx+1:breaks[i+1]])
			}
		}
		res = append(res, parsed)
	}

	return res
}

func calculateRtlLine(data []string, op string) uint64 {
	numbers := make([]uint64, 0)
	lastIdx := len(data[0])

	for i := lastIdx - 1; i > -1; i-- {
		var numString strings.Builder

		for _, l := range data {
			c := string(l[i])
			if c != "-" {
				numString.WriteString(c)
			}
		}

		num, _ := strconv.ParseUint(numString.String(), 10, 64)
		numbers = append(numbers, num)
	}

	res := numbers[0]

	for _, num := range numbers[1:] {
		switch op {
		case "+":
			res += num
		case "*":
			res *= num
			// case ":":
			// 	res \= num
		}
	}

	return res
}
