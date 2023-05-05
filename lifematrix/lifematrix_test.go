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
