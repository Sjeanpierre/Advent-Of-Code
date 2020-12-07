package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var CCC []string

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
	part2(lines)
}

//Finds out which bags can contain top level bag, inverse
func part1(lines []string) {
	var ruleSet = make(map[string][]string)
	for _, rule := range lines {
		ruleParts := strings.Split(rule, "contain")
		bagDescription, contents := strings.TrimSpace(strings.Replace(ruleParts[0], "bags", "bag", -1)), ruleParts[1]
		subBags := strings.Split(strings.Replace(contents, ".", "", -1), ",")
		for _, bag := range subBags {
			num := regexp.MustCompile(`\d+`)
			bagDes := num.ReplaceAllString(bag, "")
			bagDes = strings.TrimSpace(strings.Replace(bagDes, "bags", "bag", -1))
			ruleSet[bagDes] = append(ruleSet[bagDes], bagDescription)
		}
	}
	fmt.Println("")
	var canHold = make(map[string]bool)
	listCanHold(ruleSet, "shiny gold bag")

	for x := 0; x < 10; x++ {
		for _, bag := range CCC {
			listCanHold(ruleSet, bag)
		}

	}
	//dedup?
	for _, v := range CCC {
		canHold[v] = true
	}
	fmt.Println(len(canHold))

}

func part2(lines []string) {
	var ruleSet = make(map[string]map[string]int)
	var ruleSet2 = make(map[string][]string)
	for _, rule := range lines {
		ruleParts := strings.Split(rule, "contain")
		bagDescription, contents := strings.TrimSpace(strings.Replace(ruleParts[0], "bags", "bag", -1)), ruleParts[1]
		subBags := strings.Split(strings.Replace(contents, ".", "", -1), ",")
		var bagCounts = make(map[string]int)
		for _, bag := range subBags {
			bagDes := strings.TrimSpace(strings.Replace(bag, "bags", "bag", -1))
			ruleSet2[bagDes] = append(ruleSet2[bagDes], bagDescription)
			count, _ := strconv.Atoi(strings.TrimSpace(bagDes[:2]))
			bagCounts[bagDes[2:]] = count
		}
		ruleSet[bagDescription] = bagCounts
	}
	var mapping = make(map[int]map[string]int)
	mapping[0] = ruleSet["shiny gold bag"]

	calculateBagTree(mapping, 0, ruleSet)

	var total int
	for _, b := range mapping {
		for _, count := range b {
			total += count
		}
	}
	fmt.Println(total)
}

func calculateBagTree(mapping map[int]map[string]int, level int, ruleSet map[string]map[string]int) {
	holder := map[string]int{}
	for key, value := range mapping[level] {
		for k1, count := range getBagsAtLevel(ruleSet, key, value) {
			holder[k1] = holder[k1] + count
		}
	}
	if len(holder) == 0 {
		return
	}
	nextLevel := level + 1
	mapping[nextLevel] = holder
	calculateBagTree(mapping, nextLevel, ruleSet)
}

func getBagsAtLevel(ruleSet map[string]map[string]int, key string, value int) map[string]int {
	var tmp = map[string]int{}
	subLevel := ruleSet[key]
	for k2, v2 := range subLevel {
		if k2 == " other bag" {
			continue
		}
		for y := value; y > 0; y-- {
			tmp[k2] += v2
		}
	}
	return tmp
}

func listCanHold(s map[string][]string, description string) {
	CCC = append(CCC, s[description]...)
}
