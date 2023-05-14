package cell

type AliveStatus int

const (
	Dead AliveStatus = iota
	Alive
)

type Cell struct {
	status    AliveStatus
	Neighbors []*Cell
}

func NewAliveCell() *Cell {
	return &Cell{
		status:    Alive,
		Neighbors: make([]*Cell, 0),
	}
}

func NewDeadCell() *Cell {
	return &Cell{
		status:    Dead,
		Neighbors: make([]*Cell, 0),
	}
}

func (c *Cell) IsAlive() bool {
	return c.status == Alive
}

func (c *Cell) NumberOfLiveNeighbors() int {
	countLiveNeighbors := 0
	for _, neighbor := range c.Neighbors {
		if neighbor.IsAlive() {
			countLiveNeighbors++
		}
	}
	return countLiveNeighbors
}
