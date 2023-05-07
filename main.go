package main

import (
	"conwaysgameoflife/lifematrix"
	"conwaysgameoflife/simulation"
	"fmt"
	"time"
)

func main() {
	rows := 3
	cols := 3

	grid, err := lifematrix.NewGrid(rows, cols)
	if err != nil {
		panic(err)
	}

	grid.Randomize()

	new_simulation := simulation.NewSimulation(grid)

	for i := 0; i < 10; i++ {
		fmt.Println("New Iteration")
		new_simulation.DisplaySimulationGrid()
		time.Sleep(1 * time.Second)
		new_simulation.NextStateOfSimulation()
	}
}
