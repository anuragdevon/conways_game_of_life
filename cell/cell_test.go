package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewAliveCellGeneration(t *testing.T) {
	new_cell := NewAliveCell()
	assert.Equal(t, Alive, new_cell.status)
}

func TestValidNewDeadCellGeneration(t *testing.T) {
	new_cell := NewDeadCell()
	assert.Equal(t, Dead, new_cell.status)
}

func TestIsAliveCheckForNewCell(t *testing.T) {
	new_cell := NewAliveCell()
	if !new_cell.IsAlive() {
		t.Errorf("expected alive cell, but got dead")
	}
}

func TestValidNumberOfLiveNeighbors(t *testing.T) {
	mainCell := NewAliveCell()

	mainCell.Neighbors = []*Cell{
		NewAliveCell(),
		NewDeadCell(),
		NewAliveCell(),
	}
	assert.Equal(t, 2, mainCell.NumberOfLiveNeighbors())
}
