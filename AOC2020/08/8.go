package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

func main() {

	lines := helpers.LoadFileLines("./input.txt")

	nops, jumps := process(lines,-1,-1)

	for _, i := range nops {
		fmt.Println("running nop",i)
		_,_ = process(lines,i,-1)
	}

	for _, i := range jumps {
		fmt.Println("running jump",i)
		_,_ = process(lines,-1,i)
	}

}

func process(lines []string,nopSkip, jumpSkip int) (a, b []int){
	accumulator := 0
	x := 0
	var seen = make(map[int]bool)
	var nops []int
	var jumps []int
	for  {
		if seen[x] {
		  break
		}
		seen[x]=true
		parts := strings.Split(lines[x]," ")
		instruction, dist := parts[0], helpers.StringToInt(parts[1])
		switch instruction {
		case "nop":
			if x == nopSkip {
				x+=dist
				continue
			}
			nops = append(nops,x)
			x++
		case "acc":
			accumulator += dist
			x++
		case "jmp":
			if x == jumpSkip {
				x++
				continue
			}
			jumps = append(jumps,x)
			x+=dist
		}

		if x >= len(lines) {
			log.Fatalln("BINGO",accumulator)
		}
	}

	fmt.Println(accumulator)
	return nops,jumps
}

