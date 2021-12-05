package main

import "fmt"

func main() {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}
	dfsIterative(graph, "a")
	fmt.Println("")
	fmt.Print(dfsRecursive(graph, "a"))
	fmt.Println("")
	dfsRecursive(graph, "e")
	fmt.Println(hasPath(graph, "a", "f"))
	fmt.Println(hasPath(graph, "f", "f"))
	fmt.Println(hasPath(graph, "f", "a"))
	fmt.Println(hasPath(graph, "f", "e"))
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

func dfsIterative(graph map[string][]string, root string) {
	stack := stack{list: []string{root}}

	for stack.len() > 0 {
		current := stack.pop()
		fmt.Print(current)
		for _, v := range graph[current] {
			stack.push(v)
		}
	}
}
