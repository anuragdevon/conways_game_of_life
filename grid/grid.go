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
	if rows <= 0 || cols <= 0 {
		return nil, errors.New("invalid rows or cols values")
	}
	newGrid := &Grid{
		rows:  rows,
		cols:  cols,
		Cells: make([][]*cell.Cell, rows),
	}
	for row := range newGrid.Cells {
		newGrid.Cells[row] = make([]*cell.Cell, cols)
		for col := range newGrid.Cells[row] {
			newGrid.Cells[row][col] = cell.NewDeadCell()
		}
	}
	return newGrid, nil
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

func (grid *Grid) AddCellNeighbors() {
	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			currentCell := grid.Cells[row][col]
			for neighborRow := row - 1; neighborRow <= row+1; neighborRow++ {
				for neighborCol := col - 1; neighborCol <= col+1; neighborCol++ {
					if neighborRow == row && neighborCol == col {
						continue
					}
					if neighborRow >= 0 && neighborRow < grid.rows && neighborCol >= 0 && neighborCol < grid.cols {
						currentCell.Neighbors = append(currentCell.Neighbors, grid.Cells[neighborRow][neighborCol])
					}
				}
			}
		}
	}
}
