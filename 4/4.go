package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

var requiredFields = regexp.MustCompile(`ecl:|pid:|eyr:|hcl:|byr:|iyr:|hgt:`)
var requiredFieldsValidated = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)|pid:(\d{9})|eyr:(202[0-9]|2030)|hcl:#([0-9a-f]{6})|byr:(19[2-9][0-9]|200[0-2])|iyr:(201[0-9]|2020)|hgt:((1[5-8][0-9]|19[0-3])cm|(59|[6-7][0-9]in))`)

func main() {
	var counter int
	lines := helpers.LoadFileLines("./input.txt")
	var localLines []string
	for _, line := range lines {
		if line == "" {
			if validPassport(strings.Join(localLines, " "), true) {
				counter++
			}
			localLines = []string{}
			continue
		}
		localLines = append(localLines, line)
	}
	fmt.Println(counter)
}

func validPassport(passportData string, extendedValidation bool) bool {
	if extendedValidation {
		return 7 == (strings.Count(requiredFieldsValidated.ReplaceAllString(passportData, "+"), "+"))
	}
	return 7 == (strings.Count(requiredFields.ReplaceAllString(passportData, "+"), "+"))
}
