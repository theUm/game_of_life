package main

import (
	"bufio"
	"fmt"
	"os"
)

func NewPlaygroundFromFile(fileName string, cols, lines int) (*Playground, error) {
	fd, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("open file for read only: %w", err)
	}

	cells := make([][]bool, 0, lines)
	pg := &Playground{
		cols:  cols,
		lines: lines,
		cells: cells,
	}

	data := bufio.NewScanner(fd)
	data.Split(bufio.ScanLines) // split file by lines
	for data.Scan() {
		// todo: make check for line length here
		currLine := make([]bool, 0, cols)
		for _, r := range data.Text() { // read runes one by one
			switch r {
			case aliveChar:
				currLine = append(currLine, true)
			case deadChar:
				currLine = append(currLine, false)
			default:
				return nil, fmt.Errorf("illegal character found:  \n%c\nonly %c,%c,\\n is accepted", r, aliveChar, deadChar)
			}
		}
		pg.cells = append(pg.cells, currLine)
	}

	return pg, nil
}
