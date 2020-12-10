package main

import (
	"fmt"
	"sort"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("input.txt")
	var a []int
	for _, b := range lines {
		a = append(a,helpers.StringToInt(b))
	}
	processPart1(a)
}

func processPart1(lines []int)  {
	currentPos :=0
	sort.Ints(lines)
	adaptersIntervals := map[int]int{}

	for x:=0;x<len(lines);x++ {
		dist := lines[x]-currentPos
		adaptersIntervals[dist]++
		currentPos = lines[x]
	}
	//Add the main adapter
	adaptersIntervals[3]++
	fmt.Println(adaptersIntervals[3]*adaptersIntervals[1])
}