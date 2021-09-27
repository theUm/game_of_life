package main

import (
	"bytes"
	"flag"
	"fmt"
)

const aliveChar = '█'
const deadChar = '.'

type Playground struct {
	cols, lines int
	cells       [][]bool // current state of playground
	futureCells [][]bool // next tick of playground
}

func NewPlayground(cols, lines int) *Playground {
	cells := make([][]bool, lines)
	futureCells := make([][]bool, lines)

	// allocate the cells
	for cl := range cells {
		cells[cl] = make([]bool, cols)
		futureCells[cl] = make([]bool, cols)
	}

	return &Playground{
		cols:        cols,
		lines:       lines,
		cells:       cells,
		futureCells: futureCells,
	}
}

func (pg Playground) String() string {
	var buf bytes.Buffer
	for _, line := range pg.cells {
		for _, c := range line {
			char := deadChar
			if c {
				char = aliveChar
			}

			buf.WriteRune(char)
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

// Alive returns  whether cell is alive. Alive can handle out of bound indexes to wrap over the playground
func (pg *Playground) Alive(x, y int) bool {
	x %= pg.cols
	x += pg.cols
	x %= pg.cols

	y %= pg.lines
	y += pg.lines
	y %= pg.lines
	return pg.cells[y][x]
}

func (pg *Playground) Tick() {
	for y, line := range pg.cells {
		for x, cell := range line {
			pg.futureCells[y][x] = pg.evolute(cell, x, y)
		}
	}

	pg.cells, pg.futureCells = pg.futureCells, pg.cells // yeah like in python :)
}

func (pg *Playground) evolute(curVal bool, x, y int) bool {
	alive := 0
	for y1 := y - 1; y1 <= y+1; y1++ {
		for x1 := x - 1; x1 <= x+1; x1++ {
			if !((x1 == x) && (y1 == y)) && pg.Alive(x1, y1) { // count all alive except center one
				alive++
			}
		}
	}

	return alive == 3 || (alive == 2 && curVal)
}

func main() {
	var filePath string
	var width, height, gens int
	flag.StringVar(&filePath, "file", "centered_glider.txt", "path to txt file with initial playground state")
	flag.IntVar(&width, "w", 16, "width of playground in chars")
	flag.IntVar(&height, "h", 16, "height of playground in chars")
	flag.IntVar(&gens, "gens", 24, "ticks/generations to run")
	flag.Parse()

	// init playground
	basePG := NewPlayground(width, height)

	// populate by file
	pg, err := NewPlaygroundFromFile(filePath, basePG)
	if err != nil {
		panic(err)
	}
	fmt.Print(pg.String())

	// glider should be on bottom by the 24th iteration on 16x16 filed :/
	for i := 0; i < gens; i++ {
		pg.Tick()
		fmt.Print(pg.String(), "\n")
	}
}
