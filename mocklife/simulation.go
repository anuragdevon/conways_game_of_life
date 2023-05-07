package mocklife

import (
	"conwaysgameoflife/lifematrix"
)

type Simulation struct {
	grid *lifematrix.Grid
}

func NewSimulation(life *lifematrix.Grid) *Simulation {
	return &Simulation{
		grid: life,
	}
}
