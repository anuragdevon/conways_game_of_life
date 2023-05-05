package mocklife

import (
	"conwaysgameoflife/lifematrix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockLife(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, err := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err)

	newMockLife := CreateMockLife(testGrid)
	if newMockLife.grid != testGrid {
		t.Errorf("expected both grids to be same")
	}

}

func TestCountLiveNeighborsInMiddle(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, err := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err)
}
