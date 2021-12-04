package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	process1(lines)
	process2(lines)
}

func calculateFreq(lines []string) map[int]int {
	size := len(lines[0])
	counter := make([]int, int(size))

	for _, numbers := range lines {
		for position, digit := range strings.Split(numbers, "") {
			d := helpers.StringToInt(digit)
			if d != 0 {
				counter[position]++
			}
		}
	}

	var majorityMap = make(map[int]int)

	for position, freq := range counter {
		if float32(freq) > float32(len(lines))/2 {
			majorityMap[position] = 1
			continue
		}
		if float32(freq) < float32(len(lines))/2 {
			majorityMap[position] = 0
			continue
		} else {
			majorityMap[position] = 3
		}
	}

	return majorityMap
}

func filter(lines []string, pos, seeking int) []string {
	var result []string
	for _, line := range lines {
		if helpers.StringToInt(string(line[pos])) == seeking {
			result = append(result, line)
		}
	}
	return result
}

func process1(lines []string) {
	var delta []string
	var epsilon []string
	freqs := calculateFreq(lines)
	fmt.Println(freqs)
	for i := 0; i < len(freqs); i++ {
		fmt.Print(freqs[i])

		if freqs[i] == 1 {
			delta = append(delta, "1")
			epsilon = append(epsilon, "0")
		} else {
			delta = append(delta, "0")
			epsilon = append(epsilon, "1")
		}
	}
	deltaN, _ := strconv.ParseInt(strings.Join(delta, ""), 2, 32)
	epsilonN, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 32)
	fmt.Println("\n", deltaN*epsilonN)

}

func process2(lines []string) {
	filteredDelta := lines
	filteredEpsilon := lines
	z := 0
	for {
		z++
		fmt.Println(z)
		for i := 0; i < len(lines[0]); i++ {
			freqs := calculateFreq(filteredDelta)
			fmt.Println(freqs[i])
			switch freqs[i] {
			case 1:
				filteredDelta = filter(filteredDelta, i, 1)
			case 0:
				filteredDelta = filter(filteredDelta, i, 0)
			default:
				filteredDelta = filter(filteredDelta, i, 1)
			}
		}
		if len(filteredDelta) == 1 {
			break
		}
	}

	for {
		z++
		fmt.Println(z)
		for i := 0; i < len(lines[0]); i++ {
			freqs := calculateFreq(filteredEpsilon)
			fmt.Println(freqs[i])
			switch freqs[i] {
			case 1:
				filteredEpsilon = filter(filteredEpsilon, i, 0)
			case 0:
				filteredEpsilon = filter(filteredEpsilon, i, 1)
			default:
				filteredEpsilon = filter(filteredEpsilon, i, 0)
			}
			if len(filteredEpsilon) == 1 {
				break
			}
		}
		if len(filteredEpsilon) == 1 {
			break
		}
	}

	deltaN, _ := strconv.ParseInt(filteredDelta[0], 2, 32)
	epsilonN, _ := strconv.ParseInt(filteredEpsilon[0], 2, 32)
	fmt.Println("\n", deltaN*epsilonN)
}
