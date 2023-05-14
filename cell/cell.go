package cell

type AliveStatus int

const (
	Dead AliveStatus = iota
	Alive
)

type Cell struct {
	status    AliveStatus
	neighbors []*Cell
}

func NewAliveCell() *Cell {
	return &Cell{
		status:    Alive,
		neighbors: make([]*Cell, 0),
	}
}

func NewDeadCell() *Cell {
	return &Cell{
		status:    Dead,
		neighbors: make([]*Cell, 0),
	}
}

func (c *Cell) IsAlive() bool {
	return c.status == Alive
}

func (c *Cell) numberOfLiveNeighbors() int {
	countLiveNeighbors := 0
	for _, neighbor := range c.neighbors {
		if neighbor.IsAlive() {
			countLiveNeighbors++
		}
	}
	return countLiveNeighbors
}

func (c *Cell) UpdateCellStatus() {
	numLiveNeighbors := c.numberOfLiveNeighbors()
	if c.IsAlive() {
		if numLiveNeighbors < 2 || numLiveNeighbors > 3 {
			c.status = Dead
		}
	}
}
