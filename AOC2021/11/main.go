package main

import (
	"fmt"
	"log"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var grid = helpers.NewGrid()
var octoFlashTracker = map[helpers.Coor]bool{}
var totalFlashes int
var CYCLES = 10000

func main() {
	//lines := helpers.LoadFileLines("./sample_input.txt")
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {

	grid.PopulateWithLines(lines)

	grid.PrintGrid(len(lines), len(lines[0]))
	for cycle := 0; cycle < CYCLES; cycle++ {
		octoFlashTracker = map[helpers.Coor]bool{} //reset tracker each cycle
		increaseAll(len(lines), len(lines[0]))
		for row := 0; row < len(lines); row++ {
			for col := 0; col < len(lines[0]); col++ {
				c := helpers.Coor{Row: row, Col: col}
				if grid.GetCell(c) > 9 {
					flash(c)
					octoSearch(c)
				}
			}
		}
		//fmt.Println("---------")
		//grid.PrintGrid(len(lines), len(lines[0]))
		if cycle+1 == 100 {
			fmt.Print("----------\n", "Total Flashes: ", totalFlashes, "\n")
		}
		if len(octoFlashTracker) == grid.Size() {
			log.Fatalln("All flashed!!!", "Cycle:", cycle+1)
		}
	}
	fmt.Print("----------\n", "Total Flashes: ", totalFlashes, "\n")
}

func increaseAll(rows, cols int) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			c := helpers.Coor{Row: row, Col: col}
			grid.Incr(c)
		}
	}
}

func octoSearch(c helpers.Coor) {
	octoMembers := []helpers.Coor{c}
	for x := 0; ; x++ {
		var countFound int
		coords := octoMembers[x]
		down, up, left, right, dl, dr, ul, ur := coords.MapNeighborswDiag()
		for _, neighbor := range []helpers.Coor{down, up, left, right, dl, dr, ul, ur} {
			neighborValue, ok := grid.Locate(neighbor)
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
					grid.Incr(neighbor)
				}
			} else {
				grid.Incr(neighbor)
			}
		}
		if countFound == 0 && len(octoMembers) == x+1 {
			break
		}
	}
}

func flash(c helpers.Coor) {
	grid.SetCell(c, 0)
	totalFlashes++
	octoFlashTracker[c] = true
}
