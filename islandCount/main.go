package main

import "fmt"

const (
	water = "w"
	land  = "l"
)

func main() {
	grid := [][]string{{water, land, water, water}}
	fmt.Println(islandCount(grid))
}

func islandCount(grid [][]string) (counter int) {
	visited := make([][]bool, len(grid))
	cz := len(grid[0])
	for i := range visited {
		visited[i] = make([]bool, cz)
	}

	for r := range grid {
		for c := range grid[r] {
			if true == explore(grid, r, c, visited) {
				counter++
			}
		}
	}

	return
}

var dr, dc = [4]int{-1, 1, 0, 0}, [4]int{0, 0, -1, 1}

func explore(grid [][]string, r int, c int, visited [][]bool) (foundLend bool) {
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

	for i := 0; i < 4; i++ {
		rr := r + dr[i]
		cc := c + dc[i]
		explore(grid, rr, cc, visited)
	}

	return true
}

func outOfGrid(r int, R int, c int, C int) bool {
	return r < 0 || r >= R || c < 0 || c >= C
}

func dfs(graph map[string][]string, root string) []string {
	nodes := []string{root}
	for _, v := range graph[root] {
		nodes = append(nodes, dfs(graph, v)...)
	}
	return nodes
}
