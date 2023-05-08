package cell

import (
	"testing"
)

func TestValidNewDeadCellGeneration(t *testing.T) {
	new_cell := NewCell()
	if new_cell.status != Dead {
		t.Errorf("expected a dead cell to be created")
	}
}

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
