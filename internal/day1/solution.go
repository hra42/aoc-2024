package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// solutionPart1 calculates the total distance between sorted lists (original solution)
func solutionPart1(leftList, rightList []int) int {
	// Create copies of the slices to avoid modifying the original data
	left := make([]int, len(leftList))
	right := make([]int, len(rightList))
	copy(left, leftList)
	copy(right, rightList)

	sort.Ints(left)
	sort.Ints(right)

	// Calculate total distance
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := int(math.Abs(float64(left[i] - right[i])))
		totalDistance += distance
	}

	return totalDistance
}

// solutionPart2 calculates the similarity score between two lists
func solutionPart2(leftList, rightList []int) int {
	// Create a map to count occurrences in rightList
	rightCounts := make(map[int]int)
	for _, num := range rightList {
		rightCounts[num]++
	}

	// Calculate similarity score
	totalScore := 0
	for _, leftNum := range leftList {
		// Multiply the left number by its occurrences in the right list
		totalScore += leftNum * rightCounts[leftNum]
	}

	return totalScore
}

func CalculateSolutionPart1() (int, error) {
	// Open the input file from part1 subfolder
	inputPath := filepath.Join("input", "day1", "part1", "input.txt")
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Parse the input file
	leftList, rightList, err := parseInput(file)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}

	// Calculate the result
	result := solutionPart1(leftList, rightList)
	return result, nil
}

func CalculateSolutionPart2() (int, error) {
	// Open the input file from part2 subfolder
	inputPath := filepath.Join("input", "day1", "part2", "input.txt")
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Parse the input file
	leftList, rightList, err := parseInput(file)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}

	// Calculate the result
	result := solutionPart2(leftList, rightList)
	return result, nil
}

func parseInput(file *os.File) ([]int, []int, error) {
	scanner := bufio.NewScanner(file)
	var leftList, rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing left number: %v", err)
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing right number: %v", err)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return leftList, rightList, nil
}
