package game

import (
	"maze/maze"
	"maze/player"
	"testing"
)

func TestAddItem(t *testing.T) {
	test_file := "../maze/test_maze.txt"
	good_maze := maze.Maze{MazeFile: test_file}
	good_maze.LoadMaze()
	play := *player.NewPlayer()

	row, col := play.GetPosition()
	r := good_maze.Rooms[row][col]
	g := NewGame(&good_maze, &play, &r)
	g.DisplayPassages()
	g.DisplayItems()
}
