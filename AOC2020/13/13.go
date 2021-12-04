package main

import (
	"fmt"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)

}

func process(lines []string) {
	timeStamp := helpers.StringToInt(lines[0])
	var buses []int
	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {continue}
		buses = append(buses, helpers.StringToInt(bus))
	}
	for _, bus := range buses {
		fmt.Println(timeStamp,"%",bus,"="," mod:",timeStamp%bus,"div:",bus+(timeStamp/bus)*bus)
	}

	//Part 2 wtf?
	//https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Search_by_sieving?
	routeIDs := make(map[int]int)
	for i, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {continue}
		routeIDs[helpers.StringToInt(bus)]=i
	}
	l := 0
	product := 1
	for busFreq, pos := range routeIDs {
		for (l+pos)%busFreq != 0 {
			l += product
		}
		product *= busFreq
	}
	fmt.Println(l)
}