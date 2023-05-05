package lifematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewGridCreationWithPositiveDimension(t *testing.T) {
	rows := 10
	cols := 10
	test_grid, err := NewGrid(rows, cols)
	assert.Nil(t, err)
	if test_grid.rows != rows || test_grid.cols != cols {
		t.Errorf("expected a grid of %d rows and %d cols", rows, cols)
	}
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
	test_grid, err := NewGrid(rows, cols)
	assert.Nil(t, err)

	test_grid.Randomize()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if test_grid.cells[i][j] != 0 && test_grid.cells[i][j] != 1 {
				t.Error("expected 0 or 1, but got %d", test_grid.cells[i][j])
			}
		}
	}
}
