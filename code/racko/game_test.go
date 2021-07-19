package main

import (
	"fmt"
	"racko/card"
	"racko/player"
	"testing"
)

func BuildTestObj(deckLength int) *Game {
	handSize := 3
	hands := []*player.Hand{
		{Size: handSize},
		{Size: handSize},
	}

	for j, hand := range hands {
		for i := 1; i <= handSize; i++ {
			hand.AddToHand(&card.Card{Value: (i + j) * 2})
		}
	}

	player1 := player.Player{PlayerName: "Dave", IsComputer: false, PlayerHand: *hands[0]}
	player2 := player.Player{PlayerName: "Stan", IsComputer: false, PlayerHand: *hands[1]}
	game := Game{
		Players:           []player.Player{player1, player2},
		NumPlayers:        2,
		CurrentPlayerTurn: 1,
		CurrentRound:      1,
		Deck:              card.NewDeck(deckLength),
		Discard:           &card.Deck{},
	}
	return &game
}

func TestDrawFromDeck(t *testing.T) {
	deckLength := 5
	game := BuildTestObj(deckLength)
	game.Deck.PrintDeck()

	topCardVal := game.ShowTopOfDiscardPile()
	tempCard := game.DrawFromDeck()
	game.Discard.Discard(tempCard)

	if topCardVal != tempCard.Value {
		t.Errorf("Top card and Draw card do not match\nTop: %d\nDraw: %d", topCardVal, tempCard.Value)
	}

	nextDraw := game.ShowTopOfDiscardPile()

	if nextDraw == tempCard.Value {
		t.Errorf("Top card and Draw card match after 2nd Draw\nTop: %d\nDraw: %d", nextDraw, tempCard.Value)
	}

	for i := 0; i < deckLength*2; i++ {
		fmt.Println("\n>>> Deck")
		game.Deck.PrintDeck()
		fmt.Println("### Discard")
		game.Discard.PrintDeck()
		fmt.Println()
		game.DiscardCard(game.Players[1].SwapOutCard(game.DrawFromDeck(), 1))
	}
}
