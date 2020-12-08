package main

import (
	"fmt"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	var count int

	lines := helpers.LoadFileLines("./input.txt")

	var localLines []string
	for _, line := range lines {
		if line == "" {
			count += calculateAnyYesAnswers(localLines)
			//count += calculateAllYesAnswers(lines)
			localLines = []string{}
			continue
		}
		localLines = append(localLines,line)

	}
	fmt.Println(count)
}

func calculateAnyYesAnswers(lines []string) int{
	tracker := make(map[string]bool)
	for _, line := range lines {
		for _, char := range line {
			tracker[string(char)] = true
		}
	}
	return len(tracker)
}

func calculateAllYesAnswers(lines []string) int {
	var countAllYes int
	tracker := make(map[string]int)
	for _, line := range lines {
		for _, char := range line {
			tracker[string(char)]++
		}
	}

	for _, count := range tracker {
		if count == len(lines) {
			countAllYes++
		}
	}

	return countAllYes
}