package player

import (
	"fmt"
	card "racko/card"
	"testing"
)

func TestHand(t *testing.T) {
	var tests = []struct {
		testVal int
		c       *card.Card
	}{
		{11, &card.Card{Value: 11}},
		{22, &card.Card{Value: 22}},
		{33, &card.Card{Value: 33}},
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
		{1, &card.Card{Value: 12}},
		{2, &card.Card{Value: 45}},
		{0, &card.Card{Value: 39}},
	}

	for _, test := range swapTests {
		testname := fmt.Sprintf("position: %d\nCardVal: %v", test.position, test.c.Value)
		t.Run(testname, func(t *testing.T) {
			hand.SwapOutCard(test.c, test.position)
		})
	}

	t.Run("Should not have Racko after position swaps", func(t *testing.T) {
		if hand.HasRacko() {
			t.Errorf("Should not Have Racko")
		}
	})

	hand.SwapOutCard(&card.Card{Value: 1}, 0)
	hand.SwapOutCard(&card.Card{Value: 2}, 1)
	hand.SwapOutCard(&card.Card{Value: 3}, 2)

	t.Run("Should have Racko", func(t *testing.T) {
		if !hand.HasRacko() {
			t.Errorf("Should Have Racko")
		}
	})

	hand.SwapOutCard(&card.Card{Value: 4}, 0)
	t.Run("Should not have Racko 0 index highest", func(t *testing.T) {
		if hand.HasRacko() {
			t.Errorf("Should not Have Racko")
		}
	})

	hand.SwapOutCard(&card.Card{Value: 3}, 2)
	hand.SwapOutCard(&card.Card{Value: 5}, 1)
	t.Run("Should not have Racko middle index highest", func(t *testing.T) {
		if hand.HasRacko() {
			t.Errorf("Should not Have Racko")
		}
	})
	fmt.Println(hand.ShowHand())
}

func TestComputerPlayer(t *testing.T) {
	for i := 0; i < 60; i++ {

	}
}
