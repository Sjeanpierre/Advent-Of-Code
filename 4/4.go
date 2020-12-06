package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var requiredFields = regexp.MustCompile(`ecl:|pid:|eyr:|hcl:|byr:|iyr:|hgt:`)
var requiredFieldsValidated = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)|pid:(\d{9})|eyr:(202[0-9]|2030)|hcl:#([0-9a-f]{6})|byr:(19[2-9][0-9]|200[0-2])|iyr:(201[0-9]|2020)|hgt:((1[5-8][0-9]|19[0-3])cm|(59|[6-7][0-9]in))`)

func main() {
	extendedValidation := true
	//var lines []string
	var counter int

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Could not open input file", err)
	}
	scanner := bufio.NewScanner(file)
	var llines []string

	for scanner.Scan() {
		linedata := scanner.Text()
		if linedata == "" {
			if validPassport(strings.Join(llines, " "), extendedValidation) {
				counter += 1
			}
			llines = []string{}
			continue
		}
		llines = append(llines, linedata)
	}
	fmt.Println(counter)

}

func validPassport(passportData string, extendedValidation bool) bool {
	if extendedValidation {
		return 7 == (strings.Count(requiredFieldsValidated.ReplaceAllString(passportData, "+"), "+"))
	}
	return 7 == (strings.Count(requiredFields.ReplaceAllString(passportData, "+"), "+"))
}
