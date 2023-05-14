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

func (c *Cell) AddNeighbor(neighbor *Cell) {
	c.neighbors = append(c.neighbors, neighbor)
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
