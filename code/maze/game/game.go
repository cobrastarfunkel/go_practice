package game

import (
	"fmt"
	"math/rand"
	"maze/maze"
	"maze/passage"
	"maze/player"
	"maze/room"
	"time"
)

type Game struct {
	maze         *maze.Maze
	player       *player.Player
	currentRoom  *room.Room
	currentItems map[string]bool
}

func getPassageDescription(p *passage.Passage) string {
	var retStr string
	if !p.IsOpen && p.Key != "" {
		retStr = fmt.Sprintf("there is a locked door, it requires a %s.", p.Key)
	} else if p.IsOpen && p.Key != "" {
		retStr = "there is an open door"
	} else if p.IsOpen {
		retStr = "there is an open passage way."
	} else {
		retStr = "there is a wall."
	}
	return retStr
}

func (g *Game) DisplayPassages() {
	fmt.Printf("To the North %s\n", getPassageDescription(g.currentRoom.NorthPassage))
	fmt.Printf("To the East %s\n", getPassageDescription(g.currentRoom.EastPassage))
	fmt.Printf("To the South %s\n", getPassageDescription(g.currentRoom.SouthPassage))
	fmt.Printf("To the West %s\n", getPassageDescription(g.currentRoom.WestPassage))
}

func (g *Game) DisplayItems() {
	rand.Seed(time.Now().UnixNano())
	locations := []string{
		"wall",
		"floor",
		"ceiling",
		"shelf",
		"desk",
	}
	for _, item := range g.currentRoom.Items {
		fmt.Printf("There is a/an %s on the %s\n", item, locations[rand.Intn(len(locations))])
	}
}

func NewGame(m *maze.Maze, p *player.Player, r *room.Room) *Game {
	game := Game{
		maze:         m,
		player:       p,
		currentRoom:  r,
		currentItems: make(map[string]bool),
	}

	return &game
}
