package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var InputLinePattern = regexp.MustCompile(`^(\d+)-(\d+)\s(\w):\s(\w+)`)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	part2(lines)
}

func parseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}

func part1(lines []string) {
	var counter int
	for _, passwordLine := range lines {
		matches := InputLinePattern.FindAllStringSubmatch(passwordLine, -1)
		if len(matches[0]) != 5 {
			continue
		}
		min := parseInt(matches[0][1])
		max := parseInt(matches[0][2])
		xx := matches[0][3]
		password := matches[0][4]
		occ := strings.Count(password, xx)
		if occ >= min && occ <= max {
			counter += 1
		}
	}
	log.Println(counter)
}

func part2(lines []string) {
	var counter int
	for _, passwordLine := range lines {
		matches := InputLinePattern.FindAllStringSubmatch(passwordLine, -1)
		if len(matches[0]) != 5 {
			continue
		}
		pos1 := parseInt(matches[0][1])
		pos2 := parseInt(matches[0][2])
		xx := matches[0][3][0]
		password := matches[0][4]
		if password[pos1-1] == password[pos2-1] {
			continue
		} else {
			if password[pos1-1] == xx || password[pos2-1] == xx {
				counter += 1
			}
		}
	}
	log.Println(counter)
}
