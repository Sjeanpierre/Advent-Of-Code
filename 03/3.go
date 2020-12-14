package main

import (
	"fmt"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	part1(lines)
	part2(lines, 1, 1)
	part2(lines, 3, 1)
	part2(lines, 5, 1)
	part2(lines, 7, 1)
	part2(lines, 1, 2)
}

func part1(lines []string) {
	var counter int
	x := 0
	for i, line := range lines {
		if i == 0 {
			continue
		}
		x += 3
		pos := x % 31
		if line[pos] == '#' {
			counter += 1
		}
	}
	fmt.Println(counter)
}

func part2(lines []string, right, down int) {
	var x int
	var counter int
	for y := 0; y <= len(lines)-1; {
		if y == 0 {
			y += down
			continue
		}

		x += right
		pos := x % 31
		if lines[y][pos] == '#' {
			counter += 1
		}
		y += down
	}
	fmt.Println("Slope:", "right:", right, "down:", down, "--", counter)
}
