package main

import (
	"racko/card"
	"racko/player"
	"testing"
)

func BuildTestObj() *Game {
	deckLength := 10
	handSize := 3
	hands := []*player.Hand{
		{Size: handSize},
		{Size: handSize},
	}

	for j, hand := range hands {
		for i := 1; i <= handSize; i++ {
			hand.AddToHand(&card.Card{Value: i + j})
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
	game := BuildTestObj()
	topCardVal := game.ShowTopOfDiscardPile()
	tempCard := game.DrawFromDeck().Value

	if topCardVal != tempCard {
		t.Errorf("Top card and Draw card do not match\nTop: %d\nDraw: %d", topCardVal, tempCard)
	}

	nextDraw := game.ShowTopOfDiscardPile()

	if nextDraw == tempCard {
		t.Errorf("Top card and Draw card match after 2nd Draw\nTop: %d\nDraw: %d", nextDraw, tempCard)
	}
}
