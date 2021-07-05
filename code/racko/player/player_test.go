package player

import (
	"fmt"
	card "racko/cards"
	"testing"
)

func TestHand(t *testing.T) {
	var tests = []struct {
		testVal int
		c       *card.Card
	}{
		{111, &card.Card{Value: 111}},
		{222, &card.Card{Value: 222}},
		{333, &card.Card{Value: 333}},
	}
	hand := Hand{Size: 3}

	for _, test := range tests {
		testname := fmt.Sprintf("AddToHand() Val: %d", test.testVal)
		t.Run(testname, func(t *testing.T) {
			hand.AddToHand(test.c)
			if hand.lowestCard.Value != test.testVal {
				t.Errorf("LowestCard val bad %d\nVal: %d", hand.lowestCard.Value, test.testVal)
			}
		})
	}
	hand.ShowHand()

	var swapTests = []struct {
		position int
		c        *card.Card
	}{
		{1, &card.Card{Value: 123}},
		{2, &card.Card{Value: 456}},
		{0, &card.Card{Value: 789}},
	}

	for _, test := range swapTests {
		testname := fmt.Sprintf("position: %d\nCardVal: %v", test.position, test.c.Value)
		t.Run(testname, func(t *testing.T) {
			hand.SwapOutCard(test.c, test.position)
		})
	}
	hand.ShowHand()
}
