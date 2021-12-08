package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_traverseTree(t *testing.T) {
	nodes := buildTree()
	type args struct {
	}
	tests := []struct {
		name string
		root *Node
		want []string
	}{
		{"from root", &nodes[0], []string{"a", "b", "d", "e", "c", "f"}},
		{"from b", &nodes[1], []string{"b", "d", "e"}},
		{"from c", &nodes[2], []string{"c", "f"}},
	}
	for _, tt := range tests {
		t.Run("recursive dfs"+tt.name, func(t *testing.T) {
			if got := dfsTraversalRecursive(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run("iterative dfs"+tt.name, func(t *testing.T) {
			if got := dfsTraversalIterative(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}

	bfsTests := []struct {
		name string
		root *Node
		want []string
	}{
		{"from root", &nodes[0], []string{"a", "b", "c", "d", "e", "f"}},
		{"from b", &nodes[1], []string{"b", "d", "e"}},
		{"from c", &nodes[2], []string{"c", "f"}},
	}
	for _, tt := range bfsTests {
		t.Run("bfs "+tt.name, func(t *testing.T) {
			if got := bfsTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}

	includesTests := []struct {
		name   string
		root   *Node
		target string
		want   bool
	}{
		{"from root", &nodes[0], "d", true},
		{"from b", &nodes[1], "e", true},
		{"from b", &nodes[1], "f", false},
		{"from c", &nodes[2], "b", false},
		{"from c", &nodes[2], "f", true},
	}
	for _, tt := range includesTests {
		t.Run(fmt.Sprintf("%s includes %s %v", tt.root.val, tt.target, tt.want), func(t *testing.T) {
			if got := includesNode(tt.root, tt.target); got != tt.want {
				t.Errorf("dfsTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
