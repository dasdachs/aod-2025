package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseFileToStringArray(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return strings.Split(string(data), "\n"), nil
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Sum[I int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](values []I) I {
	var res I

	for _, n := range values {
		res = res + n
	}

	return res
}
