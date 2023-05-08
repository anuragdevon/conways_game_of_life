package main

import (
	"conwaysgameoflife/grid"
	"conwaysgameoflife/simulation"
	"fmt"
)

func main() {
	var rows, cols int
	fmt.Printf("Enter number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter number of cols: ")
	fmt.Scanln(&cols)

	matrix, err := grid.NewGrid(rows, cols)
	if err != nil {
		panic(err)
	}

	matrix.Randomize()

	new_simulation := simulation.NewSimulation(matrix)

	for i := 0; ; i++ {
		fmt.Printf("Life Cycle %d\n", i+1)
		new_simulation.UpdateGameOfLife()
	}
}
