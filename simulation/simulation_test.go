package simulation

import (
	"conwaysgameoflife/cell"
	"conwaysgameoflife/grid"
	"reflect"
	"testing"
)

func TestCreateValidNewSimulation(t *testing.T) {
	rows := 10
	cols := 10
	testGrid, _ := grid.NewGrid(rows, cols)

	testNewSimulation := NewSimulation(testGrid)
	if testNewSimulation.grid != testGrid {
		t.Errorf("expected both grids to be same")
	}
}

func TestCreateValidNextStateOfSimulation(t *testing.T) {
	rows := 3
	cols := 3
	testGrid, _ := grid.NewGrid(rows, cols)

	testGrid.Cells[0][0] = cell.NewDeadCell()
	testGrid.Cells[0][1] = cell.NewAliveCell()
	testGrid.Cells[0][2] = cell.NewDeadCell()

	testGrid.Cells[1][0] = cell.NewDeadCell()
	testGrid.Cells[1][1] = cell.NewDeadCell()
	testGrid.Cells[1][2] = cell.NewAliveCell()

	testGrid.Cells[2][0] = cell.NewAliveCell()
	testGrid.Cells[2][1] = cell.NewAliveCell()
	testGrid.Cells[2][2] = cell.NewAliveCell()

	testSimulation := NewSimulation(testGrid)

	expectedStateGrid, _ := grid.NewGrid(rows, cols)
	expectedStateGrid.Cells[0][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[0][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[0][2] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][0] = cell.NewAliveCell()
	expectedStateGrid.Cells[1][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][2] = cell.NewAliveCell()
	expectedStateGrid.Cells[2][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][1] = cell.NewAliveCell()
	expectedStateGrid.Cells[2][2] = cell.NewAliveCell()

	testSimulation.nextStateOfSimulation()

	if !reflect.DeepEqual(expectedStateGrid.Cells, testSimulation.grid.Cells) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedStateGrid.Cells, testSimulation.grid.Cells)
	}
}

func TestLiveCellToDieOfUnderpopulation(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := grid.NewGrid(rows, cols)

	test_grid.Cells[0][0] = cell.NewAliveCell()
	test_grid.Cells[0][1] = cell.NewDeadCell()
	test_grid.Cells[0][2] = cell.NewDeadCell()

	test_grid.Cells[1][0] = cell.NewDeadCell()
	test_grid.Cells[1][1] = cell.NewAliveCell()
	test_grid.Cells[1][2] = cell.NewDeadCell()

	test_grid.Cells[2][0] = cell.NewDeadCell()
	test_grid.Cells[2][1] = cell.NewDeadCell()
	test_grid.Cells[2][2] = cell.NewDeadCell()

	expectedStateGrid, _ := grid.NewGrid(rows, cols)
	expectedStateGrid.Cells[0][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[0][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[0][2] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][2] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][2] = cell.NewDeadCell()

	testSimulation := NewSimulation(test_grid)

	testSimulation.nextStateOfSimulation()

	if !reflect.DeepEqual(expectedStateGrid.Cells, testSimulation.grid.Cells) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedStateGrid.Cells, testSimulation.grid.Cells)
	}
}

func TestLiveCellToLiveForExactlyTwoOrThreeNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := grid.NewGrid(rows, cols)

	test_grid.Cells[0][0] = cell.NewDeadCell()
	test_grid.Cells[0][1] = cell.NewAliveCell()
	test_grid.Cells[0][2] = cell.NewDeadCell()

	test_grid.Cells[1][0] = cell.NewDeadCell()
	test_grid.Cells[1][1] = cell.NewDeadCell()
	test_grid.Cells[1][2] = cell.NewAliveCell()

	test_grid.Cells[2][0] = cell.NewAliveCell()
	test_grid.Cells[2][1] = cell.NewAliveCell()
	test_grid.Cells[2][2] = cell.NewAliveCell()

	expectedStateGrid, _ := grid.NewGrid(rows, cols)
	expectedStateGrid.Cells[0][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[0][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[0][2] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][0] = cell.NewAliveCell()
	expectedStateGrid.Cells[1][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][2] = cell.NewAliveCell()
	expectedStateGrid.Cells[2][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][1] = cell.NewAliveCell()
	expectedStateGrid.Cells[2][2] = cell.NewAliveCell()

	testSimulation := NewSimulation(test_grid)

	testSimulation.nextStateOfSimulation()

	if !reflect.DeepEqual(expectedStateGrid.Cells, testSimulation.grid.Cells) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedStateGrid.Cells, testSimulation.grid.Cells)
	}
}

func TestLiveCellToDieOfOverpopulation(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := grid.NewGrid(rows, cols)

	test_grid.Cells[0][0] = cell.NewAliveCell()
	test_grid.Cells[0][1] = cell.NewAliveCell()
	test_grid.Cells[0][2] = cell.NewDeadCell()

	test_grid.Cells[1][0] = cell.NewDeadCell()
	test_grid.Cells[1][1] = cell.NewAliveCell()
	test_grid.Cells[1][2] = cell.NewAliveCell()

	test_grid.Cells[2][0] = cell.NewAliveCell()
	test_grid.Cells[2][1] = cell.NewAliveCell()
	test_grid.Cells[2][2] = cell.NewAliveCell()

	testSimulation := NewSimulation(test_grid)

	expectedStateGrid, _ := grid.NewGrid(rows, cols)
	expectedStateGrid.Cells[0][0] = cell.NewAliveCell()
	expectedStateGrid.Cells[0][1] = cell.NewAliveCell()
	expectedStateGrid.Cells[0][2] = cell.NewAliveCell()
	expectedStateGrid.Cells[1][0] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][2] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][0] = cell.NewAliveCell()
	expectedStateGrid.Cells[2][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][2] = cell.NewAliveCell()

	testSimulation.nextStateOfSimulation()

	if !reflect.DeepEqual(expectedStateGrid.Cells, testSimulation.grid.Cells) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedStateGrid.Cells, testSimulation.grid.Cells)
	}
}

func TestDeadToComeAliveForExactlyThreeNeighbors(t *testing.T) {
	rows := 3
	cols := 3
	test_grid, _ := grid.NewGrid(rows, cols)

	test_grid.Cells[0][0] = cell.NewAliveCell()
	test_grid.Cells[0][1] = cell.NewDeadCell()
	test_grid.Cells[0][2] = cell.NewAliveCell()

	test_grid.Cells[1][0] = cell.NewAliveCell()
	test_grid.Cells[1][1] = cell.NewAliveCell()
	test_grid.Cells[1][2] = cell.NewAliveCell()

	test_grid.Cells[2][0] = cell.NewDeadCell()
	test_grid.Cells[2][1] = cell.NewAliveCell()
	test_grid.Cells[2][2] = cell.NewAliveCell()

	testSimulation := NewSimulation(test_grid)

	expectedStateGrid, _ := grid.NewGrid(rows, cols)
	expectedStateGrid.Cells[0][0] = cell.NewAliveCell()
	expectedStateGrid.Cells[0][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[0][2] = cell.NewAliveCell()
	expectedStateGrid.Cells[1][0] = cell.NewAliveCell()
	expectedStateGrid.Cells[1][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[1][2] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][0] = cell.NewAliveCell()
	expectedStateGrid.Cells[2][1] = cell.NewDeadCell()
	expectedStateGrid.Cells[2][2] = cell.NewAliveCell()

	testSimulation.nextStateOfSimulation()

	if !reflect.DeepEqual(expectedStateGrid.Cells, testSimulation.grid.Cells) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedStateGrid.Cells, testSimulation.grid.Cells)
	}
}
