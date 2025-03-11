package entity

const (
	nodeCapacity = 2
)

type AirportMap struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

type AirportMapConfig struct {
	AirstripNodeId string
}

type Edge struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Distance float64 `json:"distance"`
}

func NewEdge(from string, to string, distance float64) *Edge {
	return &Edge{
		From:     from,
		To:       to,
		Distance: distance,
	}
}
