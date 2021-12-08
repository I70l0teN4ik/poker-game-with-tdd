package main

import (
	"fmt"
	"math"
)

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

type NumNode struct {
	val   int
	left  *NumNode
	right *NumNode
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

func buildNumTree(numbers []int) []NumNode {
	f := NumNode{5, nil, nil}
	e := NumNode{4, nil, nil}
	d := NumNode{3, nil, nil}
	c := NumNode{2, &f, nil}
	b := NumNode{1, &d, &e}
	a := NumNode{0, &b, &c}

	return []NumNode{a, b, c, d, e, f}
}

func dfsTraversal(root *NumNode) []int {
	var result []int
	if nil == root {
		return result
	}
	result = append(result, root.val)
	result = append(result, dfsTraversal(root.left)...)
	result = append(result, dfsTraversal(root.right)...)

	return result
}

func nodeSum(root *NumNode) int {
	if nil == root {
		return 0
	}
	return root.val + nodeSum(root.left) + nodeSum(root.right)
}

func treeMinVal(root *NumNode) int {
	min := math.MaxInt
	if nil == root {
		return min
	}
	queue := []*NumNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.val < min {
			min = node.val
		}

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return min
}

func treeMinValRec(node *NumNode) int {
	if nil == node {
		return math.MaxInt
	}

	return min(node.val, min(treeMinValRec(node.left), treeMinValRec(node.right)))
}

func treeMaxPathRec(node *NumNode) int {
	if nil == node {
		return math.MinInt
	}
	if nil == node.left && nil == node.right {
		return node.val
	}

	return node.val + max(treeMaxPathRec(node.left), treeMaxPathRec(node.right))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
