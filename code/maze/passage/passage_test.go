package passage

import (
	"testing"
)

func TestRequiresKey(t *testing.T) {
	pKey := Passage{IsOpen: false, Key: "Blue"}
	pNoKey := Passage{IsOpen: false, Key: ""}

	if !pKey.RequiresKey() {
		t.Errorf("RequiresKey() failed, %v Requires a key", pKey)
	} else {
		t.Logf("%v Passed, Requires a key", pKey)
	}

	if pNoKey.RequiresKey() {
		t.Errorf("RequiresKey() Failed %v Does not Require a Key", pNoKey)
	} else {
		t.Logf("%v Passed, Requires no key", pNoKey)
	}
}

func TestOpen(t *testing.T) {
	pKey := Passage{IsOpen: false, Key: "Blue"}
	pNoKey := Passage{IsOpen: false, Key: ""}

	pKey.Open("red")
	if pKey.IsOpen {
		t.Errorf("Open failed with red, key == %s, Should Not be Open", pKey.Key)
	} else {
		t.Logf("%v Passed Open() test with bad key", pKey)
	}

	pKey.Open("Blue")
	if !pKey.IsOpen {
		t.Errorf("Open failed with Blue, key == %s, should be Open", pKey.Key)
	} else {
		t.Logf("%v Passed Open() test with good key", pKey)
	}

	pNoKey.Open("")
	if !pNoKey.IsOpen {
		t.Error("Open failed for Passage with no key requirement")
	} else {
		t.Log("Open passed for Passage with no key requirement")
	}

}
