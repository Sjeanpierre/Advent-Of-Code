package main

import (
	"fmt"
	"strings"

	"github.com/sjeanpierre/AOC2020/helpers"
)

type game struct {
	boards         []*board
	drawingNumbers []string
}

type board struct {
	locator   map[string]coor
	markedCol [5]int
	markedRow [5]int
	won       bool
}
type coor struct {
	row, col int
	marked   bool
}

func main() {
	lines := helpers.LoadFileLines("./input.txt")
	process(lines)
}

func process(lines []string) {
	var boards []*board
	fmt.Println("")
	drawingNumbers := strings.Split(lines[0], ",")
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			board := lines[i+1 : i+6]
			boards = append(boards, parseBoard(board))
		}
		i += 5
	}
	g := game{
		boards:         boards,
		drawingNumbers: drawingNumbers,
	}
	g.play()
}

func (g game) play() {
	for _, number := range g.drawingNumbers {
		for _, board := range g.boards {
			if board.won {
				continue
			}
			board.turn(number)
			if board.bingo() {
				board.won = true
				fmt.Println("Final Score", board.calculateUnmarked()*helpers.StringToInt(number))
			}
		}
	}
}

func (b *board) calculateUnmarked() int {
	sumUnmarked := 0
	for num, coor := range b.locator {
		if !coor.marked {
			sumUnmarked += helpers.StringToInt(num)
		}
	}
	return sumUnmarked
}

func (b *board) turn(drawn string) {
	pos, ok := b.locator[drawn]
	if ok {
		b.markedCol[pos.col]++
		b.markedRow[pos.row]++
		b.locator[drawn] = coor{
			row:    pos.row,
			col:    pos.col,
			marked: true,
		}
	}
}

func (b *board) bingo() bool {
	for _, marked := range b.markedRow {
		if marked > 4 {
			return true
		}
	}
	for _, marked := range b.markedCol {
		if marked > 4 {
			return true
		}
	}
	return false
}

func parseBoard(b []string) *board {
	newB := board{locator: map[string]coor{}}
	for rowIDX, row := range b {
		for colIDX, char := range strings.Fields(row) {
			newB.locator[string(char)] = coor{
				row:    rowIDX,
				col:    colIDX,
				marked: false,
			}
		}
	}
	return &newB
}
