package simulation

import (
	"conwaysgameoflife/lifematrix"
	"fmt"
)

type Simulation struct {
	grid *lifematrix.Grid
}

func NewSimulation(matrix *lifematrix.Grid) *Simulation {
	return &Simulation{
		grid: matrix,
	}
}

func (s *Simulation) NextStateOfSimulation() {
	next, _ := lifematrix.NewGrid(len(s.grid.Cells), len(s.grid.Cells[0]))
	for i := range s.grid.Cells {
		for j := range s.grid.Cells[i] {
			next.Cells[i][j] = 0
			liveNeighbors := s.grid.NumberOfLiveNeighbors(i, j)
			if (s.grid.Cells[i][j] == 1 && (liveNeighbors == 2 || liveNeighbors == 3)) || (s.grid.Cells[i][j] == 0 && liveNeighbors == 3) {
				next.Cells[i][j] = 1
			}
		}
	}
	s.grid = next
}

func (s *Simulation) DisplaySimulationGrid() {
	for i := range s.grid.Cells {
		for j := range s.grid.Cells[i] {
			if s.grid.Cells[i][j] == 0 {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
