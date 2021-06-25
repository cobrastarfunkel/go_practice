package main

import (
	"bufio"
	"fmt"
	"maze/game"
	"maze/maze"
	"maze/player"
	"os"
	"strings"
)

func makeGame() game.Game {
	test_file := "./maze/test_maze.txt"
	good_maze := &maze.Maze{MazeFile: test_file}
	good_maze.LoadMaze()
	play := player.NewPlayer()
	r := &good_maze.Rooms[0][0]
	g := game.NewGame(good_maze, play, r)
	return *g
}

func run() {
	gam := makeGame()

	for {
		reader := bufio.NewReader(os.Stdin)

		gam.DisplayPassages()
		gam.DisplayItems()
		fmt.Println("Enter a Direction [N, E, S, W]: ")
		char, _, err := reader.ReadRune()

		if err != nil {
			panic(err)
		}
		gam.MoveDirection(char)
	}
}

func start() {
	var input string
	for strings.ToLower(input) != "n" {
		fmt.Println("Play Maze game [y/n]?")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "y" {
			run()
		}
	}
}

func main() {
	start()
}
