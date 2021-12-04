package main

import (
	"fmt"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	process2(lines)
}

func process(lines []string) {
	var depth int
	var horizontal int
	for _, instruction := range lines {
		instructionParts := strings.Split(instruction," ")
		direction := instructionParts[0]
		amount := helpers.StringToInt(instructionParts[1])
		switch direction {
		case "forward":
			horizontal +=amount
		case "down":
			depth+=amount
		case "up":
			depth-=amount
		}

	}
	fmt.Println("Depth",depth)
	fmt.Println("Distance", horizontal)


}

func process2(lines []string) {
	var depth int
	var horizontal int
	var aim int
	for _, instruction := range lines {
		instructionParts := strings.Split(instruction," ")
		direction := instructionParts[0]
		amount := helpers.StringToInt(instructionParts[1])
		switch direction {
		case "forward":
			horizontal +=amount
			depth+=aim*amount
		case "down":
			aim+=amount
		case "up":
			aim-=amount
		}

	}
	fmt.Println("Depth",depth)
	fmt.Println("Distance", horizontal)



}