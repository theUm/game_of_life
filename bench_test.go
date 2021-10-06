package main

import (
	"testing"
)

func BenchmarkSample(b *testing.B) {
	basePG := NewPlayground(25, 25)
	pg, err := NewPlaygroundFromFile("centered_glider.txt", basePG)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		c := make([][]bool, len(pg.cells))
		fc := make([][]bool, len(pg.cells))
		copy(c, pg.cells)
		copy(fc, pg.futureCells)
		newPg := Playground{
			cols:        25,
			lines:       25,
			cells:       c,
			futureCells: fc,
		}

		for i := 0; i < 100; i++ {
			newPg.Tick()
		}
	}
}
