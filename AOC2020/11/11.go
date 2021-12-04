package main

import (
	"fmt"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

//Todo, they want a mark and sweep approach to this.

var PART2 bool

type coordinates struct {
	X, Y int
}

func main() {
	lines := helpers.LoadFileLines("./input.txt")

	var chart = make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		chart[i] = make([]string, len(lines[0]))
	}
	for y, line := range lines {
		for x, char := range line {
			chart[y][x] = string(char)
		}
	}
	PART2 = true
	processPartX(chart)
}

func processPartX(chart [][]string) {
	for z := 0; z > -1; z++ {
		seatChanges := make(map[coordinates]bool)
		for y := 0; y < len(chart); y++ {
			for x := 0; x < len(chart[0]); x++ {
				switch chart[y][x] {
				case "L":
					if canSit(chart, coordinates{Y: y, X: x}) {
						seatChanges[coordinates{X: x, Y: y}] = true
					}
				case "#":
					if shouldClear(chart, coordinates{Y: y, X: x}) {
						seatChanges[coordinates{X: x, Y: y}] = true
					}
				}
			}
		}

		if len(seatChanges) == 0 {
			total := 0
			for _, y := range chart {
				total += strings.Count(strings.Join(y, ""), "#")
			}
			fmt.Println(total)
			break
		}

		for c, _ := range seatChanges {
			chart[c.Y][c.X] = flip(chart[c.Y][c.X])
		}
	}
}

func flip(i string) string {
	if i == "#" {
		return "L"
	}
	return "#"
}

func canSit(chart [][]string, c coordinates) bool {
	neighbors := getAdjacentCoords(c, chart)
	if len(neighbors) == 0 {
		return true
	}
	//fmt.Println(char,neighbors)
	for _, nc := range neighbors {
		if chart[nc.Y][nc.X] == "#" {
			return false
		}
	}

	return true
}

func shouldClear(chart [][]string, c coordinates) bool {
	limit := 4
	if PART2 {
		limit = 5
	}
	neighbors := getAdjacentCoords(c, chart)
	counter := 0
	for _, nc := range neighbors {
		if chart[nc.Y][nc.X] == "#" {
			counter++
			if counter >= limit {
				return true
			}
		}
	}
	return false
}

func getAdjacentCoords(c coordinates, chart [][]string) []coordinates {
	maxCoords := coordinates{
		X: len(chart[0]) - 1,
		Y: len(chart) - 1,
	}
	level := 1
	//Don't forget to address chart as yx
	directions := []rune{'⬆', '⬇', '⬅', '➡', '↗', '↖', '↘', '↙'}
	var around []coordinates

	for _, direction := range directions {
		if PART2 {
			coord, ok := getFirstVisibleSeat(direction, c, chart)
			if !ok {
				continue
			}
			around = append(around, coord)
			continue
		}
		around = append(around, c.move(direction, level))
	}

	var returnSet []coordinates

	for _, coor := range around {
		if coor.Y < 0 || coor.X < 0 || coor.X > maxCoords.X || coor.Y > maxCoords.Y {
			continue
		}
		returnSet = append(returnSet, coor)
	}
	if len(returnSet) == 0 {
		fmt.Println("")
	}
	return returnSet
}

func getFirstVisibleSeat(direction rune, startingCoords coordinates, chart [][]string) (coordinates, bool) {
	maxCoords := coordinates{
		X: len(chart[0]) - 1,
		Y: len(chart) - 1,
	}
	for x := 1; x <= len(chart)+10; x++ {
		coor := startingCoords.move(direction, x)
		if coor.Y < 0 || coor.X < 0 || coor.X > maxCoords.X || coor.Y > maxCoords.Y {
			break
		}

		if chart[coor.Y][coor.X] != "." {
			return coor, true
		}
	}
	return coordinates{}, false
}

func (c coordinates) move(direction rune, count int) coordinates {
	switch direction {
	case '⬆':
		return coordinates{X: c.X, Y: c.Y - count}
	case '⬇':
		return coordinates{X: c.X, Y: c.Y + count}
	case '⬅':
		return coordinates{X: c.X - count, Y: c.Y}
	case '➡':
		return coordinates{X: c.X + count, Y: c.Y}
	case '↗':
		return coordinates{X: c.X + count, Y: c.Y - count}
	case '↖':
		return coordinates{X: c.X - count, Y: c.Y - count}
	case '↘':
		return coordinates{X: c.X + count, Y: c.Y + count}
	case '↙':
		return coordinates{X: c.X - count, Y: c.Y + count}
	}
	return coordinates{}
}
