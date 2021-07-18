package player

import (
	"fmt"
	card "racko/card"
	"strings"
)

type Hand struct {
	lowestCard *card.Card
	Size       int
}

// Adds a card to the Hand
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

// Loops through the Hand woring it's way back though
// the Linked List until it reaches the card at the
// correct index and returns it
func GetCardAtIndex(index int, c *card.Card) *card.Card {
	curCard := c
	i := 0

	for i <= index {
		if i == index {
			return curCard
		}
		curCard = curCard.Previous
		i++
	}

	return nil
}

// Calls GetCardAtIndex to retrieve the Card from the Hand
// Then swaps out the new card for the old in the LLinked List
func (h *Hand) SwapOutCard(c *card.Card, pos int) *card.Card {
	retCard := GetCardAtIndex(pos, h.lowestCard)

	switch {
	// Is top card
	case retCard.Next == nil:
		retCard.Previous.Next = c
		c.Next, c.Previous = nil, retCard.Previous
		h.lowestCard = c

	// Is bottom card
	case retCard.Previous == nil:
		retCard.Next.Previous = c
		c.Next, c.Previous = retCard.Next, nil

	default:
		retCard.Next.Previous = c
		retCard.Previous.Next = c
		c.Next, c.Previous = retCard.Next, retCard.Previous
	}
	retCard.Next, retCard.Previous = nil, nil

	return retCard
}

func (h *Hand) HasRacko() bool {
	curCard := h.lowestCard
	for i := 0; i < h.Size; i++ {
		if curCard.Previous == nil {
			return true
		} else if curCard.Value > curCard.Previous.Value {
			return false
		}
		curCard = curCard.Previous
	}
	return true
}

func (h *Hand) ShowHand() string {
	tempCard := h.lowestCard
	var retString string

	for i := 0; i < h.Size; i++ {
		retString = retString + fmt.Sprintf("%d:%s%d\n", i+1, strings.Repeat(" ", tempCard.Value), tempCard.Value)
		tempCard = tempCard.Previous
	}
	return retString
}
