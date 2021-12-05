package main

import (
	"fmt"
)

func main() {

	var maze = [][]string{
		{start, pass, pass, wall, pass, pass, pass},
		{pass, wall, pass, pass, pass, wall, pass},
		{pass, wall, pass, pass, pass, pass, pass},
		{pass, pass, wall, wall, pass, pass, pass},
		{wall, pass, wall, end, pass, pass, pass},
	}

	fmt.Println(findExitSteps(maze))
}

const (
	start = "S"
	end   = "E"
	pass  = "."
	wall  = "#"
)

type queue struct {
	list []int
}

func (q *queue) enqueue(i int) {
	q.list = append(q.list, i)
}

func (q *queue) dequeue() int {
	i := q.list[0]
	q.list = q.list[1:]
	return i
}

func (q *queue) len() int {
	return len(q.list)
}

// Empty queues
var rq, cq = queue{}, queue{}

func findExitSteps(maze [][]string) int {
	if 0 == len(maze) {
		panic("nothing to explore...")
	}

	R := len(maze)
	C := len(maze[0])

	nodesLeft := 1
	nodesNext := 0

	visited := [][]bool{}
	for i := 0; i < R; i++ {
		visited = append(visited, make([]bool, C))
	}

	reachedEnd := false
	counter := 0

	// starting point position (should be detected when parsing maze to matrix)
	sr, sc := 0, 0

	rq.enqueue(sr)
	cq.enqueue(sc)
	visited[sr][sc] = true

	for rq.len() > 0 {
		r := rq.dequeue()
		c := cq.dequeue()

		if end == maze[r][c] {
			reachedEnd = true
			break
		}
		nodesNext = exploreNeighbours(r, c, maze, visited, nodesNext)
		nodesLeft--
		if 0 == nodesLeft {
			nodesLeft = nodesNext
			nodesNext = 0
			counter++
		}
	}
	rq.list = []int{}
	cq.list = []int{}
	if reachedEnd {
		return counter
	}
	return -1
}

// Direction vectors
var dr, dc = []int{-1, 1, 0, 0}, []int{0, 0, 1, -1}

func exploreNeighbours(r int, c int, maze [][]string, visited [][]bool, nodesNext int) int {
	R := len(maze)
	C := len(maze[0])
	for i := 0; i < 4; i++ {
		rr := r + dr[i]
		cc := c + dc[i]

		if outOfMaze(rr, R, cc, C) {
			continue
		}
		if visited[rr][cc] || maze[rr][cc] == wall {
			continue
		}

		rq.enqueue(rr)
		cq.enqueue(cc)
		visited[rr][cc] = true
		nodesNext++
	}
	return nodesNext
}

func outOfMaze(r int, R int, c int, C int) bool {
	return r < 0 || r >= R || c < 0 || c >= C
}
