package lemin

func hasSimilarPath(groups [][][]string, g1 [][]string) bool {
	for _, g2 := range groups {
		if len(g1) == len(g2) {
			var ok = make([]bool, len(g1))
			for i, p1 := range g1 {
				p2 := g2[i]
				if len(p1) == len(p2) || len(p1) > len(p2) {
					ok = append(ok, true)
				} else {
					ok = append(ok, false)
				}
			}
			count := 0
			for _, b := range ok {
				if b {
					count++
				}
			}
			if count == len(g1) {
				return true
			}
		}
	}
	return false
}

// RmSimilarPaths removes groups that have same length and same lengths of paths
func (g *Graph) RmSimilarPaths() {
	var newGroup = [][][]string{}
	for i1, g1 := range PathGroups {
		if len(g1) == 1 && i1 != 0 {
			continue
		}
		if !hasSimilarPath(newGroup, g1) {
			newGroup = append(newGroup, g1)
		}
	}
	PathGroups = newGroup
	// fmt.Println("cleaned path groups:", pathGroups)
}
