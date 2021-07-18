package main

import (
	"racko/card"
	"racko/player"
)

type Game struct {
	Players           []player.Player
	NumPlayers        int
	CurrentPlayerTurn int
	CurrentRound      int
	Deck, Discard     *card.Deck
}

func (g *Game) DoNextTurn() {

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

func (g *Game) NewGame() {
	g.CurrentPlayerTurn = 1
	g.CurrentRound = 1
	g.Deck, g.Discard = card.NewDeck(60), &card.Deck{}
}
