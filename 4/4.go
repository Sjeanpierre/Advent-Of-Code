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
			if validPassport(strings.Join(llines," "), extendedValidation) {
				counter +=1
				//fmt.Println(strings.Join(llines," "))
			}
			llines = []string{}
			continue
		}
		llines = append(llines,linedata)
	}
	fmt.Println(counter)

}

func validPassport(passportData string, extendedValidation bool)  bool {
	if extendedValidation {
		temp := requiredFieldsValidated.ReplaceAllString(passportData,"+")
		if 7 == (strings.Count(temp,"+")) {
			parts := strings.Split(passportData," ")
			for _, pair := range parts {
				splitParts := strings.Split(pair,":")
				key := splitParts[0]
				value := splitParts[1]
				if key == "eyr" {
					i := parseInt(value)
					if i < 2020 || i > 2030 {
						log.Println("eyr issue with",passportData)
					}
				}

				if key == "byr" {
					i := parseInt(value)
					if i < 1920 || i > 2002 {
						log.Println("byr issue with",passportData)
					}
				}

				if key == "iyr" {
					i := parseInt(value)
					if i < 2010 || i > 2020 {
						log.Println("iyr issue with",passportData)
					}
				}

				if key == "hgt" {
					if strings.Contains(value,"in") {
						h := strings.Replace(value,"in","",-1)
						i := parseInt(h)
						if i < 59 || i > 76 {
							log.Println("hgt in issue with",passportData)
						}
					}
				}

				if key == "hgt" {
					if strings.Contains(value,"cm") {
						h := strings.Replace(value,"cm","",-1)
						i := parseInt(h)
						if i < 150 || i > 193 {
							log.Println("hgt cm issue with",passportData)
						}
					}
				}
			}
			return true
		}
		return false
	}

	return 7 == (strings.Count(requiredFields.ReplaceAllString(passportData,"+"),"+"))
}

func parseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}