package main

import "fmt"

func main() {
	nodes := buildTree()
	result := dfsTraversalRecursive(&nodes[0])

	fmt.Println(result)
}

func includesNode(root *Node, target string) bool {
	if nil == root {
		return false
	} else if target == root.val {
		return true
	}
	return includesNode(root.left, target) || includesNode(root.right, target)
}

func dfsTraversalRecursive(root *Node) []string {
	var result []string
	if nil == root {
		return result
	}
	result = append(result, root.val)
	result = append(result, dfsTraversalRecursive(root.left)...)
	result = append(result, dfsTraversalRecursive(root.right)...)

	return result
}

func dfsTraversalIterative(root *Node) []string {
	var result []string
	if nil == root {
		return result
	}
	stack := []*Node{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.val)

		if current.right != nil {
			stack = append(stack, current.right)
		}
		if current.left != nil {
			stack = append(stack, current.left)
		}
	}

	return result
}

func bfsTraversal(root *Node) []string {
	var result []string
	if nil == root {
		return result
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.val)

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return result
}

type Node struct {
	val   string
	left  *Node
	right *Node
}

func buildTree() []Node {
	d := Node{"d", nil, nil}
	e := Node{"e", nil, nil}
	f := Node{"f", nil, nil}
	b := Node{"b", &d, &e}
	c := Node{"c", &f, nil}
	a := Node{"a", &b, &c}

	return []Node{a, b, c, d, e, f}
}
