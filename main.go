package main

import (
	"conwaysgameoflife/lifematrix"
	"conwaysgameoflife/simulation"
	"fmt"
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

	for i := 0; ; i++ {
		fmt.Printf("Life Cycle %d\n", i+1)
		new_simulation.UpdateGameOfLife()
	}
}
