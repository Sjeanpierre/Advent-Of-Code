package main

import (
	"fmt"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./sample_input.txt")
	//lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
		//do something here with lines
	}
}
