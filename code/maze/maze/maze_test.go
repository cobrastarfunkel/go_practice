package maze

import (
	"testing"
)

func TestLoadMaze(t *testing.T) {
	test_file := "./test_maze.txt"
	good_maze := Maze{MazeFile: test_file}

	// bad_file := "./bad_maze.txt"
	// bad_maze := Maze{MazeFile: bad_file}

	good_maze.LoadMaze()
	if good_maze.cols != 3 || good_maze.rows != 2 {
		t.Errorf("Load maze failed to get size Row: %d Col: %d", good_maze.rows, good_maze.cols)
	} else {
		t.Log("LoadMaze size test passed")
	}

	//bad_maze.LoadMaze()
}
