package mocklife

import (
	"conwaysgameoflife/lifematrix"
)

type MockLife struct {
	grid *lifematrix.Grid
}

func CreateMockLife(life *lifematrix.Grid) *MockLife {
	return &MockLife{
		grid: life,
	}
}
