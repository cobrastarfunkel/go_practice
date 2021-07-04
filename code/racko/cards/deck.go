package card

import (
	"fmt"
	"math/rand"
	"time"
)

type deck struct {
	Top                     *Card
	NUM_CARDS_IN_RACKO_DECK int
	NumCardsInDeck          int
}

func (d *deck) MakeFullDeck() {
	for i := 1; i <= d.NUM_CARDS_IN_RACKO_DECK; i++ {
		d.Push(i)
	}
	d.Shuffle()
}

func (d *deck) Push(val int) {
	if d.Top == nil {
		d.Top = &Card{Next: nil, Previous: nil, Value: val}
	} else {
		// New Top Card
		nextTop := Card{Next: nil, Previous: d.Top, Value: val}

		// Prev Top card Next should refernce new Top
		d.Top.Next = &nextTop

		d.Top = &nextTop
	}
	d.NumCardsInDeck++
}

func (d *deck) Discard(c *Card) {
	if d.Top == nil {
		d.Top = c
	} else {
		d.Top.Next = c
		c.Previous = d.Top
		d.Top = c
	}
	d.NumCardsInDeck++
}

func (d *deck) Draw() *Card {
	retCard := d.Top
	if d.Top.Previous != nil {
		d.Top = retCard.Previous
	}
	retCard.Previous = nil
	retCard.Next = nil
	d.Top.Next = nil

	d.NumCardsInDeck--
	return retCard
}

func (d *deck) Peek() int {
	return d.Top.Value
}

func removeElement(list []*Card, i int) []*Card {
	list[i] = list[len(list)-1]
	list[len(list)-1] = nil
	return list[:len(list)-1]
}

func (d *deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	var list []*Card
	origDeckSize := d.NumCardsInDeck

	for i := 0; i < origDeckSize; i++ {
		list = append(list, d.Draw())
	}

	maxCount := len(list)
	counter := len(list)

	d.Top = nil

	for i := 0; i < counter; i++ {
		max := maxCount
		rando := rand.Intn(max)

		if d.Top == nil {
			d.Top = &Card{Value: list[rando].Value, Previous: nil, Next: nil}
			d.NumCardsInDeck++
		} else {
			newTop := Card{Previous: d.Top, Value: list[rando].Value, Next: nil}
			d.Top.Next = &newTop
			d.Top = &newTop
			d.NumCardsInDeck++
		}

		list = removeElement(list, rando)
		maxCount--
	}
}

func (d *deck) InsertAt(c *Card, index int) {
	if index < 0 || index > d.NumCardsInDeck {
		return
	}

	tempCard := d.Top
	lastCard := false
	i := 0

	for i <= index {
		if i == index {
			if lastCard {
				tempCard.Previous = c
				c.Next = tempCard
			} else if tempCard.Next != nil {
				c.Next = tempCard.Next
				tempCard.Next.Previous = c
				c.Previous = tempCard
				tempCard.Next = c
			} else {
				c.Previous = tempCard
				tempCard.Next = c
				d.Top = c
			}
			d.NumCardsInDeck++
		}
		if tempCard.Previous != nil {
			tempCard = tempCard.Previous
		} else {
			lastCard = true
		}

		i++
	}
}

func (d *deck) Printdeck() {
	list := d.Top
	fmt.Printf("Num Cards %d\n", d.NumCardsInDeck)
	for i := 1; i <= d.NumCardsInDeck; i++ {
		fmt.Printf("Card %v\n", list)
		list = list.Previous
	}
}

func NewDeck(length int) *deck {
	deck := deck{
		NUM_CARDS_IN_RACKO_DECK: length,
	}
	deck.MakeFullDeck()
	return &deck
}
