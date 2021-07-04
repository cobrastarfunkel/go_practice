package player

import (
	"fmt"
	card "racko/cards"
)

type Hand struct {
	lowestCard *card.Card
	Size       int
}

func (h *Hand) AddToHand(card *card.Card) {
	tempCard := h.lowestCard
	if h.lowestCard == nil {
		h.lowestCard = card
	} else {
		card.Previous, card.Next = tempCard, nil
		tempCard.Next = card
		h.lowestCard = card
	}
}

func (h *Hand) SwapOutCard(c *card.Card, pos int) *card.Card {
	var retCard *card.Card
	tempCard := h.lowestCard
	i := 0

	for i <= pos {
		if i == pos {
			retCard = tempCard
			switch {
			case tempCard.Next == nil:
				tempCard.Previous.Next = c
				c.Next, c.Previous = nil, tempCard.Previous
				h.lowestCard = c

			case tempCard.Previous == nil:
				tempCard.Next.Previous = c
				c.Next, c.Previous = tempCard.Next, nil

			default:
				tempCard.Next.Previous = c
				tempCard.Previous.Next = c
				c.Next, c.Previous = tempCard.Next, tempCard.Previous
			}
			retCard.Next, retCard.Previous = nil, nil
		} else {
			tempCard = tempCard.Previous
		}
		i++
	}
	return retCard
}

func (h *Hand) ShowHand() {
	tempCard := h.lowestCard
	for i := 0; i < h.Size; i++ {
		fmt.Printf("Pos: %d\nVal: %d\n\n", i, tempCard.Value)
		tempCard = tempCard.Previous
	}
}
