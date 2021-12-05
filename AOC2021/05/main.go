package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var V = make(map[string]int)

func main() {
	lines := helpers.LoadFileLines("./sample_input.txt")
	process(lines)
}

func process(lines []string) {
	fmt.Println(lines)
	for _, line := range lines {
		walkCoordinates(line)
	}
	var crossedThreshold int
	for _, count := range V {
		if count > 1 {
			crossedThreshold++
		}
	}
	fmt.Println(crossedThreshold)
}

func parseCoordinates(coor string) (int, int, int, int) {
	coor = strings.ReplaceAll(coor, "->", ",")
	coor = strings.ReplaceAll(coor, " ", "")
	coorParts := strings.Fields(coor)
	coorParts = strings.Split(coor, ",")
	xStart, yStart, xEnd, yEnd := helpers.StringToInt(coorParts[0]),
		helpers.StringToInt(coorParts[1]),
		helpers.StringToInt(coorParts[2]),
		helpers.StringToInt(coorParts[3])
	return xStart, yStart, xEnd, yEnd
}

func walkCoordinates(coor string) {
	xStart, yStart, xEnd, yEnd := parseCoordinates(coor)
	xs := []int{xStart, xEnd}
	sort.Ints(xs)
	startX, finishX := xs[0], xs[1]
	var xLine []int
	for x := startX; x <= finishX; x++ {
		xLine = append(xLine, x)
	}
	ys := []int{yStart, yEnd}
	sort.Ints(ys)
	startY, finishY := ys[0], ys[1]
	var yLine []int
	for y := startY; y <= finishY; y++ {
		yLine = append(yLine, y)
	}

	if len(xLine) == len(yLine) {
		xSeen := make(map[int]bool)
		ySeen := make(map[int]bool)
		for _, xcor := range xLine {
			for _, ycor := range yLine {
				if !(xcor == xStart || xcor == xEnd) {
					if !(ycor == yStart || ycor == yEnd) {
						_, xok := xSeen[xcor]
						_, yok := ySeen[ycor]
						if !xok && !yok {
							V[fmt.Sprintf("%d-%d", xcor, ycor)]++
							xSeen[xcor] = true
							ySeen[ycor] = true
						}

					}
				}
			}
		}
	} else {
		for _, xcor := range xLine {
			for _, ycor := range yLine {
				V[fmt.Sprintf("%d-%d", xcor, ycor)]++
			}
		}
	}
}
