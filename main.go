package main

import (
	"fmt"
	"graph/connections"
	"graph/grid"
	"graph/path"
	"graph/traversal"
)

func edgeListToAdjacencyList(edgeList [][]string, isDirected bool) map[string][]string {
	adj := make(map[string][]string)
	if !isDirected {
		for _, edge := range edgeList {
			adj[edge[0]] = append(adj[edge[0]], edge[1])
			adj[edge[1]] = append(adj[edge[1]], edge[0])
		}
	} else {
		for _, edge := range edgeList {
			adj[edge[0]] = append(adj[edge[0]], edge[1])
		}
	}

	return adj
}

func main() {
	adj := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}
	adj1 := map[string][]string{
		"f": {"g", "i"},
		"g": {"h"},
		"h": {},
		"i": {"g", "k"},
		"j": {"i"},
		"k": {},
	}
	edg := [][]string{
		{"i", "j"}, {"k", "i"}, {"m", "k"}, {"k", "l"}, {"o", "n"},
	}
	edjToAdj := edgeListToAdjacencyList(edg, false)
	/* adj2 := map[int][]int{
		3: {},
		4: {6},
		6: {4, 5, 7, 8},
		8: {6},
		7: {6},
		5: {6},
		1: {2},
		2: {1},
	} */
	adj3 := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0, 8},
		8: {0, 5},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}
	edg2 := [][]string{
		{"w", "x"}, {"x", "y"}, {"z", "y"}, {"z", "v"}, {"w", "v"},
	}
	edgToAdj := edgeListToAdjacencyList(edg2, false)
	grid1 := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "L", "W"},
		{"W", "W", "L", "L", "W"},
		{"L", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "W"},
	}
	fmt.Println("DFS:")
	traversal.DFSIter(adj, "a")
	fmt.Println("\nDFS:")
	traversal.DFSRec(adj, "a")
	fmt.Println("\nBFS:")
	traversal.BFS(adj, "a")
	fmt.Println("\nPath Exists:")
	fmt.Println(path.HasPathDFS(adj1, "f", "k"))
	fmt.Println(path.HasPathBFS(adj1, "f", "k"))
	fmt.Println("Undirected Path:", path.UndirectedPath(edjToAdj, "j", "m"))
	fmt.Println("Total components:", connections.ConnectedComponentsCount(adj3, connections.Visited))
	fmt.Println("Largest component:", connections.LargestComponent(adj3))
	fmt.Println("Shortest Path:", path.ShortestPath(edgToAdj, "w", "z"))
	fmt.Println("Island Count:", grid.IslandCount(grid1))
	fmt.Println("Smallest Island:", grid.SmallestIsland(grid1))
}
