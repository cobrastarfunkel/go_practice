package player

import (
	"fmt"
	card "racko/card"
)

type Player struct {
	PlayerName string
	PlayerHand Hand
	IsComputer bool
}

func (p *Player) TurnPrompt() string {
	return fmt.Sprintf("%s's turn\n", p.PlayerName)
}

func (p *Player) ShowHand() string {
	return p.PlayerHand.ShowHand()
}

func (p *Player) SwapOutCard(c *card.Card, pos int) *card.Card {
	return p.PlayerHand.SwapOutCard(c, pos)
}

func (p *Player) HasRacko() bool {
	return p.PlayerHand.HasRacko()
}

func (p *Player) MakeChoice(discardVal int) string {
	if discardVal == 1 {
		return "p"
	}

	tempCard := p.PlayerHand.lowestCard
	for i := 0; i < p.PlayerHand.Size; i++ {
		if tempCard.Previous == nil {
			return "d"
		}
		if discardVal > tempCard.Value && discardVal < tempCard.Previous.Value {
			return "p"
		}
	}
	return "d"
}

func (p *Player) ComputerChooseSlot(cardVal int) int {
	return 0
}
