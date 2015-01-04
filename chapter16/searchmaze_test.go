package chapter16

import (
	"fmt"
	"testing"
)

func TestSearchMaze(t *testing.T) {
	maze := [][]int{
		{0, 1, 1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 0, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 1, 0, 0},
		{1, 1, 1, 0, 0, 0, 1, 1, 0, 1},
		{1, 0, 0, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 1, 0, 1, 0, 0, 1},
		{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 1},
	}

	fmt.Println("TestSearchMaze:", SearchMaze(maze, Point{9, 0}, Point{0, 9}))
	fmt.Println()
}

func TestSearchMazeForUnreachableEndPoint(t *testing.T) {
	maze := [][]int{
		{0, 1, 1, 1, 1, 1, 0, 0, 0, 1},
		{1, 1, 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 1, 1, 0, 0, 1, 0, 0},
		{1, 1, 1, 0, 0, 0, 1, 1, 0, 1},
		{1, 0, 0, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 1, 0, 1, 0, 0, 1},
		{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 1},
	}

	path := SearchMaze(maze, Point{9, 0}, Point{0, 9})

	if len(path) > 0 {
		t.Fatalf("TestSearchMazeForUnreachableEndPoint: want empty path, got %v", path)
	}
}
