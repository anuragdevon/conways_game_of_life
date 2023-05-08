package grid

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
	test_grid, _ := NewGrid(rows, cols)

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

func TestAliveCellForSetCellAliveToOne(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := NewGrid(rows, cols)

	test_grid.Randomize()

	test_grid.SetCellAlive(1, 1, 1)

	assert.Equal(t, 1, test_grid.Cells[1][1])
}

func TestDeadCellForSetCellAliveToZero(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := NewGrid(rows, cols)

	test_grid.Randomize()

	test_grid.SetCellAlive(1, 1, 0)

	assert.Equal(t, 0, test_grid.Cells[1][1])
}

func TestNoChangeInCellValueForSetAliveCellForInvalidCell(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := NewGrid(rows, cols)

	test_grid.Randomize()

	earlier_value := test_grid.Cells[1][1]
	test_grid.SetCellAlive(3, 2, 1)

	assert.Equal(t, earlier_value, test_grid.Cells[1][1])
}

func TestNoChangeInCellValueForInvalidAliveValue(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := NewGrid(rows, cols)

	test_grid.Randomize()

	earlier_value := test_grid.Cells[1][1]
	test_grid.SetCellAlive(1, 1, 3)

	assert.Equal(t, earlier_value, test_grid.Cells[1][1])
}

func TestValidNumberOfLiveNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := NewGrid(rows, cols)

	test_grid.SetCellAlive(0, 0, 1)
	test_grid.SetCellAlive(0, 1, 1)
	test_grid.SetCellAlive(0, 2, 1)

	test_grid.SetCellAlive(1, 0, 0)
	test_grid.SetCellAlive(1, 1, 1)
	test_grid.SetCellAlive(1, 2, 0)

	test_grid.SetCellAlive(2, 0, 1)
	test_grid.SetCellAlive(2, 1, 1)
	test_grid.SetCellAlive(2, 2, 1)

	expected_live_neighbors_for_cell_1_1 := 6
	actual_live_neighbors_for_cell_1_1 := test_grid.NumberOfLiveNeighbors(1, 1)

	if actual_live_neighbors_for_cell_1_1 != expected_live_neighbors_for_cell_1_1 {
		t.Errorf("Expected %d neighbors but got %d", expected_live_neighbors_for_cell_1_1, actual_live_neighbors_for_cell_1_1)
	}
}
