package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var PARTB bool

type rule struct {
	fieldName                                    string
	rangeALow, rangeAHigh, rangeBLow, rangeBHigh int
}

func main() {
	PARTB = true
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {

	//Process Rules
	var ruleSet []rule
	for _, line := range lines {
		if line == "" {
			break
		}
		lineParts := strings.Split(line, ":")
		ruleThresholdParts := strings.Split(lineParts[1], "or")
		ruleA := strings.Split(strings.TrimSpace(ruleThresholdParts[0]), "-")
		ruleB := strings.Split(strings.TrimSpace(ruleThresholdParts[1]), "-")
		ruleSet = append(
			ruleSet, rule{
				fieldName:  strings.TrimSpace(lineParts[0]),
				rangeALow:  helpers.StringToInt(ruleA[0]),
				rangeAHigh: helpers.StringToInt(ruleA[1]),
				rangeBLow:  helpers.StringToInt(ruleB[0]),
				rangeBHigh: helpers.StringToInt(ruleB[1]),
			},
		)

	}

	//Process ticket data
	var inTickets bool
	var ticketData = make(map[int][]int)
	sumBad :=0
	for _, line := range lines {
		if inTickets {
			if !allValid(ruleSet,line) {
				continue
			}
			for column, value := range strings.Split(line, ",") {
				if !PARTB && !anyValid(ruleSet,helpers.StringToInt(value)) {
					sumBad +=helpers.StringToInt(value)
					continue
				}
				ticketData[column] = append(ticketData[column], helpers.StringToInt(value))
			}
		}
		if line == "nearby tickets:" {
			inTickets = true
		}
	}
	ruleMapping := calculate(ticketData, ruleSet)
	fmt.Println(ruleMapping)
	fmt.Println(sumBad)

}



func calculate(ticketData map[int][]int, rules []rule) map[int]rule {
	var mapping = make(map[int]rule)
	for column, data := range ticketData {
		sort.Ints(data)
		highScore := float64(0)
		for _, r := range rules {
			validPercentage := r.validate(data)
			if validPercentage > highScore {
				mapping[column] = r
				highScore = validPercentage
			}
			fmt.Println("Column", column, validPercentage, "%", "valid for field", r.fieldName)
		}
	}
	return mapping
}

func (r rule) validate(d []int) float64 {
	countValid := 0
	for _, value := range d {
		if value >= r.rangeALow && value <= r.rangeAHigh || value >= r.rangeBLow && value <= r.rangeBHigh {
			countValid++
		}
	}
	if countValid == len(d) {
		fmt.Println("")
	}
	return float64(countValid) / float64(len(d))
}

func allValid(rules []rule, ticket string) bool {
	ticketData := strings.Split(ticket,",")
	for _, v := range ticketData {
		if !anyValid(rules,helpers.StringToInt(v)) {
			return false
		}
	}
	return true
}

func anyValid(rules []rule, value int) bool {
	for _, r := range rules {
		if value >= r.rangeALow && value <= r.rangeAHigh || value >= r.rangeBLow && value <= r.rangeBHigh {
			return true
		}
	}
	return false
}