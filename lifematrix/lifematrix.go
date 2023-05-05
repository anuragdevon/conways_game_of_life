package lifematrix

import "errors"

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
