package main

import (
	"fmt"
	"racko/player"
	"strconv"
	"strings"
)

func GetNumPlayers() int {
	var numPlayers string
	fmt.Print("How many Players(Including Computer) Max 4: ")
	fmt.Scanln(&numPlayers)

	intPlayers, err := strconv.Atoi(numPlayers)
	if err != nil {
		fmt.Println("Invalid Input")
		return GetNumPlayers()
	}

	if intPlayers < 5 && intPlayers > 1 {
		return intPlayers
	} else {
		fmt.Println("Invalid Input")
		return GetNumPlayers()
	}
}

func GetPlayerName(index int) string {
	var name string
	fmt.Printf("Enter player %d's name: ", index)
	fmt.Scanln(&name)
	return name
}

func GetIsComputer() bool {
	var isComp string
	fmt.Println("Is Player a computer (y/n): ")
	fmt.Scanln(&isComp)
	switch strings.ToLower(isComp) {
	case "y":
		return true
	case "n":
		return false
	default:
		fmt.Println("Invalid Entry")
		return GetIsComputer()
	}
}

func GetPlayerInfo() []player.Player {
	var players []player.Player
	numPlayers := GetNumPlayers()

	for i := 0; i < numPlayers; i++ {
		players = append(players, player.Player{
			PlayerName: GetPlayerName(i),
			IsComputer: GetIsComputer(),
			PlayerHand: player.Hand{Size: 10},
		})
	}
	return players
}

func MakeHands(g *Game, p *player.Player) {
	for i := 0; i < 10; i++ {
		p.PlayerHand.AddToHand(g.DrawFromDeck())
	}
}

func start(g *Game) {
	fmt.Println("Welcome to Racko!")
	g.Deck.Shuffle()
	for i := range g.Players {
		MakeHands(g, &g.Players[i])
	}
	g.Discard.Discard(g.DrawFromDeck())

	for {
		if g.IsGameOver() {
			break
		}
		// Player hand isn't keeping 10 cards after p and index selection
		g.DoNextTurn()
	}
	fmt.Printf("Congratulations %s, you won!", g.GetPlayer().PlayerName)
}

func createGame() {
	players := GetPlayerInfo()
	game := NewGame(players)
	start(game)
}

func main() {
	createGame()
}
