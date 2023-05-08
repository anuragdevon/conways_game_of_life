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

	testSimulation := NewSimulation(test_grid)

	expectedNextStateGrid := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
	}

	testSimulation.nextStateOfSimulation()

	actualNextStateGrid := make([][]int, len(testSimulation.grid.Cells))
	for i := range testSimulation.grid.Cells {
		actualNextStateGrid[i] = make([]int, len(testSimulation.grid.Cells[i]))
		for j := range testSimulation.grid.Cells[i] {
			if testSimulation.grid.Cells[i][j].IsAlive() {
				actualNextStateGrid[i][j] = 1
			}
		}
	}

	if !reflect.DeepEqual(expectedNextStateGrid, actualNextStateGrid) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedNextStateGrid, actualNextStateGrid)
	}
}

func TestLiveCellToDieOfUnderpopulation(t *testing.T) {
	test_grid, _ := grid.NewGrid(3, 3)

	test_grid.Cells[0][0] = cell.NewAliveCell()
	test_grid.Cells[0][1] = cell.NewDeadCell()
	test_grid.Cells[0][2] = cell.NewDeadCell()

	test_grid.Cells[1][0] = cell.NewDeadCell()
	test_grid.Cells[1][1] = cell.NewAliveCell()
	test_grid.Cells[1][2] = cell.NewDeadCell()

	test_grid.Cells[2][0] = cell.NewDeadCell()
	test_grid.Cells[2][1] = cell.NewDeadCell()
	test_grid.Cells[2][2] = cell.NewDeadCell()

	expectedNextStateGrid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	testSimulation := NewSimulation(test_grid)

	testSimulation.nextStateOfSimulation()

	actualNextStateGrid := make([][]int, len(testSimulation.grid.Cells))
	for i := range testSimulation.grid.Cells {
		actualNextStateGrid[i] = make([]int, len(testSimulation.grid.Cells[i]))
		for j := range testSimulation.grid.Cells[i] {
			if testSimulation.grid.Cells[i][j].IsAlive() {
				actualNextStateGrid[i][j] = 1
			}
		}
	}

	if !reflect.DeepEqual(expectedNextStateGrid, actualNextStateGrid) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedNextStateGrid, actualNextStateGrid)
	}
}

func TestLiveCellToLiveForExactlyTwoOrThreeNeighbors(t *testing.T) {
	test_grid, _ := grid.NewGrid(3, 3)

	test_grid.Cells[0][0] = cell.NewDeadCell()
	test_grid.Cells[0][1] = cell.NewAliveCell()
	test_grid.Cells[0][2] = cell.NewDeadCell()

	test_grid.Cells[1][0] = cell.NewDeadCell()
	test_grid.Cells[1][1] = cell.NewDeadCell()
	test_grid.Cells[1][2] = cell.NewAliveCell()

	test_grid.Cells[2][0] = cell.NewAliveCell()
	test_grid.Cells[2][1] = cell.NewAliveCell()
	test_grid.Cells[2][2] = cell.NewAliveCell()

	expectedNextStateGrid := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
	}
	testSimulation := NewSimulation(test_grid)

	testSimulation.nextStateOfSimulation()

	actualNextStateGrid := make([][]int, len(testSimulation.grid.Cells))
	for i := range testSimulation.grid.Cells {
		actualNextStateGrid[i] = make([]int, len(testSimulation.grid.Cells[i]))
		for j := range testSimulation.grid.Cells[i] {
			if testSimulation.grid.Cells[i][j].IsAlive() {
				actualNextStateGrid[i][j] = 1
			}
		}
	}

	if !reflect.DeepEqual(expectedNextStateGrid, actualNextStateGrid) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedNextStateGrid, actualNextStateGrid)
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

	expectedNextStateGrid := [][]int{
		{1, 1, 1},
		{0, 0, 0},
		{1, 0, 1},
	}

	testSimulation.nextStateOfSimulation()

	actualNextStateGrid := make([][]int, len(testSimulation.grid.Cells))
	for i := range testSimulation.grid.Cells {
		actualNextStateGrid[i] = make([]int, len(testSimulation.grid.Cells[i]))
		for j := range testSimulation.grid.Cells[i] {
			if testSimulation.grid.Cells[i][j].IsAlive() {
				actualNextStateGrid[i][j] = 1
			}
		}
	}

	if !reflect.DeepEqual(expectedNextStateGrid, actualNextStateGrid) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedNextStateGrid, actualNextStateGrid)
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

	expectedNextStateGrid := [][]int{
		{1, 0, 1},
		{1, 0, 0},
		{1, 0, 1},
	}

	testSimulation.nextStateOfSimulation()

	actualNextStateGrid := make([][]int, len(testSimulation.grid.Cells))
	for i := range testSimulation.grid.Cells {
		actualNextStateGrid[i] = make([]int, len(testSimulation.grid.Cells[i]))
		for j := range testSimulation.grid.Cells[i] {
			if testSimulation.grid.Cells[i][j].IsAlive() {
				actualNextStateGrid[i][j] = 1
			}
		}
	}

	if !reflect.DeepEqual(expectedNextStateGrid, actualNextStateGrid) {
		t.Errorf("Expected next state grid:\n%v\nbut got:\n%v", expectedNextStateGrid, actualNextStateGrid)
	}
}
