package mocklife

import (
	"conwaysgameoflife/lifematrix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockLife(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, err1 := lifematrix.NewGrid(rows, cols)
	assert.Nil(t, err1)

	newMockLife, err2 := MockLife(testGrid)
	assert.Nil(t, err1)
	if newMockLife.grid != testGrid {
		t.Errorf("expected both grids to be same")
	}

}
