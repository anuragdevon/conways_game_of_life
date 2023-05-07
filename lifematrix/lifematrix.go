package lifematrix

import (
	"errors"
	"math/rand"
)

type Grid struct {
	rows  int
	cols  int
	Cells [][]int
}

func NewGrid(rows, cols int) (*Grid, error) {
	if rows <= 0 && cols <= 0 {
		return &Grid{}, errors.New("invalid rows or cols values")
	}
	newLifeMatix := &Grid{
		rows:  rows,
		cols:  cols,
		Cells: make([][]int, rows),
	}
	for row := range newLifeMatix.Cells {
		newLifeMatix.Cells[row] = make([]int, cols)
	}
	return newLifeMatix, nil
}

func (grid *Grid) Randomize() {
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.cols; j++ {
			grid.Cells[i][j] = rand.Intn(2)
		}
	}
}

func (grid *Grid) SetCellAlive(x, y, alive int) {
	if x < grid.rows && y < grid.cols && x >= 0 && y >= 0 {
		grid.Cells[x][y] = alive
	}
}
