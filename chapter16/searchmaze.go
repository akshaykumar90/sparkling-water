// problem 16.1

package chapter16

import "fmt"

const (
	UNDISCOVERED = iota
	DISCOVERED
	VISITED
)

var edges = []struct {
	dy, dx int
}{
	{-1, 0},
	{0, +1},
	{+1, 0},
	{0, -1},
}

type Point struct {
	y, x int
}

func dfs(maze [][]int, state [][]int, start, end Point) bool {
	state[start.y][start.x] = DISCOVERED

	if start.x == end.x && start.y == end.y {
		fmt.Println(start.y, start.x)
		return true
	}

	for _, e := range edges {
		next_y, next_x := start.y+e.dy, start.x+e.dx
		if next_y >= 0 && next_y < len(maze) &&
			next_x >= 0 && next_x < len(maze[0]) &&
			maze[next_y][next_x] == 1 &&
			state[next_y][next_x] == UNDISCOVERED {
			if found := dfs(maze, state, Point{next_y, next_x}, end); found {
				fmt.Println(start.y, start.x)
				return true
			}
		}
	}

	state[start.y][start.x] = VISITED

	return false
}

func SearchMaze(maze [][]int, start, end Point) {
	state := make([][]int, len(maze))
	for i := 0; i < len(maze); i++ {
		state[i] = make([]int, len(maze[0]))
	}

	if found := dfs(maze, state, start, end); !found {
		fmt.Println("No path exists.")
	}

}
