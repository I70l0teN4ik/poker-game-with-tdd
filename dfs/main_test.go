package main

import (
	"reflect"
	"testing"
)

func Test_hasPath(t *testing.T) {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}
	type args struct {
		graph  map[string][]string
		root   string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should find path", args{graph, "a", "f"}, true},
		{"should be ok for self", args{graph, "f", "f"}, true},
		{"should not find path", args{graph, "f", "a"}, false},
		{"should not find path", args{graph, "f", "e"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPath(tt.args.graph, tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_pop(t *testing.T) {
	t.Run("push increase size of the stuck", func(t *testing.T) {
		stack := stack{[]string{"a"}}
		stack.pop()
		if len(stack.list) != 0 {
			t.Errorf("did not change the size of stack")
		}
	})
	t.Run("pops elements in reverse order", func(t *testing.T) {
		want := "b"
		stack := stack{[]string{"a", want}}
		got := stack.pop()

		if want != got {
			t.Errorf("popped %s from stack wanted: %s", got, want)
		}
	})
}

func Test_stack_push(t *testing.T) {
	stack := stack{}
	t.Run("push increase size of the stuck", func(t *testing.T) {
		initLen := len(stack.list)
		stack.push("a")
		stack.push("b")
		if len(stack.list) != initLen+2 {
			t.Errorf("did not change the size of stack")
		}
	})
}

func Test_buildGraph(t *testing.T) {
	acNodes := [][]string{{"a", "b"}, {"a", "c"}, {"b", "d"}, {"c", "e"}, {"d", "f"}}
	cNodes := [][]string{{"a", "b"}, {"a", "c"}, {"b", "a"}, {"b", "d"}, {"c", "e"}, {"d", "f"}}
	tests := []struct {
		name      string
		nodes     [][]string
		directed  bool
		wantGraph map[string][]string
	}{
		{
			"acyclic",
			acNodes,
			true,
			map[string][]string{
				"a": {"b", "c"},
				"b": {"d"},
				"c": {"e"},
				"d": {"f"},
				"e": {},
				"f": {},
			},
		},
		{
			"cyclic",
			cNodes,
			true,
			map[string][]string{
				"a": {"b", "c"},
				"b": {"a", "d"},
				"c": {"e"},
				"d": {"f"},
				"e": {},
				"f": {},
			},
		},
		{
			"undirected acyclic",
			acNodes,
			false,
			map[string][]string{
				"a": {"b", "c"},
				"b": {"a", "d"},
				"c": {"a", "e"},
				"d": {"b", "f"},
				"e": {"c"},
				"f": {"d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGraph := buildGraph(tt.nodes, tt.directed); !reflect.DeepEqual(gotGraph, tt.wantGraph) {
				t.Errorf("buildGraph() = %v, want %v", gotGraph, tt.wantGraph)
			}
		})
	}
}
