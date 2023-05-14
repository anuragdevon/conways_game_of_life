package simulation

import (
	"conwaysgameoflife/cell"
	"conwaysgameoflife/grid"
	"reflect"
	"testing"
)

func TestNewSimulationCreatesValidSimulation(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, _ := grid.NewGrid(rows, cols)

	testNewSimulation := NewSimulation(testGrid)
	if testNewSimulation.grid != testGrid {
		t.Errorf("expected both grids to be same")
	}
}

func TestNextStateOfSimulationForLiveCellToDieForLessThanTwoNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, _ := grid.NewGrid(rows, cols)

	testGrid.Cells = [][]*cell.Cell{
		{cell.NewAliveCell(), cell.NewDeadCell(), cell.NewDeadCell()},
		{cell.NewDeadCell(), cell.NewAliveCell(), cell.NewDeadCell()},
		{cell.NewDeadCell(), cell.NewDeadCell(), cell.NewDeadCell()},
	}

	expectedGridState, _ := grid.NewGrid(rows, cols)
	expectedGridState.Cells = [][]*cell.Cell{
		{cell.NewDeadCell(), cell.NewDeadCell(), cell.NewDeadCell()},
		{cell.NewDeadCell(), cell.NewDeadCell(), cell.NewDeadCell()},
		{cell.NewDeadCell(), cell.NewDeadCell(), cell.NewDeadCell()},
	}
	expectedSimulationState := NewSimulation(expectedGridState)

	testSimulation := NewSimulation(testGrid)
	testSimulation.nextStateOfSimulation()

	if !reflect.DeepEqual(expectedSimulationState.grid, testSimulation.grid) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedSimulationState.grid, testSimulation.grid)
	}
}

func TestNextStateOfSimulationForLiveCellToLiveForTwoOrThreeNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, _ := grid.NewGrid(rows, cols)

	testGrid.Cells = [][]*cell.Cell{
		{cell.NewDeadCell(), cell.NewAliveCell(), cell.NewDeadCell()},
		{cell.NewDeadCell(), cell.NewDeadCell(), cell.NewAliveCell()},
		{cell.NewAliveCell(), cell.NewAliveCell(), cell.NewAliveCell()},
	}

	expectedGridState, _ := grid.NewGrid(rows, cols)
	expectedGridState.Cells = [][]*cell.Cell{
		{cell.NewDeadCell(), cell.NewDeadCell(), cell.NewDeadCell()},
		{cell.NewAliveCell(), cell.NewDeadCell(), cell.NewAliveCell()},
		{cell.NewDeadCell(), cell.NewAliveCell(), cell.NewAliveCell()},
	}
	expectedSimulationState := NewSimulation(expectedGridState)

	testSimulation := NewSimulation(testGrid)

	testSimulation.nextStateOfSimulation()

	if !reflect.DeepEqual(expectedSimulationState.grid, testSimulation.grid) {
		t.Errorf("Expected next state:\n%v\nbut got:\n%v", expectedSimulationState.grid, testSimulation.grid)
	}
}

func TestNextStateOfSimulationForLiveCellToDieForMoreThanThreeNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, _ := grid.NewGrid(rows, cols)

	testGrid.Cells = [][]*cell.Cell{
		{cell.NewAliveCell(), cell.NewAliveCell(), cell.NewDeadCell()},
		{cell.NewDeadCell(), cell.NewAliveCell(), cell.NewAliveCell()},
		{cell.NewAliveCell(), cell.NewAliveCell(), cell.NewAliveCell()},
	}

	testSimulation := NewSimulation(testGrid)
	testSimulation.nextStateOfSimulation()

	expectedGridState, _ := grid.NewGrid(rows, cols)
	expectedGridState.Cells = [][]*cell.Cell{
		{cell.NewAliveCell(), cell.NewAliveCell(), cell.NewAliveCell()},
		{cell.NewDeadCell(), cell.NewDeadCell(), cell.NewDeadCell()},
		{cell.NewAliveCell(), cell.NewDeadCell(), cell.NewAliveCell()},
	}
	expectedSimulationState := NewSimulation(expectedGridState)

	if !reflect.DeepEqual(expectedGridState, testSimulation.grid) {
		t.Errorf("Expected next state:\n%v\nbut got:\n%v", expectedSimulationState.grid, testSimulation.grid)
	}
}

func TestNextStateOfSimulationForDeadToComeAliveForExactlyThreeNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, _ := grid.NewGrid(rows, cols)

	testGrid.Cells = [][]*cell.Cell{
		{cell.NewAliveCell(), cell.NewDeadCell(), cell.NewAliveCell()},
		{cell.NewAliveCell(), cell.NewAliveCell(), cell.NewAliveCell()},
		{cell.NewDeadCell(), cell.NewAliveCell(), cell.NewAliveCell()},
	}

	testSimulation := NewSimulation(testGrid)
	testSimulation.nextStateOfSimulation()

	expectedGridState, _ := grid.NewGrid(rows, cols)
	expectedGridState.Cells = [][]*cell.Cell{
		{cell.NewAliveCell(), cell.NewDeadCell(), cell.NewAliveCell()},
		{cell.NewAliveCell(), cell.NewDeadCell(), cell.NewDeadCell()},
		{cell.NewAliveCell(), cell.NewDeadCell(), cell.NewAliveCell()},
	}
	expectedSimulationState := NewSimulation(expectedGridState)

	if !reflect.DeepEqual(expectedGridState, testSimulation.grid) {
		t.Errorf("Expected next state:\n%v\nbut got:\n%v", expectedSimulationState.grid, testSimulation.grid)
	}
}
