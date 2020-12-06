package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

var maxRow = 127
var maxSeat = 8

func main() {
	var lines []string

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Could not open input file", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//fmt.Println(len(lines))
	s, r := part1(lines)
	a := part2()
	fmt.Println(s, a, r)
	var lol = make(map[int]bool)
	var missing []int
	for _, seatID := range s {
		lol[seatID] = true
	}

	for _, seatID := range a {
		_, ok := lol[seatID]
		if !ok {
			missing = append(missing, seatID)
		}
	}
	fmt.Println(missing)
}

func part1(lines []string) ([]int, []int) {
	var seatIDs []int
	var rows []int
	for _, lineData := range lines {
		rowMax := maxRow
		rowMin := 0
		colMax := maxSeat
		colMin := 0
		rowCoordinates := lineData[0:7]
		var finalRow int
		var finalCol int
		seatCoordinates := lineData[len(lineData)-3:]
		for i, direction := range rowCoordinates {
			dir := string(direction)
			fmt.Println(i)
			switch dir {
			case string('F'):
				mid := math.Round(float64((rowMin + rowMax) / 2))
				diff := rowMax - rowMin
				rowMax = int(mid)
				if i == len(rowCoordinates)-1 {
					if diff > 1 {
						finalRow = int(mid)
					} else {
						finalRow = rowMin
					}
				}
			case string('B'):
				mid := math.Round(float64((rowMin + rowMax) / 2))
				rowMin = int(mid)
				if i == len(rowCoordinates)-1 {
					if rowMax-rowMin > 1 {
						finalRow = int(mid)
					}
					finalRow = rowMax
				}
			}
		}
		for i, direction := range seatCoordinates {
			dir := string(direction)
			switch dir {
			case "L":
				mid := math.Round(float64((colMin + colMax) / 2))
				colMax = int(mid)
				if i == len(seatCoordinates)-1 {
					finalCol = colMin
				}
			case "R":
				mid := math.Round(float64((colMin + colMax) / 2))
				colMin = int(mid)
				if i == len(seatCoordinates)-1 {
					finalCol = colMax - 1
				}
			}

		}

		if (finalRow*8 + finalCol) == 86 {
			fmt.Println("")
		}

		seatIDs = append(seatIDs, ((finalRow * 8) + finalCol))
		rows = append(rows, finalRow)
	}
	sort.Ints(seatIDs)
	sort.Ints(rows)
	return seatIDs, rows
	//fmt.Println(seatIDs)
}

func part2() []int {
	var allSeats []int
	for x := 9; x <= 119; x++ {
		for y := 0; y <= 7; y++ {
			allSeats = append(allSeats, (x*8 + y))
		}
	}
	sort.Ints(allSeats)
	return allSeats
}
