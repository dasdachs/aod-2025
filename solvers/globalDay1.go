package solvers

import (
	"fmt"
	"os"
	"strconv"

	"firstbulb.xyz/aod-2025/utils"
)

func Day1Part1(path string) {
	position := 50
	secretCnt := 0

	if path == "" {
		path = "data/test/day_1.txt"
	}

	input, err := utils.ParseFileToStringArray(path)

	if err != nil {
		fmt.Printf("Could not process file %s", path)
		os.Exit(1)
	}

	for _, data := range input {
		if data == "" {
			continue
		}

		newPosition, isSecret := getPosition(data, position)

		fmt.Printf("The positon is %d\n", newPosition)
		position = newPosition
		if isSecret {
			secretCnt++
		}
	}

	fmt.Printf("The positon is %d\n", position)
	fmt.Printf("The secret is %d", secretCnt)
}

func Day1Part2(path string) {
	position := 50
	secretCnt := 0

	if path == "" {
		path = "data/test/day_1.txt"
	}

	input, err := utils.ParseFileToStringArray(path)

	if err != nil {
		fmt.Printf("Could not process file %s", path)

		os.Exit(1)
	}

	for _, data := range input {
		if data == "" {
			continue
		}

		newPosition, secretCount := getPosition2(data, position)

		position = newPosition
		secretCnt += secretCount
	}

	fmt.Printf("The positon is %d\n", position)
	fmt.Printf("The secret is %d", secretCnt)
}

func getPosition(move string, currentPosition int) (int, bool) {
	fmt.Printf("Move %s\n", move)
	direction := move[0]

	moves, err := strconv.Atoi(move[1:])
	if err != nil {
		panic("String not correct order")
	}

	newPosition := 0

	if direction == 'L' {
		newPosition = mod(currentPosition-moves, 100)
	} else {
		newPosition = mod(currentPosition+moves, 100)
	}

	return newPosition, newPosition == 0
}

func getPosition2(move string, currentPosition int) (int, int) {
	direction := move[0]

	moves, err := strconv.Atoi(move[1:])
	if err != nil {
		panic("String not correct order")
	}

	newPosition := 0
	total := 0

	if direction == 'L' {
		if currentPosition == 0 {
			total = moves / 100
		} else if moves > currentPosition {
			total = (moves-currentPosition)/100 + 1
		} else if moves == currentPosition {
			total = 1
		}

		newPosition = mod(currentPosition-moves, 100)
	} else {
		total = (currentPosition + moves) / 100
		newPosition = mod(currentPosition+moves, 100)
	}

	if total > 0 {
		fmt.Printf("current position %d, move %d, zeroes %d\n", newPosition, moves, total)
	}
	return newPosition, total
}

func mod(a, b int) int {
	r := a % b
	if (r < 0 && b > 0) || (r > 0 && b < 0) {
		return r + b
	}
	return r
}
