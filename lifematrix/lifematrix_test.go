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
	count_number_of_ones, count_number_of_zeros := 0, 0
	test_grid.Randomize()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if test_grid.Cells[i][j] == 0 {
				count_number_of_zeros += 1
			} else if test_grid.Cells[i][j] == 1 {
				count_number_of_ones += 1
			} else {
				t.Error("expected 0 or 1, but got ", test_grid.Cells[i][j])
			}
		}
	}
	if count_number_of_ones == 0 || count_number_of_zeros == 0 {
		t.Errorf("Expected a random allocation of dead and live cells")
	}
}

func TestAliveCellForSetCellAliveToTrue(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, err := NewGrid(rows, cols)
	assert.Nil(t, err)
	test_grid.Randomize()

	test_grid.setCellAlive(1, 1)

	assert.Equal(t, 1, test_grid.Cells[1][1])
}
