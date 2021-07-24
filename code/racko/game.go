package main

import (
	"fmt"
	"racko/card"
	"racko/player"
	"strconv"
	"strings"
)

type Game struct {
	Players           []player.Player
	NumPlayers        int
	CurrentPlayerTurn int
	CurrentRound      int
	Deck, Discard     *card.Deck
}

func (g *Game) PrintPrompt() {
	fmt.Println(strings.Repeat("#", 60))
	fmt.Printf("Round: %d\n", g.CurrentRound)
	fmt.Println(g.GetPlayer().TurnPrompt())
	fmt.Println(g.GetPlayer().ShowHand())
	fmt.Printf("Available card in discard pile: %d\n", g.Discard.Peek())
}

func PlayerPileOption() string {
	var input string
	fmt.Println("Enter 'p' to get the card from the discard pile, or 'd' to draw the unknown card from the top of the deck: ")
	fmt.Scanln(&input)
	return input
}

func PlayerIndexOption(val int) int {
	var input string
	fmt.Printf("Enter the slot number from the left edge of the display that you want to replace with %d. Less than 1 simply discards the card:\n", val)
	fmt.Scanln(&input)

	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Input")
		PlayerIndexOption(val)
	}

	if intInput < 11 {
		return intInput
	} else {
		fmt.Println("Invalid Input")
		PlayerIndexOption(val)
	}
	return 0
}

func (g *Game) DoNextTurn() {
	g.PrintPrompt()
	curPlayer := &g.Players[g.CurrentPlayerTurn]
	if curPlayer.IsComputer {

	} else {
		var index int
		switch PlayerPileOption() {
		case "p":
			card := g.Discard.Draw()
			index = PlayerIndexOption(card.Value)
			if index < 0 {
				g.Discard.Discard(card)
			} else {
				g.Discard.Discard(curPlayer.SwapOutCard(card, index))
			}
		case "d":
			card := g.DrawFromDeck()
			index = PlayerIndexOption(card.Value)
			if index < 0 {
				g.Discard.Discard(card)
			} else {
				g.Discard.Discard(curPlayer.SwapOutCard(card, index))
			}
		default:
			fmt.Println("Invalid Option")
			g.DoNextTurn()
		}
	}
	if g.CurrentPlayerTurn >= g.NumPlayers-1 {
		g.CurrentPlayerTurn = 0
	} else {
		g.CurrentPlayerTurn++
	}
	fmt.Printf(">>> G Current Player %d\n", g.CurrentPlayerTurn)
}

func (g *Game) IsGameOver() bool {
	return g.Players[g.CurrentPlayerTurn].HasRacko()
}

func (g *Game) GetPlayer() *player.Player {
	return &g.Players[g.CurrentPlayerTurn]
}

func (g *Game) ShowTopOfDiscardPile() int {
	return g.Deck.Peek()
}

func (g *Game) DrawFromDeck() *card.Card {
	retCard := g.Deck.Draw()
	if g.Deck.NumCardsInDeck < 1 {
		g.Deck = g.Discard
		g.Deck.Shuffle()
		g.Discard = &card.Deck{}
		g.Discard.Push(g.Deck.Draw().Value)
	}
	return retCard
}

func (g *Game) DiscardCard(c *card.Card) {
	g.Discard.Discard(c)
}

func NewGame(players []player.Player) *Game {
	g := Game{
		Players:           players,
		NumPlayers:        len(players),
		CurrentPlayerTurn: 0,
		CurrentRound:      1,
		Deck:              card.NewDeck(60),
		Discard:           &card.Deck{},
	}
	return &g
}
