package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var t = map[int]int{2: 1, 4: 4, 3: 7, 7: 8}

func main() {
	lines := helpers.LoadFileLines("./sample_input.txt")
	//lines := helpers.LoadFileLines("./input.txt")
	process2(lines)
}

func process(lines []string) {

	var counter int
	for _, line := range lines {
		p := strings.Trim(strings.Split(line, "|")[1], " ")
		for _, segment := range strings.Fields(p) {
			_, ok := t[len(segment)]
			if ok {
				counter++
			}
		}
	}
	fmt.Println(counter)
}

func process2(lines []string) {
	var sum int
	for _, line := range lines {
		parts := strings.Split(line, "|")
		singalPattern := strings.Trim(parts[0], " ")
		output := strings.Trim(parts[1], " ")
		sum += translate(decode(singalPattern), output)
	}
	fmt.Println(sum)

}

func translate(key map[string]int, output string) int {
	var result []int
	for _, o := range strings.Fields(output) {
		s := strings.Split(o, "")
		sort.Strings(s)
		result = append(result, key[strings.Join(s, "")])
	}
	num := fmt.Sprintf("%+v", result)
	num = strings.Replace(strings.Trim(num, "[]"), " ", "", -1)
	return helpers.StringToInt(num)
}
func decode(signal string) map[string]int {
	tracker := map[int]string{}
	for _, value := range strings.Fields(signal) {
		switch len(value) {
		case 2:
			tracker[1] = value
		case 3:
			tracker[7] = value
		case 4:
			tracker[4] = value
		case 7:
			tracker[8] = value
		default:
			fmt.Println(value)
		}
	}
	seen := map[string]int{tracker[1]: 1, tracker[7]: 7, tracker[4]: 4, tracker[8]: 8}
	for _, value := range strings.Fields(signal) {
		if seen[value] != 0 {
			continue
		}
		if isNum(tracker[1], value, 0) && isNum(tracker[8], value, 1) && isNum(tracker[4], value, 1) {
			tracker[0] = value
			seen[value] = 0
			continue
		}
		if isNum(tracker[4], value, 2) && isNum(tracker[8], value, 2) {
			tracker[2] = value
			seen[value] = 2
			continue
		}
		if isNum(tracker[4], value, 1) && isNum(tracker[8], value, 2) && isNum(tracker[7], value, 2) {
			tracker[3] = value
			seen[value] = 3
			continue
		}

		if isNum(tracker[8], value, 1) && isNum(tracker[1], value, 1) {
			tracker[6] = value
			seen[value] = 6
			continue
		}
		if isNum(tracker[4], value, 0) && isNum(tracker[7], value, 0) {
			tracker[9] = value
			seen[value] = 9
			continue
		}

		tracker[5] = value
		seen[value] = 5
	}

	seenSorted := map[string]int{}
	for key, v := range seen {
		s := strings.Split(key, "")
		sort.Strings(s)
		seenSorted[strings.Join(s, "")] = v
	}

	return seenSorted
}

func isZero(four, candidate string) bool {
	for _, char := range strings.Split(candidate, "") {
		four = strings.Replace(four, char, "", -1)
	}
	if len(four) == 1 {
		return true
	}
	return false
}

func isNum(known, candidate string, remainder int) bool {
	for _, char := range strings.Split(candidate, "") {
		known = strings.Replace(known, char, "", -1)
	}
	if len(known) == remainder {
		return true
	}
	return false
}
