package path

import (
	"graph/traversal"
)

func HasPathDFS(adj map[string][]string, source, destination string) bool {
	for _, node := range adj[source] {
		if node == destination {
			return true
		}
		if HasPathDFS(adj, node, destination) {
			return true
		} else {
			continue
		}
	}
	return false
}

func HasPathBFS(adj map[string][]string, source, destination string) bool {
	var queue traversal.Queue
	queue.Enqueue(source)
	for len(queue) > 0 {
		current := queue.Dequeue()
		if current == destination {
			return true
		}
		for _, node := range adj[current] {
			queue.Enqueue(node)
		}
	}
	return false
}

var visited = make(map[string]bool)

func UndirectedPath(adj map[string][]string, source, destination string) bool {
	if source == destination {
		return true
	}
	for _, node := range adj[source] {
		if _, exists := visited[node]; !exists {
			visited[node] = true
			if UndirectedPath(adj, node, destination) {
				return true
			}
		}
	}
	return false
}

func ShortestPath(adj map[string][]string, source, destination string) int {
	var queue [][]interface{}
	dist := 0
	vis := map[string]bool{}
	vis[source] = true
	queue = append(queue, []interface{}{source, dist})
	for len(queue) > 0 {
		popped := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		current, distance := popped[0].(string), popped[1].(int)
		if current == destination {
			return distance
		}
		for _, node := range adj[current] {
			if _, ok := vis[node]; !ok {
				vis[node] = true
				queue = append(queue, []interface{}{node, distance + 1})
			}
		}
	}
	return -1
}
