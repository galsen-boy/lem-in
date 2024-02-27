package lemin

import (
	"fmt"
	"os"
)

// checks if path intersects other paths in group
func intersects(group [][]string, path []string) bool {
	for i, room := range path {
		for _, pathG := range group {
			for _, roomG := range pathG {
				if room == roomG && i != len(path)-1 {
					return true
				}
			}
		}
	}
	return false
}

// PathGroups - groups of non-intersecting paths
var PathGroups = [][][]string{}

// GetNonInterPaths gets all non intersecting paths and groups them
func (g *Graph) GetNonInterPaths() {
	g.SortPaths()
	// trim start room bc we don't need it
	for i, p := range Paths {
		Paths[i] = p[1:]
	}

	if len(Paths) == 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	for i1, p1 := range Paths {
		var group = [][]string{}
		group = append(group, p1)
		PathGroups = append(PathGroups, group)
		for i2 := i1 + 1; i2 < len(Paths); i2++ {
			p2 := Paths[i2]
			if !intersects(group, p2) {
				group = append(group, p2)
			}
		}
		PathGroups = append(PathGroups, group)
	}

}

// SortPaths sorts paths by their length
func (g *Graph) SortPaths() {
	for i1 := 0; i1 < len(Paths); i1++ {
		for i2 := 0; i2 < len(Paths); i2++ {
			if len(Paths[i1]) < len(Paths[i2]) {
				Paths[i1], Paths[i2] = Paths[i2], Paths[i1]
			}
		}
	}
	// fmt.Println("sorted paths:", paths)
}
