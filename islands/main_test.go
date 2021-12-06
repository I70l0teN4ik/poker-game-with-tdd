package main

import (
	"reflect"
	"testing"
)

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

func Test_minIslandSize(t *testing.T) {
	tests := []struct {
		name string
		grid [][]string
		want int
	}{
		{"one row one isle", [][]string{{water, land, water, water}}, 1},
		{"one row 2 isles", [][]string{{water, land, land, water}}, 2},
		{"2 and 1", [][]string{{water, land, water, water, water}, {water, land, water, land, water}}, 1},
		{"3", [][]string{{land, water, water, water, water}, {land, land, water, water, water}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIslandSize(tt.grid); got != tt.want {
				t.Errorf("minIslandSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripIslandsOnEdges(t *testing.T) {
	tests := []struct {
		name string
		grid [][]string
		want [][]string
	}{
		{
			"one row one isle",
			[][]string{{water, land, water, water}},
			[][]string{{water, water, water, water}},
		},
		{
			"clears both rows",
			[][]string{{water, land, water, water, water}, {water, land, water, land, water}},
			[][]string{{water, water, water, water, water}, {water, water, water, water, water}},
		},
		{
			"top",
			[][]string{
				{land, water, land, water, water, water},
				{water, land, land, water, land, water},
				{water, land, water, land, land, water},
				{water, land, water, water, water, water},
			},
			[][]string{
				{water, water, water, water, water, water},
				{water, water, water, water, land, water},
				{water, water, water, land, land, water},
				{water, water, water, water, water, water},
			},
		},
		{
			"left edge",
			[][]string{
				{land, water, water, water, water, water},
				{land, land, land, water, land, water},
				{water, land, water, land, land, water},
				{water, land, water, water, water, water},
			},
			[][]string{
				{water, water, water, water, water, water},
				{water, water, water, water, land, water},
				{water, water, water, land, land, water},
				{water, water, water, water, water, water},
			},
		},
		{
			"bottom edge",
			[][]string{
				{land, water, water, water, water, water},
				{water, land, land, water, land, water},
				{water, land, water, land, land, water},
				{water, land, water, water, water, water},
			},
			[][]string{
				{water, water, water, water, water, water},
				{water, water, water, water, land, water},
				{water, water, water, land, land, water},
				{water, water, water, water, water, water},
			},
		},
		{
			"right edge",
			[][]string{
				{land, water, water, water, water, water},
				{water, land, land, water, land, land},
				{water, land, water, land, land, water},
				{water, water, water, water, water, water},
			},
			[][]string{
				{water, water, water, water, water, water},
				{water, land, land, water, land, land},
				{water, land, water, land, land, water},
				{water, water, water, water, water, water},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripIslandsOnEdges(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stripIslandsOnEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
