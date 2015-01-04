// Problem 16.1

package chapter16

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

func SearchMaze(maze [][]int, start, end Point) []Point {
	path := make([]Point, 0)

	isFeasible := func(point Point) bool {
		return point.y >= 0 && point.y < len(maze) &&
			point.x >= 0 && point.x < len(maze[0]) &&
			maze[point.y][point.x] == 1
	}

	var dfs func(Point) bool
	dfs = func(point Point) bool {
		if point == end {
			return true
		}

		for _, e := range edges {
			next := Point{point.y + e.dy, point.x + e.dx}
			if isFeasible(next) {
				maze[next.y][next.x] = 0
				path = append(path, next)
				if dfs(next) {
					return true
				}
				path = path[:len(path)-1]
			}
		}

		return false
	}

	maze[start.y][start.x] = 0
	path = append(path, start)
	if !dfs(start) {
		path = path[:len(path)-1]
	}

	return path
}
