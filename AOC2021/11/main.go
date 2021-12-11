package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

type coor struct {
	row, col int
}

var grid = map[coor]int{}
var octoFlashTracker = map[coor]bool{}
var totalFlashes int
var CYCLES = 10000

func main() {
	//lines := helpers.LoadFileLines("./sample_input.txt")
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {
	for rowIdx, line := range lines {
		for colIdx, value := range strings.Split(line, "") {
			grid[coor{row: rowIdx, col: colIdx}] = helpers.StringToInt(value)
		}
	}

	printGrid(len(lines), len(lines[0]))
	for cycle := 0; cycle < CYCLES; cycle++ {
		octoFlashTracker = map[coor]bool{} //reset tracker each cycle
		increaseAll(len(lines), len(lines[0]))
		for row := 0; row < len(lines); row++ {
			for col := 0; col < len(lines[0]); col++ {
				c := coor{row: row, col: col}
				if grid[c] > 9 {
					flash(c)
					octoSearch(c)
				}
			}
		}
		//fmt.Println("---------")
		//printGrid(len(lines), len(lines[0]))
		if cycle+1 == 100 {
			fmt.Print("----------\n", "Total Flashes: ", totalFlashes, "\n")
		}
		if len(octoFlashTracker) == len(grid) {
			log.Fatalln("All flashed!!!", "Cycle:", cycle+1)
		}
	}
	fmt.Print("----------\n", "Total Flashes: ", totalFlashes, "\n")
}

func increaseAll(rows, cols int) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			c := coor{row: row, col: col}
			grid[c]++
		}
	}
}

func octoSearch(c coor) {
	octoMembers := []coor{c}
	for x := 0; ; x++ {
		var countFound int
		coords := octoMembers[x]
		down, up, left, right, dl, dr, ul, ur := mapNeighbors(coords)
		for _, neighbor := range []coor{down, up, left, right, dl, dr, ul, ur} {
			neighborValue, ok := grid[neighbor]
			if octoFlashTracker[neighbor] {
				continue
			}
			if !ok { //Out of bounds
				continue
			}
			if neighborValue+1 > 9 {
				if !octoFlashTracker[neighbor] {
					flash(neighbor) //reset
					octoMembers = append(octoMembers, neighbor)
					countFound++
				} else {
					grid[neighbor]++
				}
			} else {
				grid[neighbor]++
			}
		}
		if countFound == 0 && len(octoMembers) == x+1 {
			break
		}
	}
}

func flash(c coor) {
	grid[c] = 0
	totalFlashes++
	octoFlashTracker[c] = true
}

func mapNeighbors(c coor) (coor, coor, coor, coor, coor, coor, coor, coor) {
	return coor{row: c.row + 1, col: c.col}, //down
		coor{row: c.row - 1, col: c.col}, //up
		coor{row: c.row, col: c.col - 1}, //left
		coor{row: c.row, col: c.col + 1}, //right
		coor{row: c.row + 1, col: c.col - 1}, //down-left
		coor{row: c.row + 1, col: c.col + 1}, //down-right
		coor{row: c.row - 1, col: c.col - 1}, //up-left
		coor{row: c.row - 1, col: c.col + 1} //up-right

}

func printGrid(rows, cols int) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			fmt.Print(grid[coor{row: row, col: col}])

		}
		fmt.Print("\n")
	}
}
