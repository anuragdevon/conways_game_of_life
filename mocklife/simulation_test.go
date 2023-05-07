package mocklife

import (
	"conwaysgameoflife/lifematrix"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidNewSimulation(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, err := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err)

	testNewSimulation := NewSimulation(testGrid)
	if testNewSimulation.grid != testGrid {
		t.Errorf("expected both grids to be same")
	}
}

func TestValidNumberOfLiveNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, err := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err)

	testGrid.SetCellAlive(0, 0, 1)
	testGrid.SetCellAlive(0, 1, 1)
	testGrid.SetCellAlive(0, 2, 1)

	testGrid.SetCellAlive(1, 0, 0)
	testGrid.SetCellAlive(1, 1, 0)
	testGrid.SetCellAlive(1, 2, 0)

	testGrid.SetCellAlive(2, 0, 1)
	testGrid.SetCellAlive(2, 1, 1)
	testGrid.SetCellAlive(0, 0, 1)

	expected_live_neighbors_for_cell_1_1 := 6
	actual_live_neighbors_for_cell_1_1 := testGrid.NumberOfLiveNeighbors(1, 1)

	if actual_live_neighbors_for_cell_1_1 != expected_live_neighbors_for_cell_1_1 {
		t.Errorf("Expected %d neighbors but got %d", actual_live_neighbors_for_cell_1_1, expected_live_neighbors_for_cell_1_1)
	}

}

func TestNextStateOfSimulation(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, err := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err)

	testGrid.SetCellAlive(0, 0, 0)
	testGrid.SetCellAlive(0, 1, 1)
	testGrid.SetCellAlive(0, 2, 0)

	testGrid.SetCellAlive(1, 0, 1)
	testGrid.SetCellAlive(1, 1, 1)
	testGrid.SetCellAlive(1, 2, 1)

	testGrid.SetCellAlive(2, 0, 0)
	testGrid.SetCellAlive(2, 1, 1)
	testGrid.SetCellAlive(0, 0, 0)

	testNewSimulation := NewSimulation(testGrid)

	expectedNextStateGrid := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}

	testNewSimulation.NextStateOfSimulation()
	if !reflect.DeepEqual(testNewSimulation.grid.Cells, expectedNextStateGrid) {
		t.Errorf("expected state did not match with evaluated state")
	}
}
