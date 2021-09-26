package main

import (
	"bytes"
	"fmt"
)

const aliveChar = '█'
const deadChar = '.'

type Playground struct {
	cols, lines int
	cells       [][]bool
}

func NewPlayground(cols, lines int) Playground {
	cells := make([][]bool, lines)
	f := Playground{
		cols:  cols,
		lines: lines,
		cells: cells,
	}

	for cl, _ := range cells {
		cells[cl] = make([]bool, cols)
	}
	return f
}

func (p Playground) String() string {
	var buf bytes.Buffer
	for _, line := range p.cells {
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

func main() {
	pg, err := NewPlaygroundFromFile("centered_glider.txt", 3, 3)
	if err != nil {
		panic(err)
	}
	fmt.Print(pg.String())

}
