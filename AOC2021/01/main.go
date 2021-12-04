package main

import (
	"fmt"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	process2(lines)
}

func process(lines []string) {
	var sequence []int
	var largerCount int
	for _, value := range lines {
		sequence = append(sequence, helpers.StringToInt(value))
	}

	for key, value := range sequence {
		if key == 0 {
			fmt.Printf("%d (N/A - no Previous number)\n",value)
			continue
		}
		if value > sequence[key-1] {
			fmt.Printf("%d (increased)\n",value)
			largerCount++
			continue
		}
		fmt.Println(value,"(decreased or no change)")
	}
	fmt.Println("Number larger", largerCount)

}

func process2(lines []string) {
	var sequence []int
	var largerCount int
	var windowSum int
	for _, value := range lines {
		sequence = append(sequence, helpers.StringToInt(value))
	}

	for key, value := range sequence {
		fmt.Println("Number larger", largerCount)
		window := value + sequence[key+1] + sequence[key+2]
		switch {
		case windowSum == 0 && key < 2 :
			fmt.Printf("%d (no previous sum)\n", window)
		case window > windowSum:
			largerCount++
			fmt.Printf("%d (increased)\n",window)
		case window < windowSum:
			fmt.Println(window,"(decreased)")
		default:
			fmt.Println(value,"(no change)")
		}
		windowSum = window
	}


}