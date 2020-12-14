package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {
	batchSize := 25
	x, y := 0, batchSize
	for oo := 0; oo < len(lines); oo++ {
		previousSet := lines[x:y]
		if !match(previousSet, lines, y) {
			log.Printf("Could not find match for %s at index %d", lines[y], y+1)
			break
		}
		x++
		y++
	}
	findSeries(helpers.StringToInt(lines[y]), lines)

	fmt.Println("")

}

func match(section []string, lines []string, y int) bool {
	searchInt := y
	numToFind := helpers.StringToInt(lines[searchInt])
	for _, s1 := range section {
		for _, s2 := range section {
			if numToFind == helpers.StringToInt(s1)+helpers.StringToInt(s2) && s1 != s2 {
				fmt.Printf("(%s,%s)==%s\n", s1, s2, lines[searchInt])
				return true
			}
		}
	}
	return false
}

func findSeries(numToFind int, numList []string) {
	start := 0
	counter := 0
	for oo := 0; oo < len(numList); oo++ {
		counter += helpers.StringToInt(numList[start+oo])
		if counter == numToFind {
			var series []int
			for _, num := range numList[start : start+oo] {
				series = append(series, helpers.StringToInt(num))
			}
			sort.Ints(series)
			low := series[0]
			high := series[len(series)-1]
			log.Println("BINGO-Part-2", start, "distance", oo, low+high)
			break
		}
		if counter > numToFind {
			start++
			counter = 0
			oo = 0
			continue
		}
	}
}
