package player

import (
	"testing"
)

func TestLoadMaze(t *testing.T) {
	player := NewPlayer()

	if player.GetMoves() != 0 {
		t.Errorf("Player moves should be 0 Player Moves: %d", player.GetMoves())
	}

	player.IncrementMoves()
	if player.GetMoves() != 1 {
		t.Errorf("Player moves should be 1 Player Moves: %d", player.GetMoves())
	}

	player.SetPosition(2, 4)
	row, col := player.GetPosition()
	if row != 2 || col != 4 {
		t.Errorf("Player position set wrong Row: %d Col %d", row, col)
	}

	testItem := "RedKey"
	player.AddItem(testItem)
	if !player.HasItem(testItem) {
		t.Error("Item not added to Player inventory")
	}

	player.UseItem(testItem)
	if player.HasItem(testItem) {
		t.Error("Player Item not Removed")
	}
}
