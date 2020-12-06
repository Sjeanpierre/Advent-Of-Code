package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var lines []string
	var count int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Could not open input file", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineData := scanner.Text()
		if lineData == "" {
			//We've reached the end of our group, do calculations with data
			//count += calculateAnyYesAnswers(lines)
			count += calculateAllYesAnswers(lines)
			lines = []string{}
			continue
		}
		lines = append(lines, scanner.Text())
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
			tracker[string(char)] = tracker[string(char)]+1
		}
	}

	for _, count := range tracker {
		if count == len(lines) {
			countAllYes += 1
		}
	}

	return countAllYes
}