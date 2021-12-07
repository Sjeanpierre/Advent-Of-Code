package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var B = true

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {
	var positions []int
	for _, position := range strings.Split(lines[0], ",") {
		positions = append(positions, helpers.StringToInt(position))
	}
	bestFuel := 10000000000
	var bestPosition int
	for _, position := range positions {
		fuelCost := calculate(positions, position)
		if fuelCost < bestFuel {
			bestFuel = fuelCost
			bestPosition = position
		}
		fmt.Println(fuelCost)
	}
	fmt.Println(bestPosition)
	fmt.Println(bestFuel)
}

func calculate(positions []int, candidate int) int {
	fmt.Println("Calculating position", candidate)
	var fuel int
	for _, position := range positions {
		if B {
			fuel += exoFuel(int(math.Abs(float64(position) - float64(candidate))))
		} else {
			fuel += int(math.Abs(float64(position) - float64(candidate)))
		}

	}

	return fuel
}

func exoFuel(f int) int {
	var c int
	for i := 0; i <= f; i++ {
		c += i
	}
	return c
}
