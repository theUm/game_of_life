package main

import (
	"bytes"
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
	//basePG := NewPlayground(5, 5)
	//pg, err := NewPlaygroundFromFile("3x3.txt", basePG)
	basePG := NewPlayground(16, 16)
	pg, err := NewPlaygroundFromFile("centered_glider.txt", basePG)
	if err != nil {
		panic(err)
	}
	fmt.Print(pg.String())

	// todo: put those tests onto test for Alive func
	//fmt.Println(pg.Alive(-6, 1))
	//fmt.Println(pg.Alive(-5, 1))
	//fmt.Println(pg.Alive(-4, 1))
	//
	//fmt.Println(pg.Alive(-3, 1))
	//fmt.Println(pg.Alive(-2, 1))
	//fmt.Println(pg.Alive(-1, 1))
	//
	//fmt.Println(pg.Alive(0, 1))
	//fmt.Println(pg.Alive(1, 1))
	//fmt.Println(pg.Alive(2, 1))
	//
	//fmt.Println(pg.Alive(3, 1))
	//fmt.Println(pg.Alive(4, 1))
	//fmt.Println(pg.Alive(5, 1))
	//fmt.Println(pg.Alive(5,1))

	//pg.evolute(true, 1, 1)
	//fmt.Println("")

	// glider should be on bottom by the 24th iteration on 16x16 filed. its damn hard to clear the cli on Windows :/
	for i := 0; i < 24; i++ {
		pg.Tick()
		fmt.Print(pg.String(), "\n")
	}
}
