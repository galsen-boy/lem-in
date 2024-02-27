package lemin

// Graph : start, end - source & sink, edges - connections between rooms
type Graph struct {
	Start, End string
	Edges      map[string][]string
}

// Connect rooms
func (g *Graph) Connect(r1, r2 string) {
	if g.Edges == nil {
		g.Edges = make(map[string][]string)
	}
	g.Edges[r1] = append(g.Edges[r1], r2)
	g.Edges[r2] = append(g.Edges[r2], r1)
}

// Ant : path, current room index, previous room, ignore
type Ant struct {
	Path     []string
	RoomID   int
	Previous string
	Ignore   bool
}

// Task is content of file in arguments
var Task string
