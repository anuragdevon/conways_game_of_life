package grid

import (
	"conwaysgameoflife/cell"
	"errors"
	"math/rand"
)

type Grid struct {
	rows  int
	cols  int
	Cells [][]*cell.Cell
}

func NewGrid(rows, cols int) (*Grid, error) {
	if rows <= 0 && cols <= 0 {
		return &Grid{}, errors.New("invalid rows or cols values")
	}
	newLifeMatix := &Grid{
		rows:  rows,
		cols:  cols,
		Cells: make([][]*cell.Cell, rows),
	}
	for row := range newLifeMatix.Cells {
		newLifeMatix.Cells[row] = make([]*cell.Cell, cols)
		for col := range newLifeMatix.Cells[row] {
			newLifeMatix.Cells[row][col] = cell.NewDeadCell()
		}
	}
	return newLifeMatix, nil
}

func (grid *Grid) Randomize() {
	for i := 0; i < grid.rows; i++ {
		for j := 0; j < grid.cols; j++ {
			if rand.Intn(2) == 0 {
				grid.Cells[i][j] = cell.NewDeadCell()
			} else {
				grid.Cells[i][j] = cell.NewAliveCell()
			}
		}
	}
}

func (grid *Grid) NumberOfLiveNeighbors(row, col int) int {
	countLiveNeighbors := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < grid.rows && j >= 0 && j < grid.cols && !(i == row && j == col) {
				if grid.Cells[i][j].IsAlive() {
					countLiveNeighbors += 1
				}
			}
		}
	}
	return countLiveNeighbors
}
