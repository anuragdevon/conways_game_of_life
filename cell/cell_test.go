package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAliveCellCreation(t *testing.T) {
	newCell := NewAliveCell()
	assert.Equal(t, Alive, newCell.status)
}

func TestNewDeadCellCreation(t *testing.T) {
	newCell := NewDeadCell()
	assert.Equal(t, Dead, newCell.status)
}

func TestIsAliveForNewAliveCell(t *testing.T) {
	newCell := NewAliveCell()
	if !newCell.IsAlive() {
		t.Errorf("expected alive cell, but got dead")
	}
}

func TestIsAliveForNewDeadCell(t *testing.T) {
	newCell := NewDeadCell()
	if newCell.IsAlive() {
		t.Errorf("expected dead cell, but got alive")
	}
}

func TestNumberOfLiveNeighborsReturnCountOfLiveNeighbors(t *testing.T) {
	targetCell := NewAliveCell()

	targetCell.Neighbors = []*Cell{
		NewAliveCell(),
		NewDeadCell(),
		NewAliveCell(),
	}
	assert.Equal(t, 2, targetCell.NumberOfLiveNeighbors())
}
