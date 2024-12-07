package main

import (
	"flag"
	"fmt"

	"github.com/hra42/aoc-2024/internal/day1"
	"github.com/hra42/aoc-2024/internal/day2"
)

func main() {
	day := flag.Int("day", 1, "Select the day to run (1-24)")
	part := flag.Int("part", 1, "Select the part to run (1-2)")
	flag.Parse()

	switch *day {
	case 1:
		var result int
		var err error
		if *part == 1 {
			result, err = day1.CalculateSolutionPart1()
		} else {
			result, err = day1.CalculateSolutionPart2()
		}
		if err != nil {
			fmt.Printf("Error calculating solution for day 1 part %d: %v\n", *part, err)
			return
		}
		fmt.Printf("Day 1 Part %d: %d\n", *part, result)
	case 2:
		var result int
		var err error
		if *part == 1 {
			result, err = day2.CalculateSolutionPart1()
		} else {
			result, err = day2.CalculateSolutionPart2()
		}
		if err != nil {
			fmt.Printf("Error calculating solution for day 2 part %d: %v\n", *part, err)
			return
		}
		fmt.Printf("Day 2 Part %d: %d\n", *part, result)
	default:
		fmt.Println("Day not implemented yet")
	}
}
