package traversal

import (
	"fmt"
)

type Stack []string

func (s *Stack) Push(val string) {
	*s = append([]string{val}, *s...)
}

func (s *Stack) Pop() string {
	if len(*s) >= 1 {
		val := (*s)[0]
		*s = (*s)[1:]
		return val
	} else {
		return ""
	}
}

type Queue []string

func (q *Queue) Enqueue(val string) {
	*q = append([]string{val}, *q...)
}

func (q *Queue) Dequeue() string {
	if len(*q) >= 1 {
		val := (*q)[len(*q)-1]
		*q = (*q)[:len(*q)-1]
		return val
	} else {
		return ""
	}
}

func DFSIter(adj map[string][]string, start string) {
	var stack Stack
	stack.Push(start)
	for len(stack) > 0 {
		current := stack.Pop()
		fmt.Printf("%s -> ", current)
		for _, nodes := range adj[current] {
			stack.Push(nodes)
		}
	}
}

func DFSRec(adj map[string][]string, start string) {
	fmt.Printf("%s -> ", start)
	for _, nodes := range adj[start] {
		DFSRec(adj, nodes)
	}
}

func BFS(adj map[string][]string, start string) {
	var queue Queue
	queue.Enqueue(start)
	for len(queue) > 0 {
		current := queue.Dequeue()
		fmt.Printf("%s -> ", current)
		for _, nodes := range adj[current] {
			queue.Enqueue(nodes)
		}
	}
}
