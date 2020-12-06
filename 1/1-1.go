package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	var lines []int

	file, err := os.Open("./1-1_input.txt")
	if err != nil {
		log.Fatalln("Could not open input file", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		lines = append(lines, int(i))
	}
	//part1(lines)
	part2(lines)

}

func part1(lines []int) {
	for _, num := range lines {
		for _, numb := range lines {
			if (num + numb) == 2020 {
				log.Println(num * numb)
				os.Exit(0)
			}

		}
	}
}

func part2(lines []int) {
	for _, num1 := range lines {
		for _, num2 := range lines {
			for _, num3 := range lines {
				if (num1 + num2 + num3) == 2020 {
					log.Println(num1 * num2 * num3)
					os.Exit(0)
				}
			}
		}
	}
}
