package simulation

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

func TestCreateValidNextStateOfSimulation(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, err := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err)

	testGrid.SetCellAlive(0, 0, 0)
	testGrid.SetCellAlive(0, 1, 1)
	testGrid.SetCellAlive(0, 2, 0)

	testGrid.SetCellAlive(1, 0, 0)
	testGrid.SetCellAlive(1, 1, 0)
	testGrid.SetCellAlive(1, 2, 1)

	testGrid.SetCellAlive(2, 0, 1)
	testGrid.SetCellAlive(2, 1, 1)
	testGrid.SetCellAlive(2, 2, 1)

	testNewSimulation := NewSimulation(testGrid)

	expectedNextStateGrid := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
	}

	testNewSimulation.nextStateOfSimulation()
	if !reflect.DeepEqual(testNewSimulation.grid.Cells, expectedNextStateGrid) {
		t.Errorf("expected state did not match with evaluated state")
	}
}

func TestLiveCellToDieOfUnderpopulation(t *testing.T) {
	testGrid, _ := lifematrix.NewGrid(3, 3)

	testGrid.SetCellAlive(0, 0, 1)
	testGrid.SetCellAlive(0, 1, 0)
	testGrid.SetCellAlive(0, 2, 0)
	testGrid.SetCellAlive(1, 0, 0)
	testGrid.SetCellAlive(1, 1, 1)
	testGrid.SetCellAlive(1, 2, 0)
	testGrid.SetCellAlive(2, 0, 0)
	testGrid.SetCellAlive(2, 1, 0)
	testGrid.SetCellAlive(2, 2, 0)

	expectedNextStateGrid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	testSimulation := NewSimulation(testGrid)
	testSimulation.nextStateOfSimulation()
	assert.Equal(t, expectedNextStateGrid, testSimulation.grid.Cells)
}
