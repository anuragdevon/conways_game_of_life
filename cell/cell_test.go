package cell

import (
	"testing"
)

func TestValidNewAliveCellGeneration(t *testing.T) {
	new_cell := NewAliveCell()
	if new_cell.status != Alive {
		t.Errorf("expected alive cell, but got dead")
	}
}

func TestValidNewDeadCellGeneration(t *testing.T) {
	new_cell := NewDeadCell()
	if new_cell.status != Dead {
		t.Errorf("expected dead cell, but got alive")
	}
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

	if mainCell.NumberOfLiveNeighbors() != 2 {
		t.Errorf("Expected 2 live neighbors but got %d", mainCell.NumberOfLiveNeighbors())
	}
}
