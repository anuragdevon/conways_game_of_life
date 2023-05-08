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
	return &Simulation{
		grid: matrix,
	}
}

func (s *Simulation) nextStateOfSimulation() {
	next, _ := grid.NewGrid(len(s.grid.Cells), len(s.grid.Cells[0]))
	for i := range s.grid.Cells {
		for j := range s.grid.Cells[i] {
			nextCell := cell.NewDeadCell()
			liveNeighbors := s.grid.NumberOfLiveNeighbors(i, j)
			if s.grid.Cells[i][j].IsAlive() && (liveNeighbors == 2 || liveNeighbors == 3) {
				nextCell = cell.NewAliveCell()
			} else if !s.grid.Cells[i][j].IsAlive() && liveNeighbors == 3 {
				nextCell = cell.NewAliveCell()
			}
			next.Cells[i][j] = nextCell
		}
	}
	s.grid = next
}

func (s *Simulation) displaySimulationGrid() {
	for i := range s.grid.Cells {
		for j := range s.grid.Cells[i] {
			if s.grid.Cells[i][j].IsAlive() {
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
