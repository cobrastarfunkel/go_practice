package main

import (
	"fmt"
	"testing"
)

func TestBuildDeck(t *testing.T) {
	length := 10
	deck := NewDeck(length)
	var list []int

	curTop := deck.Top
	for i := 1; i <= length; i++ {
		fmt.Printf("Curtop %d\n", curTop.Value)
		list = append(list, curTop.Value)
		curTop = curTop.Previous
	}

	if len(list) != length {
		t.Errorf("Length of Deck list %d != Length of Deck %d", len(list), length)
	}
	if deck.NumCardsInDeck != length {
		t.Errorf("NumCardsInDeck: %d != Length: %d", deck.NumCardsInDeck, length)
	}
}

func TestDiscard(t *testing.T) {
	length := 10
	testVal := 22
	discard := NewDeck(length)
	discard.Discard(&Card{Value: testVal})

	if discard.NumCardsInDeck != length+1 {
		t.Errorf("NumCardsInDeck: %d != Length: %d", discard.NumCardsInDeck, length+1)
	}
	if discard.Peek() != testVal {
		t.Errorf("Wrong val for Top in discard %d, Peek() Fail", discard.Top.Value)
	}
	if discard.Top.Next != nil {
		t.Errorf("Next for Top in Discard should == nil\nActual: %v", discard.Top.Next)
	}
	if discard.Top.Previous == nil {
		t.Error("Discard Previous Card should not == nil")
	}

	//discard.Printdeck()
}

func TestDraw(t *testing.T) {
	length := 10
	testVal := 22
	deck := NewDeck(length)
	deck.Discard(&Card{Value: testVal})

	card := deck.Draw()

	if card.Value != testVal {
		t.Errorf("Card %d != testVal %d", card.Value, testVal)
	}
	if deck.Top.Next != nil {
		t.Error("Next Card should be Nil for new Top after Draw")
	}
	if deck.NumCardsInDeck != length {
		t.Errorf("NumCardsInDeck should = length (%d) but = %d", length, deck.NumCardsInDeck)
	}
	//deck.Printdeck()
	deck.Shuffle()
	deck.Printdeck()
}

func TestInsertAt(t *testing.T) {
	length := 10
	testVal := 999
	deck := NewDeck(length)
	card := Card{Value: testVal}
	card2 := Card{Value: 333}
	card3 := Card{Value: 444}
	//deck.Printdeck()
	deck.InsertAt(&card, length)

	deck.Printdeck()

	deck.InsertAt(&card2, 5)

	deck.Printdeck()

	deck.InsertAt(&card3, 0)

	deck.Printdeck()
}
