package helpers


import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func LoadFileLines(filepath string) []string {
	var lines []string

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Could not open input file", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func StringToInt(s string) int {
	i,err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("welp you should have known better",err,s)
	}
	return i
}