package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

type coor struct {
	row, col int
}

var grid = map[coor]int{}
var basinTracker = map[coor]bool{}

func main() {
	//lines := helpers.LoadFileLines("./sample_input.txt")
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {
	for idxRow, row := range lines {
		for idxCol, column := range strings.Split(row, "") {
			grid[coor{row: idxRow, col: idxCol}] = helpers.StringToInt(column)
		}
	}

	var tracker []coor
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[0]); col++ {
			c := coor{row: row, col: col}
			if lowest(c) {
				tracker = append(tracker, c)
			}
		}
	}
	var basinSizes []int
	var basins [][]coor
	var sum int
	for _, c := range tracker {
		sum += grid[c] + 1
		foundBasinMembers := basinSearch(c)
		basins = append(basins, foundBasinMembers)
		basinSizes = append(basinSizes, len(foundBasinMembers))
	}
	sort.Ints(basinSizes)
	top3Product := 1
	for _, value := range basinSizes[len(basinSizes)-3:] {
		top3Product *= value
	}
	fmt.Println(sum)
	fmt.Println(top3Product)
}

func lowest(c coor) bool {
	down, up, left, right := mapNeighbors(c)
	cVal := grid[c]
	if cVal == 9 { //even if surrounded by other 9s a 9 is still a high point
		return false
	}
	for _, neighbor := range []coor{down, up, left, right} {
		neighborValue, ok := grid[neighbor]
		if !ok { //Out of bounds
			continue
		}
		if cVal > neighborValue { //A neighbor value is lower than the current value
			return false
		}
	}
	return true
}

func basinSearch(c coor) []coor {
	basinMembers := []coor{c}
	basinTracker[c] = true
	for x := 0; ; x++ {
		var countFound int
		coords := basinMembers[x]
		down, up, left, right := mapNeighbors(coords)
		for _, neighbor := range []coor{down, up, left, right} {
			neighborValue, ok := grid[neighbor]
			if !ok || neighborValue == 9 { //Out of bounds or highest elevation found
				continue
			}
			if !basinTracker[neighbor] {
				basinMembers = append(basinMembers, neighbor)
				basinTracker[neighbor] = true
				countFound++
			}
		}
		if countFound == 0 && len(basinMembers) == x+1 {
			break
		}
	}
	return basinMembers
}

func mapNeighbors(c coor) (coor, coor, coor, coor) {
	return coor{row: c.row + 1, col: c.col}, coor{row: c.row - 1, col: c.col},
		coor{row: c.row, col: c.col - 1}, coor{row: c.row, col: c.col + 1}
}
