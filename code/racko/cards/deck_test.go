package card

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
	deck.Shuffle()
	deck.Printdeck()
}

func TestInsertAt(t *testing.T) {
	length := 10

	var tests = []struct {
		testVal, index int
		card           Card
	}{
		{999, length, Card{Value: 999}},
		{222, 4, Card{Value: 222}},
		{333, 0, Card{Value: 333}},
	}
	deck := NewDeck(length)

	loopCards := func(index int, val int) {
		tempCard := deck.Top
		for i := 0; i <= index; i++ {
			if i == index && tempCard.Value != val {
				t.Errorf("InserAt %d wrong\nShould be %d\nIs: %d", index, val, tempCard.Value)
			}
			tempCard = tempCard.Previous
		}
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Val: %d, Index: %d", test.testVal, test.index)
		t.Run(testName, func(t *testing.T) {
			deck.InsertAt(&test.card, test.index)
			loopCards(test.index, test.testVal)
		})
	}
}
