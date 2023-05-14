package simulation

import (
	"conwaysgameoflife/cell"
	"conwaysgameoflife/grid"
	"fmt"
	"time"
)

type Simulation struct {
	grid *grid.Grid
}

func NewSimulation(matrix *grid.Grid) *Simulation {
	matrix.AddCellNeighbors()
	return &Simulation{
		grid: matrix,
	}

}

func (s *Simulation) nextStateOfSimulation() {
	nextSimGrid, _ := grid.NewGrid(len(s.grid.Cells), len(s.grid.Cells[0]))
	for currentRow := range s.grid.Cells {
		for currentCol := range s.grid.Cells[currentRow] {
			nextCell := cell.NewDeadCell()
			liveNeighbors := s.grid.Cells[currentRow][currentCol].NumberOfLiveNeighbors()
			if s.grid.Cells[currentRow][currentCol].IsAlive() && (liveNeighbors == 2 || liveNeighbors == 3) {
				nextCell = cell.NewAliveCell()
			} else if !s.grid.Cells[currentRow][currentCol].IsAlive() && liveNeighbors == 3 {
				nextCell = cell.NewAliveCell()
			}
			nextSimGrid.Cells[currentRow][currentCol] = nextCell
		}
	}
	nextSimGrid.AddCellNeighbors()
	s.grid = nextSimGrid
}

func (s *Simulation) displaySimulationGrid() {
	for currentRow := range s.grid.Cells {
		for currentCol := range s.grid.Cells[currentRow] {
			if s.grid.Cells[currentRow][currentCol].IsAlive() {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (s *Simulation) UpdateGameOfLife() {
	s.displaySimulationGrid()
	time.Sleep(1 * time.Second)
	s.nextStateOfSimulation()
}
