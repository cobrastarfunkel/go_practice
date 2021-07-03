package main

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

	for i := 0; i < counter; i++ {
		max := maxCount
		rando := rand.Intn(max)

		if d.Top == nil {
			d.Top = list[rando]
			d.NumCardsInDeck++
		} else {
			newTop := Card{Previous: d.Top, Value: list[rando].Value}
			d.Top.Next = list[rando]
			d.Top = &newTop
			d.NumCardsInDeck++
		}

		list = removeElement(list, rando)
		maxCount--
	}
}

func (d *deck) Printdeck() {
	list := d.Top
	for i := 1; i <= d.NumCardsInDeck; i++ {
		fmt.Printf("Card %d\n", list.Value)
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
