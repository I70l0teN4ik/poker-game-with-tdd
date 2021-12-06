package main

import "testing"

func Test_islandCount(t *testing.T) {
	tests := []struct {
		name string
		grid [][]string
		want int
	}{
		{"one row one isle", [][]string{{water, land, water, water}}, 1},
		{"one row 2 isles", [][]string{{water, land, water, land, water}}, 2},
		{"2 isles", [][]string{{water, land, water, water, water}, {water, land, water, land, water}}, 2},
		{"3 isles", [][]string{{land, water, water, water, water}, {water, land, water, land, water}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandCount(tt.grid); got != tt.want {
				t.Errorf("islandCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
