package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	process(lines)
}

func process(lines []string) {
	accumulator := 0
	x := 0
	var seen = make(map[int]bool)

	for  {
		if seen[x] {
			break
		}
		seen[x]=true
		parts := strings.Split(lines[x]," ")
		instruction, dist := parts[0], stringToInt(parts[1])
		switch instruction {
		case "nop":
			x++
		case "acc":
			accumulator += dist
			x++
		case "jmp":
			x+=dist
		}
	}

	fmt.Println(accumulator)



}

func stringToInt(s string) int {
	i,_ := strconv.Atoi(s)
	return i
}