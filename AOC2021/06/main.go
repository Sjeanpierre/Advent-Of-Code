package main

import (
	"fmt"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

type fish struct {
	members map[int]int
}

func main() {
	schoolOfFish := fish{members: map[int]int{}}
	lines := helpers.LoadFileLines("./input.txt")
	for _, f := range strings.Split(lines[0], ",") {
		schoolOfFish.members[helpers.StringToInt(f)]++
	}
	s := schoolOfFish.process(256)
	var total int
	for _, count := range s.members {
		total += count
	}
	fmt.Println(total)
}

func (f fish) process(days int) fish {
	if days == 0 {
		return f
	}
	nf := fish{members: map[int]int{}}
	newChildren := f.members[0]
	nf.members[7] = f.members[8]
	nf.members[6] = f.members[7]
	nf.members[5] = f.members[6]
	nf.members[4] = f.members[5]
	nf.members[3] = f.members[4]
	nf.members[2] = f.members[3]
	nf.members[1] = f.members[2]
	nf.members[0] = f.members[1]
	nf.members[8] += newChildren
	nf.members[6] += newChildren
	return nf.process(days - 1)
}
