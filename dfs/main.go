package main

import "fmt"

func main() {
	nodes := [][]string{{"a", "b"}, {"a", "c"}, {"b", "d"}, {"c", "e"}, {"d", "f"}}
	graph := buildGraph(nodes, true)

	fmt.Println(hasPath(graph, "a", "f"))
	fmt.Println(hasPath(graph, "f", "f"))
	fmt.Println(hasPath(graph, "f", "a"))
	fmt.Println(hasPath(graph, "f", "e"))
}

func buildGraph(nodes [][]string, directed bool) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range nodes {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)

		if _, ok := graph[y]; !ok && directed {
			graph[y] = []string{}
		} else if !directed {
			graph[y] = append(graph[y], x)
		}
	}
	return graph
}

func hasPath(graph map[string][]string, root, target string) bool {
	nodes := dfsRecursive(graph, root)
	for _, v := range nodes {
		if target == v {
			return true
		}
	}
	return false
}

func dfsRecursive(graph map[string][]string, root string) []string {
	nodes := []string{root}
	for _, v := range graph[root] {
		nodes = append(nodes, dfsRecursive(graph, v)...)
	}
	return nodes
}

type stack struct {
	list []string
}

func (s *stack) push(i string) {
	s.list = append(s.list, i)
}

func (s *stack) pop() string {
	i := s.list[s.lastIndex()]
	s.list = s.list[:s.lastIndex()]
	return i
}

func (s *stack) len() int {
	return len(s.list)
}

func (s *stack) lastIndex() int {
	return s.len() - 1
}

func dfsIterative(graph map[string][]string, root string) []string {
	stack := stack{list: []string{root}}
	var nodes []string

	for stack.len() > 0 {
		current := stack.pop()
		nodes = append(nodes, current)
		for _, v := range graph[current] {
			stack.push(v)
		}
	}

	return nodes
}
