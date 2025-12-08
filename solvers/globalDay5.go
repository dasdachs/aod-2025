package solvers

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"firstbulb.xyz/aoc-2025/utils"
)

func Day5Part1(path string) {
	if path == "" {
		path = "input/test/day_5.txt"
	}

	input, err := utils.ParseFileToStringArray(path)

	if err != nil {
		panic("Coucl not parse file")
	}

	splited := splitInput(input)
	fresh := splited[0]
	ingredients := splited[1]

	freshIdRanges := freshIds(fresh)

	var count uint64
	lowest := freshIdRanges[0][0]

	for _, ingredient := range ingredients {
		id, _ := strconv.ParseUint(ingredient, 10, 64)

		if id < lowest {
			continue
		}

		for _, r := range freshIdRanges {
			if id >= r[0] && id <= r[1] {
				count++
				break
			}
		}
	}
	fmt.Println("Fresh ingredients ", count)
}

func Day5Part2(path string) {
	if path == "" {
		path = "input/test/day_5.txt"
	}

	input, err := utils.ParseFileToStringArray(path)
	if err != nil {
		panic("Could not parse file")
	}
	splited := splitInput(input)
	fresh := splited[0]

	freshIdRanges := freshIds(fresh)

	if len(freshIdRanges) == 0 {
		fmt.Println("Total number of fresh ingredients ", 0)
		return
	}

	mergedRanges := make([][2]uint64, 0)
	mergedRanges = append(mergedRanges, freshIdRanges[0])

	for i := 1; i < len(freshIdRanges); i++ {
		currentRange := freshIdRanges[i]
		lastMergedRange := &mergedRanges[len(mergedRanges)-1]

		if currentRange[0] <= lastMergedRange[1] {
			if currentRange[1] > lastMergedRange[1] {
				lastMergedRange[1] = currentRange[1]
			}
		} else {
			mergedRanges = append(mergedRanges, currentRange)
		}
	}

	var total uint64
	for _, r := range mergedRanges {
		total += r[1] - r[0] + 1
	}

	fmt.Println("Total number of fresh ingredients ", total)
}

func splitInput(input []string) [2][]string {

	freshIds := make([]string, 0)
	ingredients := make([]string, 0)

	split := false

	for _, line := range input {
		if line == "" {
			split = true
			continue
		}
		if !split {
			freshIds = append(freshIds, line)
		} else {
			ingredients = append(ingredients, line)
		}
	}

	var splited [2][]string

	splited[0] = freshIds
	splited[1] = ingredients

	return splited
}

func freshIds(input []string) [][2]uint64 {
	ids := make([][2]uint64, 0)

	for _, line := range input {
		r := regexp.MustCompile(`\d+`)

		numbers := r.FindAllString(line, -1)

		if len(numbers) < 2 {
			continue
		}

		lower, _ := strconv.ParseUint(numbers[0], 10, 64)
		upper, _ := strconv.ParseUint(numbers[1], 10, 64)

		if lower > upper {
			tmp := lower
			upper = lower
			lower = tmp
		}

		ids = append(ids, [2]uint64{lower, upper})
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i][0] < ids[j][0]
	})

	return ids
}
