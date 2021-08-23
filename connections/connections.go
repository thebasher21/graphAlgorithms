package connections

var Visited = make(map[int]bool)

func explore(adj map[int][]int, node int, vis map[int]bool) bool {
	if _, ok := vis[node]; ok {
		return false
	}
	vis[node] = true
	for _, neighbour := range adj[node] {
		explore(adj, neighbour, vis)
	}
	return true
}

func exploreSize(adj map[int][]int, node int, vis map[int]bool) int {
	if _, ok := vis[node]; ok {
		return 0
	}
	vis[node] = true
	size := 1
	for _, neighbour := range adj[node] {
		size += exploreSize(adj, neighbour, vis)
	}
	return size
}

func ConnectedComponentsCount(adj map[int][]int, vis map[int]bool) int {
	count := 0
	for node := range adj {
		if explore(adj, node, vis) {
			count++
		}
	}
	return count
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LargestComponent(adj map[int][]int) int {
	max := 0
	vis := map[int]bool{}
	for node := range adj {
		if _, ok := vis[node]; !ok {
			max = Max(max, exploreSize(adj, node, vis))
		}
	}
	return max
}
