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

var dr, dc = [4]int{-1, 1, 0, 0}, [4]int{0, 0, -1, 1}

func explore(grid [][]string, r int, c int, visited [][]bool) (size int) {
	if outOfGrid(r, len(grid), c, len(grid[0])) {
		return
	}
	if water == grid[r][c] {
		return
	}
	if v := visited[r][c]; v {
		return
	}
	visited[r][c] = true
	size++

	for i := 0; i < 4; i++ {
		rr := r + dr[i]
		cc := c + dc[i]
		size += explore(grid, rr, cc, visited)
	}

	return
}

func outOfGrid(r int, R int, c int, C int) bool {
	return r < 0 || r >= R || c < 0 || c >= C
}

func initVisitedGrid(grid [][]string) [][]bool {
	visited := make([][]bool, len(grid))
	cz := len(grid[0])
	for i := range visited {
		visited[i] = make([]bool, cz)
	}
	return visited
}
