package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var opening = map[string]bool{"[": true, "(": true, "{": true, "<": true}
var closingMatches = map[string]string{"]": "[", ")": "(", "}": "{", ">": "<"}
var openingMatches = map[string]string{"[": "]", "(": ")", "{": "}", "<": ">"}
var mismatchPoints = map[string]int{"]": 57, ")": 3, "}": 1197, ">": 25137}
var corruptedLines = map[int]bool{}
var closingPoints = map[string]int{"]": 2, ")": 1, "}": 3, ">": 4}

func main() {
	//lines := helpers.LoadFileLines("./sample_input.txt")
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
	process2(lines)
}

func process(lines []string) {
	var score int
LineLoop:
	for lineIdx, line := range lines {
		var stack []string
		for idx, char := range strings.Split(line, "") {

			if opening[char] {
				stack = append(stack, char) //push onto stack
				continue
			}
			lastPos := len(stack) - 1
			if stack[lastPos] == closingMatches[char] {
				stack = stack[:lastPos] //pop from stack
				continue
			}
			fmt.Printf("Bad char detected at position %d on line %d. Char was %s\n", idx, lineIdx, char)
			corruptedLines[lineIdx] = true
			score += mismatchPoints[char]
			continue LineLoop
		}
	}
	fmt.Println(score)
}

func process2(lines []string) {
	var scores []int

	for lineIdx, line := range lines {
		var score int
		if corruptedLines[lineIdx] {
			continue
		}
		var neededClosings []string
		var stack []string
		for _, char := range strings.Split(line, "") {

			if opening[char] {
				stack = append(stack, char) //push onto stack
				continue
			}
			lastPos := len(stack) - 1
			if stack[lastPos] == closingMatches[char] {
				stack = stack[:lastPos] //pop from stack
				continue
			}
		}
		for i := len(stack) - 1; i >= 0; i-- {
			unclosed := stack[i]
			neededClosings = append(neededClosings, openingMatches[unclosed])
			score = (score * 5) + closingPoints[openingMatches[unclosed]]
		}
		fmt.Println(neededClosings)
		scores = append(scores, score)
	}
	fmt.Println(scores)
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
