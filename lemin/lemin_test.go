package lemin

import (
	"testing"
)

func TestGetAllPaths(t *testing.T) {
	// Création d'un graphe de test
	graph := Graph{
		Start: "A",
		End:   "D",
		Edges: map[string][]string{
			"A": {"B", "C"},
			"B": {"C", "D"},
			"C": {"D"},
			"D": {},
		},
	}

	// Appel de la méthode GetAllPaths
	graph.GetAllPaths(graph.Start)

	// Vérification des chemins trouvés
	expectedPaths := [][]string{
		{"A", "B", "C", "D"},
		{"A", "B", "D"},
		{"A", "C", "D"},
	}
	if len(Paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, but got %d", len(expectedPaths), len(Paths))
	}

	for i, expectedPath := range expectedPaths {
		if len(Paths[i]) != len(expectedPath) {
			t.Errorf("Expected path length %d, but got %d", len(expectedPath), len(Paths[i]))
		}

		for j, expectedNode := range expectedPath {
			if Paths[i][j] != expectedNode {
				t.Errorf("Expected node %s, but got %s", expectedNode, Paths[i][j])
			}
		}
	}
}

func TestSortPaths(t *testing.T) {
	// Mocking Paths
	Paths = [][]string{
		{"A", "B", "C", "D"},
		{"A", "B", "D"},
		{"A", "C", "D"},
	}

	// Appel de la méthode SortPaths
	var graph Graph
	graph.SortPaths()

	// Vérification de l'ordre des chemins
	expectedPaths := [][]string{
		{"A", "B", "D"},
		{"A", "C", "D"},
		{"A", "B", "C", "D"},
	}
	for i := range Paths {
		if len(Paths[i]) != len(expectedPaths[i]) {
			t.Errorf("Expected path length %d, but got %d", len(expectedPaths[i]), len(Paths[i]))
		}

		for j := range Paths[i] {
			if Paths[i][j] != expectedPaths[i][j] {
				t.Errorf("Expected node %s, but got %s", expectedPaths[i][j], Paths[i][j])
			}
		}
	}
}

func TestIntersects(t *testing.T) {
	// Mocking a group and a path
	group := [][]string{
		{"A", "B", "C", "D"},
		{"B", "C", "E"},
		{"C", "D", "E"},
	}
	path := []string{"B", "C", "D"}

	// Appel de la fonction intersects
	result := intersects(group, path)

	// Vérification du résultat attendu
	expectedResult := true
	if result != expectedResult {
		t.Errorf("Expected intersects(group, path) to be %v, but got %v", expectedResult, result)
	}
}
