package main

import (
	"flag"
	"fmt"
	"os"

	"firstbulb.xyz/aoc-2025/solvers"
)

func main() {
	day := flag.Int("day", 0, "Required: the day, e.g. 1.")
	part := flag.Int("part", 1, "Optional: the part of the challenge.")
	aod := flag.String("aod", "global", "Optional: which advent od code you want to solve, global or inova.")
	dataSource := flag.String("data", "", "Optional: The path to the data source file, default to the challenge data.")
	helpFlag := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	fmt.Printf("Solving challenge for day: %d (part: %d) in the %s advent of code challenge.\n", *day, *part, *aod)
	fmt.Println("------------------------------")

	if *dataSource != "" {
		fmt.Printf("Using data source: %s\n", *dataSource)
		fmt.Println("------------------------------")
	}

	fmt.Println("Godspeed!")
	fmt.Println("------------------------------")

	if *aod == "global" {
		executeGlobalSolver(*day, *part, *dataSource)
	}
}

func printHelp() {
	fmt.Println("Usage: [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func executeGlobalSolver(solver, part int, dataSource string) {
	switch solver {
	case 1:
		if part == 1 {
			solvers.Day1Part1(dataSource)
		} else {
			solvers.Day1Part2(dataSource)
		}
	case 2:
		if part == 1 {
			solvers.Day2Part1(dataSource)
		} else {
			solvers.Day2Part2(dataSource)
		}
	case 3:
		if part == 1 {
			solvers.Day3Part1(dataSource)
		} else {
			solvers.Day3Part2(dataSource)
		}
	case 4:
		if part == 1 {
			solvers.Day4Part1(dataSource)
		} else {
			solvers.Day4Part2(dataSource)
		}
	// case 5:
	// 	if part == 1 {
	// 		solvers.Day5Part1(dataSource)
	// 	} else {
	// 		solvers.Day5Part2(dataSource)
	// 	}
	// case 6:
	// 	if part == 1 {
	// 		solvers.Day6Part1(dataSource)
	// 	} else {
	// 		solvers.Day6Part2(dataSource)
	// 	}
	// case 7:
	// 	if part == 1 {
	// 		solvers.Day7Part1(dataSource)
	// 	} else {
	// 		solvers.Day7Part2(dataSource)
	// 	}
	// case 8:
	// 	if part == 1 {
	// 		solvers.Day8Part1(dataSource)
	// 	} else {
	// 		solvers.Day8Part2(dataSource)
	// 	}
	// case 9:
	// 	if part == 1 {
	// 		solvers.Day9Part1(dataSource)
	// 	} else {
	// 		solvers.Day9Part2(dataSource)
	// 	}
	// case 10:
	// 	if part == 1 {
	// 		solvers.Day10Part1(dataSource)
	// 	} else {
	// 		solvers.Day10Part2(dataSource)
	// 	}
	// case 11:
	// 	if part == 1 {
	// 		solvers.Day11Part1(dataSource)
	// 	} else {
	// 		solvers.Day11Part2(dataSource)
	// 	}
	// case 12:
	// 	if part == 1 {
	// 		solvers.Day12Part1(dataSource)
	// 	} else {
	// 		solvers.Day12Part2(dataSource)
	// 	}

	default:
		fmt.Println("Solver not yet implemented.")
	}
}
