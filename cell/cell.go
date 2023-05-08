package cell

type AliveStatus int

const (
	Dead AliveStatus = iota
	Alive
)

type Cell struct {
	status AliveStatus
}

func NewAliveCell() *Cell {
	return &Cell{status: Alive}
}

func NewDeadCell() *Cell {
	return &Cell{status: Dead}
}

func (c *Cell) IsAlive() bool {
	return c.status == Alive
}
