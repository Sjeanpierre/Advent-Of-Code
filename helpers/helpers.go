package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coor struct {
	Row, Col int
}

type grid struct {
	cells map[Coor]int
}

func NewGrid() grid {
	return grid{cells: map[Coor]int{}}
}

func (g *grid) Incr(coor Coor) {
	g.cells[coor]++
}
func (g grid) Size() int {
	return len(g.cells)
}

func (g grid) Locate(coor Coor) (int, bool) {
	val, ok := g.cells[coor]
	return val, ok
}
func (g grid) GetCell(coor Coor) int {
	return g.cells[coor]
}

func (g *grid) SetCell(coor Coor, value int) {
	g.cells[coor] = value
}

func (g grid) PopulateWithLines(lines []string) {
	for rowIdx, line := range lines {
		for colIdx, value := range strings.Split(line, "") {
			g.cells[Coor{Row: rowIdx, Col: colIdx}] = StringToInt(value)
		}
	}
}

func (g grid) PrintGrid(rows, cols int) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			fmt.Print(g.cells[Coor{Row: row, Col: col}])

		}
		fmt.Println()
	}
}

func (c Coor) MapNeighbors() (Coor, Coor, Coor, Coor) {
	return Coor{Row: c.Row + 1, Col: c.Col}, //down
		Coor{Row: c.Row - 1, Col: c.Col}, //up
		Coor{Row: c.Row, Col: c.Col - 1}, //left
		Coor{Row: c.Row, Col: c.Col + 1} //right
}

func (c Coor) MapNeighborswDiag() (Coor, Coor, Coor, Coor, Coor, Coor, Coor, Coor) {
	return Coor{Row: c.Row + 1, Col: c.Col}, //down
		Coor{Row: c.Row - 1, Col: c.Col}, //up
		Coor{Row: c.Row, Col: c.Col - 1}, //left
		Coor{Row: c.Row, Col: c.Col + 1}, //right
		Coor{Row: c.Row + 1, Col: c.Col - 1}, //down-left
		Coor{Row: c.Row + 1, Col: c.Col + 1}, //down-right
		Coor{Row: c.Row - 1, Col: c.Col - 1}, //up-left
		Coor{Row: c.Row - 1, Col: c.Col + 1} //up-right
}

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
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("welp you should have known better", err, s)
	}
	return i
}
