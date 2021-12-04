package main

import (
	"fmt"
	"sort"

	"github.com/sjeanpierre/AOC2020/helpers"
)

type ways struct {
	defaultVal int
	m map[int]int
}

//small reimplementation of DefaultDict from Python to help
func defaultDict(defaultVal int) *ways {
	return &ways{defaultVal: defaultVal, m: make(map[int]int)}
}

func (w ways) Get(v int)  int {
	val, ok := w.m[v]
	if !ok {
		return w.defaultVal
	}
	return val
}

func (w ways) Set(pos, v int)  {
	w.m[pos]=v
}

func main() {
	lines := helpers.LoadFileLines("input.txt")
	var a []int
	for _, b := range lines {
		a = append(a, helpers.StringToInt(b))
	}
	processPart2(a)
}

func processPart1(lines []int) {
	currentPos := 0
	sort.Ints(lines)
	adaptersIntervals := map[int]int{}

	for x := 0; x < len(lines); x++ {
		dist := lines[x] - currentPos
		adaptersIntervals[dist]++
		currentPos = lines[x]
	}
	//Add the main adapter
	adaptersIntervals[3]++
	fmt.Println(adaptersIntervals[3] * adaptersIntervals[1])
}

func processPart2(lines []int) {
	lines = append(lines,0)
	sort.Ints(lines)
	//Add the last position for the adapter
	lines = append(lines,lines[len(lines)-1]+3)
	ways := defaultDict(0)
	ways.Set(0,1)
	for x:=1;x<len(lines);x++ {
		ways.Set(lines[x], ways.Get(lines[x]-1)+ways.Get(lines[x]-2)+ways.Get(lines[x]-3))
		fmt.Println("")
	}

	fmt.Println(ways.Get(lines[len(lines)-1]))
}
