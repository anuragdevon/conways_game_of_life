package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	mainCell.neighbors = append(mainCell.neighbors, NewAliveCell())
	mainCell.neighbors = append(mainCell.neighbors, NewDeadCell())
	mainCell.neighbors = append(mainCell.neighbors, NewAliveCell())

	if mainCell.numberOfLiveNeighbors() != 2 {
		t.Errorf("Expected 2 live neighbors but got %d", mainCell.numberOfLiveNeighbors())
	}
}

func TestUpdateCellStatusOfAliveCellWithLessThanTwoLiveNeighbors(t *testing.T) {
	aliveCellWithFewerThanTwoLiveNeighbors := NewAliveCell()
	aliveCellWithFewerThanTwoLiveNeighbors.neighbors = []*Cell{
		NewDeadCell(),
		NewDeadCell(),
		NewDeadCell(),
	}
	aliveCellWithFewerThanTwoLiveNeighbors.UpdateCellStatus()
	assert.Equal(t, aliveCellWithFewerThanTwoLiveNeighbors.IsAlive(), false)
}

func TestUpdateCellStatusOfAliveCellWithMoreThanThreeLiveNeighbors(t *testing.T) {
	aliveCellWithMoreThanThreeLiveNeighbors := NewAliveCell()
	aliveCellWithMoreThanThreeLiveNeighbors.neighbors = []*Cell{
		NewAliveCell(),
		NewAliveCell(),
		NewAliveCell(),
		NewAliveCell(),
	}
	aliveCellWithMoreThanThreeLiveNeighbors.UpdateCellStatus()
	assert.Equal(t, aliveCellWithMoreThanThreeLiveNeighbors.IsAlive(), false)
}
