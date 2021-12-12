package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var t = map[int]int{2: 1, 4: 4, 3: 7, 7: 8}

func main() {
	//lines := helpers.LoadFileLines("./sample_input.txt")
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
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
	seen := map[string]int{}
	for _, value := range strings.Fields(signal) {
		switch len(value) {
		case 2:
			tracker[1] = value
			seen[value] = 1
		case 3:
			tracker[7] = value
			seen[value] = 7
		case 4:
			tracker[4] = value
			seen[value] = 4
		case 7:
			tracker[8] = value
			seen[value] = 8
		default:
			fmt.Println(value)
		}
	}
	var upperRight string
	var lowerRight string
	var top string
	oneParts := strings.Split(tracker[1], "")
	if strings.Count(signal, oneParts[0]) == 8 {
		upperRight = oneParts[0]
		lowerRight = oneParts[1]
	} else {
		upperRight = oneParts[1]
		lowerRight = oneParts[0]
	}
	top = strings.Replace(strings.Replace(tracker[7], oneParts[0], "", -1), oneParts[1], "", -1)

	for _, value := range strings.Fields(signal) {
		if !strings.Contains(value, lowerRight) {
			tracker[2] = value
			seen[value] = 2
		}
		if sortedString(strings.Replace(tracker[8], upperRight, "", -1)) == sortedString(value) {
			tracker[6] = value
			seen[value] = 6
		}
		r := strings.Join([]string{top, tracker[4]}, "")
		if len(value) == 6 && len(deleteChars(value, r)) == 1 {
			tracker[9] = value
			seen[value] = 9
		}
		if len(value) == 5 && !strings.Contains(value, upperRight) {
			tracker[5] = value
			seen[value] = 5
		}
		if len(value) == 5 && strings.Contains(value, oneParts[0]) && strings.Contains(value, oneParts[1]) {
			tracker[3] = value
			seen[value] = 3
		}
	}
	for _, value := range strings.Fields(signal) {
		fmt.Println("Mapping is", value, seen[value])
		if _, ok := seen[value]; !ok {
			fmt.Println("We have not come across this sequence yet, must be 0")
			tracker[0] = value
			seen[value] = 0
		}
	}
	seenSorted := map[string]int{}
	for key, v := range seen {
		s := strings.Split(key, "")
		sort.Strings(s)
		seenSorted[strings.Join(s, "")] = v
	}
	fmt.Println("")
	return seenSorted
}

func sortedString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func deleteChars(oldstring, chars string) string {
	for _, char := range strings.Split(chars, "") {
		oldstring = strings.Replace(oldstring, char, "", -1)
	}
	return oldstring
}
