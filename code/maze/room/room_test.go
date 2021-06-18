package room

import (
	"maze/passage"
	"testing"
)

func TestAddItem(t *testing.T) {
	r := Room{
		Name:         "Test Room",
		Items:        []string{},
		NorthPassage: &passage.Passage{IsOpen: false, Key: "Blue"},
		EastPassage:  &passage.Passage{IsOpen: false, Key: ""},
		SouthPassage: &passage.Passage{IsOpen: false, Key: "Red"},
		WestPassage:  &passage.Passage{IsOpen: true, Key: ""},
	}

	testItem := "blueKey"
	r.AddItem(testItem)

	tempItem := r.AquireNextItem()
	if tempItem != testItem {
		t.Errorf("AddItem() failed %s not found", testItem)
	} else {
		t.Log("AddItem Passed")
	}

	if r.AquireNextItem() != "" {
		t.Error("AquireNextItem() Failed")
	} else {
		t.Log("AquireNextItem() Passed")
	}

}
