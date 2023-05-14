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

func TestValidAdditionOfNeighbors(t *testing.T) {
	mainCell := NewAliveCell()
	neighborCell := NewDeadCell()

	mainCell.AddNeighbor(neighborCell)

	if len(mainCell.neighbors) != 1 || mainCell.neighbors[0] != neighborCell {
		t.Errorf("got unexpected neighbors")
	}
}
