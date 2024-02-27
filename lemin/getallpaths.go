package lemin

var visited = make(map[string]bool)
var path = []string{}

// Paths - all found paths
var Paths = [][]string{}

// GetAllPaths implements DFS algorithm to find all paths from start to end rooms
func (g *Graph) GetAllPaths(room string) {
	if visited[room] {
		return
	}
	visited[room] = true
	path = append(path, room)
	if room != g.End {
		for _, neighbor := range g.Edges[room] {
			g.GetAllPaths(neighbor)
		}
	} else {
		var copy = make([]string, len(path))
		for i, v := range path {
			copy[i] = v
		}
		Paths = append(Paths, copy)
	}
	path = path[:len(path)-1]
	visited[room] = false
}
