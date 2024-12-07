package day2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// isValidSequence checks if a sequence of numbers follows the rules:
// 1. All increasing or all decreasing
// 2. Adjacent numbers differ by 1-3
func isValidSequence(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}

	// Determine if sequence should be increasing or decreasing
	isIncreasing := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		// Check if difference is between 1 and 3
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}

		// Check if sequence maintains direction
		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}

	return true
}

// isValidSequenceWithDampener checks if a sequence can be made valid by removing one number
func isValidSequenceWithDampener(numbers []int) bool {
	// First check if it's already valid without removing any number
	if isValidSequence(numbers) {
		return true
	}

	// Try removing each number once
	for i := range numbers {
		// Create a new slice without the current number
		dampened := make([]int, 0, len(numbers)-1)
		dampened = append(dampened, numbers[:i]...)
		dampened = append(dampened, numbers[i+1:]...)

		if isValidSequence(dampened) {
			return true
		}
	}

	return false
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// parseLine converts a string of space-separated numbers into a slice of integers
func parseLine(line string) ([]int, error) {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))

	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("error parsing number: %v", err)
		}
		numbers[i] = num
	}

	return numbers, nil
}

func CalculateSolutionPart1() (int, error) {
	inputPath := filepath.Join("input", "day2", "part1", "input.txt")
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		numbers, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		if isValidSequence(numbers) {
			safeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return safeReports, nil
}

func CalculateSolutionPart2() (int, error) {
	inputPath := filepath.Join("input", "day2", "part2", "input.txt")
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		numbers, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		if isValidSequenceWithDampener(numbers) {
			safeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return safeReports, nil
}
