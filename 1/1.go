package main

import (
	"log"
	"os"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	var numbers []int
	for _, line := range lines {
		numbers = append(numbers,helpers.StringToInt(line))
	}
	//part1(numbers)
	part2(numbers)

}

func part1(lines []int) {
	for _, num := range lines {
		for _, numb := range lines {
			if (num + numb) == 2020 {
				log.Println(num * numb)
				os.Exit(0)
			}

		}
	}
}

func part2(lines []int) {
	for _, num1 := range lines {
		for _, num2 := range lines {
			for _, num3 := range lines {
				if (num1 + num2 + num3) == 2020 {
					log.Println(num1 * num2 * num3)
					os.Exit(0)
				}
			}
		}
	}
}
