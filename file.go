package main

import (
	"bufio"
	"fmt"
	"os"
)

func NewPlaygroundFromFile(fileName string, pg *Playground) (*Playground, error) {
	fd, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("open file for read only: %w", err)
	}

	data := bufio.NewScanner(fd)
	data.Split(bufio.ScanLines) // split file by lines
	lineNum := 0                // lines counter
	for data.Scan() {
		// check whether all lines have provided length
		if len([]rune(data.Text())) != pg.cols {
			return nil, fmt.Errorf("line length is not same as provided length. the line:\n%s\nexpected to be %d chars length", data.Text(), pg.cols)
		}
		currLine := make([]bool, 0, pg.cols)
		for _, r := range data.Text() { // read runes one by one
			switch r {
			case aliveChar:
				currLine = append(currLine, true)
			case deadChar:
				currLine = append(currLine, false)
			default:
				return nil, fmt.Errorf("illegal character found:  \n%c\nonly %c,%c,\\n is accepted", r, aliveChar, deadChar)
			}
			pg.cells[lineNum] = currLine
		}
		lineNum++
	}

	return pg, nil
}
