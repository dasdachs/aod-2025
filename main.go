package main

import (
	"flag"
	"fmt"
	"os"
)

// Solver is an interface that all Advent of Code day solutions must implement.
type Solver interface {
	SolvePart1(inputFile string) (string, error)
	SolvePart2(inputFile string) (string, error)
}

var solvers = make(map[string]Solver)

func registerSolver(name string, solver Solver) {
	solvers[name] = solver
}

func main() {
	task := flag.String("task", "", "The Advent of Code task to run (e.g., 'day1')")
	part := flag.Int("part", 0, "The part of the task to run (1 or 2)")
	inputFile := flag.String("input", "", "Path to the input file for the task")
	flag.Parse()

	if *task == "" {
		fmt.Println("Error: --task argument is required.")
		flag.Usage()
		os.Exit(1)
	}
	if *part != 1 && *part != 2 {
		fmt.Println("Error: --part must be 1 or 2.")
		flag.Usage()
		os.Exit(1)
	}
	if *inputFile == "" {
		fmt.Println("Error: --input argument is required.")
		flag.Usage()
		os.Exit(1)
	}

	solver, ok := solvers[*task]
	if !ok {
		fmt.Printf("Error: Unknown task '%s'.\n", *task)
		os.Exit(1)
	}

	var result string
	var err error

	switch *part {
	case 1:
		result, err = solver.SolvePart1(*inputFile)
	case 2:
		result, err = solver.SolvePart2(*inputFile)
	}

	if err != nil {
		fmt.Printf("Error running task %s part %d: %v\n", *task, *part, err)
		os.Exit(1)
	}

	fmt.Printf("Result for %s part %d: %s\n", *task, *part, result)
}
