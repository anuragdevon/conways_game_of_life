package mocklife

import (
	"conwaysgameoflife/lifematrix"
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

func TestCalculateValidLiveNeighbors(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, err1 := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err1)

	testGrid.Cells[0][0] = 1
	testGrid.Cells[0][1] = 1
	testGrid.Cells[0][2] = 0
	testGrid.Cells[1][0] = 0
	testGrid.Cells[1][1] = 0
	testGrid.Cells[1][2] = 1
	testGrid.Cells[2][0] = 1
	testGrid.Cells[2][1] = 0
	testGrid.Cells[2][2] = 0

	liveNeighbors, err2 := CalculateLiveNeighbors(1, 1)
	assert.Nil(t, err2)

	if liveNeighbors != 2 {
		t.Error("expected 2 live neighbors, got %d", liveNeighbors)
	}
}
