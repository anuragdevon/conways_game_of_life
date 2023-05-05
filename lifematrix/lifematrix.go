package lifematrix

import (
	"errors"
	"math/rand"
)

type Grid struct {
	rows  int
	cols  int
	cells [][]int
}

func NewGrid(rows, cols int) (*Grid, error) {
	if rows <= 0 && cols <= 0 {
		return &Grid{}, errors.New("invalid rows or cols values")
	}
	newLifeMatix := &Grid{
		rows:  rows,
		cols:  cols,
		cells: make([][]int, rows),
	}
	for row := range newLifeMatix.cells {
		newLifeMatix.cells[row] = make([]int, cols)
	}
	return newLifeMatix, nil
}

func (g *Grid) Randomize() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			g.cells[i][j] = rand.Intn(2)
		}
	}
}
