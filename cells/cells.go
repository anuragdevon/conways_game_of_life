package cell

type AliveStatus int

const (
	Dead AliveStatus = iota
	Alive
)

type Cell struct {
	status AliveStatus
}
