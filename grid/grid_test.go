package grid

import (
	"conwaysgameoflife/cell"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewGridCreationWithPositiveDimension(t *testing.T) {
	rows := 10
	cols := 10
	test_grid, err := NewGrid(rows, cols)
	assert.Nil(t, err)
	assert.Equal(t, rows, test_grid.rows)
	assert.Equal(t, cols, test_grid.cols)
}

func TestInvalidGridCreationWithZeroDimension(t *testing.T) {
	rows := 0
	cols := 0
	_, err := NewGrid(rows, cols)
	assert.NotNil(t, err)
}

func TestInvalidGridCreationWithNegativeDimension(t *testing.T) {
	rows := -1
	cols := 0
	_, err := NewGrid(rows, cols)
	assert.NotNil(t, err)
}

func TestValidRandomCellValuesInitialization(t *testing.T) {
	rows := 10
	cols := 10
	test_grid, _ := NewGrid(rows, cols)

	count_number_of_ones, count_number_of_zeros := 0, 0
	test_grid.Randomize()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if test_grid.Cells[i][j].IsAlive() {
				count_number_of_ones += 1
			} else {
				count_number_of_zeros += 1
			}
		}
	}
	assert.NotEqual(t, 0, count_number_of_ones)
	assert.NotEqual(t, 0, count_number_of_zeros)
}

func TestValidSetCellNeighbors(t *testing.T) {
	rows := 2
	cols := 2
	testGrid, _ := NewGrid(rows, cols)

	testGrid.Cells[0][0] = cell.NewAliveCell()
	testGrid.Cells[1][0] = cell.NewAliveCell()
	testGrid.Cells[1][1] = cell.NewAliveCell()

	testGrid.SetCellNeighbors()
	assert.Equal(t, 2, testGrid.Cells[0][0].NumberOfLiveNeighbors())
	assert.Equal(t, 3, testGrid.Cells[0][1].NumberOfLiveNeighbors())
	assert.Equal(t, 2, testGrid.Cells[1][0].NumberOfLiveNeighbors())
	assert.Equal(t, 2, testGrid.Cells[1][1].NumberOfLiveNeighbors())
}
