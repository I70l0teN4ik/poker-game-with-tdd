package main

import "testing"

func Test_findExitSteps(t *testing.T) {
	tests := []struct {
		name string
		want int
		maze [][]string
	}{
		{"no exit", -1, [][]string{{start, pass, pass, wall}, {wall, wall, wall, pass}}},
		{"straight", 4, [][]string{{start, pass, pass, pass, end}}},
		{"found exit", 9, [][]string{
			{start, pass, pass, wall, pass, pass, pass},
			{pass, wall, pass, pass, pass, wall, pass},
			{pass, wall, pass, pass, pass, pass, pass},
			{pass, pass, wall, wall, pass, pass, pass},
			{wall, pass, wall, end, pass, pass, pass},
		}},
		{"found exit", 11, [][]string{
			{start, pass, pass, wall, pass, pass, pass},
			{pass, wall, pass, pass, pass, wall, pass},
			{pass, wall, pass, pass, pass, pass, pass},
			{pass, pass, wall, wall, wall, pass, pass},
			{wall, pass, wall, end, pass, pass, pass},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findExitSteps(tt.maze); got != tt.want {
				t.Errorf("findExitSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
