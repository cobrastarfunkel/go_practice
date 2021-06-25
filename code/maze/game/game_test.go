package game

import (
	"maze/maze"
	"maze/player"
	"testing"
)

func TestGame(t *testing.T) {
	test_file := "../maze/test_maze.txt"
	good_maze := &maze.Maze{MazeFile: test_file}
	good_maze.LoadMaze()
	play := player.NewPlayer()

	row, col := play.GetPosition()
	r := &good_maze.Rooms[row][col]
	g := NewGame(good_maze, play, r)
	g.DisplayPassages()
	g.DisplayItems()

	g.DisplayItems()
	g.MoveDirection('s')
	g.MoveDirection('e')
	g.DisplayPassages()
	g.DisplayItems()
	g.MoveDirection('e')
	g.DisplayPassages()
	g.DisplayItems()
	g.MoveDirection('w')
	g.DisplayPassages()
	g.DisplayItems()
	g.MoveDirection('w')
	g.DisplayPassages()
	g.DisplayItems()
	g.MoveDirection('s')
	g.DisplayPassages()
	g.DisplayItems()
	g.MoveDirection('E')
	g.DisplayPassages()
	g.DisplayItems()
	g.MoveDirection('E')
	g.DisplayPassages()
	g.DisplayItems()
}
