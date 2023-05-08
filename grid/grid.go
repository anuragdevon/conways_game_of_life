package grid

import (
	"errors"
	"math/rand"
)

// Associate behavior with data => keep a data behavior entity
// Avoid using common nouns as => as entities
// Mock and spy
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
	if x < grid.rows && y < grid.cols && x >= 0 && y >= 0 && (alive == 0 || alive == 1) {
		grid.Cells[x][y] = alive
	}
}

func (grid *Grid) NumberOfLiveNeighbors(row, col int) int {
	countLiveNeighbors := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < grid.rows && j >= 0 && j < grid.cols && !(i == row && j == col) {
				if grid.Cells[i][j] == 1 {
					countLiveNeighbors += 1
				}
			}
		}
	}
	return countLiveNeighbors
}
