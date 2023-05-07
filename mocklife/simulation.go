package mocklife

import (
	"conwaysgameoflife/lifematrix"
)

type Simulation struct {
	grid *lifematrix.Grid
}

func NewSimulation(matrix *lifematrix.Grid) *Simulation {
	return &Simulation{
		grid: matrix,
	}
}
