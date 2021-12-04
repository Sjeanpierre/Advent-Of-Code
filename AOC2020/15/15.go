package main

import (
	"fmt"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var PART2 = true

func main() {
	var t []int
	lines := helpers.LoadFileLines("./input.txt")
	for _, v := range strings.Split(lines[0], ",") {
		t = append(t, helpers.StringToInt(v))
	}
	process(t)
}

func process(lines []int) {
	findPos := 2020
	if PART2 {
		findPos = 30000000
	}

	var tracker = make(map[int]int,findPos+1)
	for i, v := range lines {
		tracker[v] = i + 1
	}
	num := 0
	for x := len(lines)+1; x > 0; x++ {
		if x > findPos-3 && x < findPos+3 {
			fmt.Println(x, ":", num)
		}

		if x > findPos+3 {
			break
		}
		lastPos, ok := tracker[num]
		if !ok {
			tracker[num] = x
			num = 0
			continue
		}
		tracker[num]=x
		num = x-lastPos
	}
}
