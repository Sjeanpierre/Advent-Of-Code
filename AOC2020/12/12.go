package main

import (
	"fmt"

	"github.com/sjeanpierre/AOC2020/helpers"
)


func main() {

	lines := helpers.LoadFileLines("./input.txt")
	process(lines)

}

func process(lines []string) {
	pX := 0
	pY := 0

	heading := 0

	headings := []byte{'E', 'N', 'W', 'S'} //ccw
	for _, instruction := range lines {
		amount := helpers.StringToInt(instruction[1:len(instruction)])


		dir := instruction[0]


		switch instruction[0] {
		case 'R':
			heading -= amount / 90
			heading += 4
			heading %= 4
		case 'L':
			heading += amount / 90
			heading %= 4
		case 'F':
			dir = headings[heading]
		}

		switch dir {
		case 'N':
			pY += amount
		case 'S':
			pY -= amount
		case 'E':
			pX += amount
		case 'W':
			pX -= amount
		}
	}

	drMan := 0
	if pX >= 0 {
		drMan += pX
	} else {
		drMan -= pX
	}
	if pY >= 0 {
		drMan += pY
	} else {
		drMan -= pY
	}
	fmt.Println(drMan)

	// Part B.
	wayX := 10
	wayY := 1
	pX = 0
	pY = 0

	r := [][4]int{
		// 0
		{1, 0, 0, 1},
		// 90
		{0, -1, 1, 0},
		// 180
		{-1, 0, 0, -1},
		// 270
		{0, 1, -1, 0},
	}

	for _, s := range lines {
		param := helpers.StringToInt(s[1:len(s)])


		switch s[0] {
		case 'R':
			idx := (4 + (-param / 90)) % 4
			newWayX := r[idx][0]*wayX + r[idx][1]*wayY
			newWayY := r[idx][2]*wayX + r[idx][3]*wayY
			wayX = newWayX
			wayY = newWayY
		case 'L':
			idx := (param / 90) % 4
			newWayX := r[idx][0]*wayX + r[idx][1]*wayY
			newWayY := r[idx][2]*wayX + r[idx][3]*wayY
			wayX = newWayX
			wayY = newWayY
		case 'F':
			pX += wayX * param
			pY += wayY * param
		case 'N':
			wayY += param
		case 'S':
			wayY -= param
		case 'E':
			wayX += param
		case 'W':
			wayX -= param
		}
	}

	drMan = 0
	if pX >= 0 {
		drMan += pX
	} else {
		drMan -= pX
	}
	if pY >= 0 {
		drMan += pY
	} else {
		drMan -= pY
	}
	fmt.Println(drMan)
}
