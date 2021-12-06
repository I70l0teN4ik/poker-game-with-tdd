package main

import (
	"fmt"
	"math"
)

const (
	water = "w"
	land  = "l"
)

func main() {
	grid := [][]string{{water, land, water, water}}
	fmt.Println(islandCount(grid))
}

func islandCount(grid [][]string) (counter int) {
	visited := initVisitedGrid(grid)

	for r := range grid {
		for c := range grid[r] {
			if explore(grid, r, c, visited) > 0 {
				counter++
			}
		}
	}

	return
}

func minIslandSize(grid [][]string) int {
	visited := initVisitedGrid(grid)
	size := math.MaxInt

	for r := range grid {
		for c := range grid[r] {
			if currSize := explore(grid, r, c, visited); currSize > 0 && currSize < size {
				size = currSize
			}
		}
	}
	return size
}

func stripIslandsOnEdges(grid [][]string) [][]string {
	visited := initVisitedGrid(grid)

	for r := range grid {
		for c := range grid[r] {
			explore(grid, r, c, visited)
		}
	}

	for r := range grid {
		for c := range grid[r] {
			if visited[r][c].toRemove {
				grid[r][c] = water
			}
		}
	}
	return grid
}

var dr, dc = [4]int{-1, 1, 0, 0}, [4]int{0, 0, -1, 1}

func explore(grid [][]string, r int, c int, visited [][]node) (size int) {
	rz := len(grid)
	cz := len(grid[0])
	if outOfGrid(r, rz, c, cz) {
		return
	}
	if water == grid[r][c] {
		return
	}
	if visited[r][c].visited {
		return
	}
	visited[r][c].visited = true
	size++

	if onTheEdge(r, rz, c, cz) || connectedToEdge(r, c, visited) {
		visited[r][c].toRemove = true
	}

	for i := 0; i < 4; i++ {
		rr := r + dr[i]
		cc := c + dc[i]
		size += explore(grid, rr, cc, visited)
	}

	return
}

func connectedToEdge(r int, c int, visited [][]node) bool {
	for i := 0; i < 4; i++ {
		rr := r + dr[i]
		cc := c + dc[i]
		if outOfGrid(rr, len(visited), cc, len(visited[0])) {
			continue
		}
		if visited[rr][cc].toRemove {
			return true
		}
	}
	return false
}

func outOfGrid(r int, R int, c int, C int) bool {
	return r < 0 || r >= R || c < 0 || c >= C
}

func onTheEdge(r int, rowsSize int, c int, columnsSize int) bool {
	return r == 0 || r == (rowsSize-1) || c == 0 || c == (columnsSize-1)
}

type node struct {
	visited  bool
	toRemove bool
}

func initVisitedGrid(grid [][]string) [][]node {
	visited := make([][]node, len(grid))
	cz := len(grid[0])
	for i := range visited {
		visited[i] = make([]node, cz)
	}
	return visited
}
