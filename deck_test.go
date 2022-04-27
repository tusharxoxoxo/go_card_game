package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newdeck()

	if len(d) != 52 {
		t.Errorf("Expected output was 522 but we get %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}
}

func SaveTofileTest(t *testing.T) {
	os.Remove("_decktesting")

	d := newdeck()
	d.saveTofile("_decktesting")

	loadedDeck := readFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected length was 52 but we got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
