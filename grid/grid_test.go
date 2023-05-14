package grid

import (
	"conwaysgameoflife/cell"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGridFormWithPositiveDimension(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, err := NewGrid(rows, cols)
	assert.Nil(t, err)
	assert.Equal(t, rows, testGrid.rows)
	assert.Equal(t, cols, testGrid.cols)
}

func TestNewGridNotFormWithZeroDimension(t *testing.T) {
	rows := 0
	cols := 0
	_, err := NewGrid(rows, cols)
	assert.NotNil(t, err)
}

func TestNewGridNotFormWithNegativeDimension(t *testing.T) {
	rows := -1
	cols := 0
	_, err := NewGrid(rows, cols)
	assert.NotNil(t, err)
}

func TestRandomizeAllocatesRandomCellValues(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, _ := NewGrid(rows, cols)

	count_number_of_ones, count_number_of_zeros := 0, 0
	testGrid.Randomize()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if testGrid.Cells[i][j].IsAlive() {
				count_number_of_ones += 1
			} else {
				count_number_of_zeros += 1
			}
		}
	}
	assert.NotEqual(t, 0, count_number_of_ones)
	assert.NotEqual(t, 0, count_number_of_zeros)
}

func TestAddCellNeighborsAllocatesNeighborsToCell(t *testing.T) {
	rows := 2
	cols := 2
	testGrid, _ := NewGrid(rows, cols)

	testGrid.Cells[0][0] = cell.NewAliveCell()
	testGrid.Cells[1][0] = cell.NewAliveCell()
	testGrid.Cells[1][1] = cell.NewAliveCell()

	testGrid.AddCellNeighbors()
	assert.Equal(t, 2, testGrid.Cells[0][0].NumberOfLiveNeighbors())
	assert.Equal(t, 3, testGrid.Cells[0][1].NumberOfLiveNeighbors())
	assert.Equal(t, 2, testGrid.Cells[1][0].NumberOfLiveNeighbors())
	assert.Equal(t, 2, testGrid.Cells[1][1].NumberOfLiveNeighbors())
}
