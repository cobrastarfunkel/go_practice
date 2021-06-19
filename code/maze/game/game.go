package game

import (
	"fmt"
	"math/rand"
	"maze/maze"
	"maze/passage"
	"maze/player"
	"maze/room"
	"time"
	"unicode"
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

func (g *Game) ExitFound() bool {
	if g.getCurrentRoomName() == "The Exit" {
		return true
	}
	return false
}

func (g *Game) DisplayPassages() {
	fmt.Printf("\n\nYou are now in the %s\n", g.getCurrentRoomName())
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
		itemLoc := locations[rand.Intn(len(locations))]
		fmt.Printf("There is a/an %s on the %s\n", item, itemLoc)

		g.player.AddItem(g.currentRoom.AquireNextItem())

		fmt.Printf("You grab the %s from the %s\n\n", item, itemLoc)
	}
}

func (g *Game) checkPassageReqs(p *passage.Passage) bool {
	if p.IsOpen {
		return true
	} else if p.RequiresKey() {
		if g.player.HasItem(p.Key) {
			p.Open(g.player.UseItem(p.Key))
			fmt.Printf("You unlock the door using the %s\n", p.Key)

			return true
		} else {
			fmt.Printf("This door requires the %s\n", p.Key)
		}
	} else {
		fmt.Println("You can't go that way.")
	}
	return false
}

func (g *Game) validDirection(direction rune) bool {
	switch direction {
	case 'N':
		return g.checkPassageReqs(g.currentRoom.NorthPassage)
	case 'E':
		return g.checkPassageReqs(g.currentRoom.EastPassage)
	case 'S':
		return g.checkPassageReqs(g.currentRoom.SouthPassage)
	case 'W':
		return g.checkPassageReqs(g.currentRoom.WestPassage)
	default:
		return false
	}
}

func (g *Game) MoveDirection(direction rune) {
	direction = unicode.ToUpper(direction)
	directions := map[rune]string{
		'N': "North",
		'E': "East",
		'S': "South",
		'W': "West",
	}
	fmt.Printf("You try to go %s\n", directions[direction])

	if g.validDirection(direction) {
		curRow, curCol := g.player.GetPosition()
		switch direction {
		case 'N':
			g.player.SetPosition(curRow-1, curCol)
			g.currentRoom = &g.maze.Rooms[curRow-1][curCol]
			nextPass := g.maze.Rooms[curRow-1][curCol].SouthPassage
			nextPass.Open(nextPass.Key)
		case 'E':
			g.player.SetPosition(curRow, curCol+1)
			g.currentRoom = &g.maze.Rooms[curRow][curCol+1]
			nextPass := g.maze.Rooms[curRow][curCol+1].WestPassage
			nextPass.Open(nextPass.Key)
		case 'S':
			g.player.SetPosition(curRow+1, curCol)
			g.currentRoom = &g.maze.Rooms[curRow+1][curCol]
			nextPass := g.maze.Rooms[curRow+1][curCol].NorthPassage
			nextPass.Open(nextPass.Key)
		case 'W':
			g.player.SetPosition(curRow, curCol-1)
			g.currentRoom = &g.maze.Rooms[curRow][curCol-1]
			nextPass := g.maze.Rooms[curRow][curCol-1].EastPassage
			nextPass.Open(nextPass.Key)
		}
	}
}

func (g *Game) getCurrentRoomName() string {
	return g.currentRoom.Name
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
