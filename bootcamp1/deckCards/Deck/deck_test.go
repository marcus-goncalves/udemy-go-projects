package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("expected 52 cards. got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected Ace of Spades. got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("expected King of Clubs. got %v", d[len(d)-1])
	}
}

func TestSaveFileAndLoadDeck(t *testing.T) {
	file := "./_decktest"

	os.Remove(file)

	deck := NewDeck()
	deck.SaveToFile(file)

	loadedDeck := NewDeckFromFile(file)
	if len(loadedDeck) != 52 {
		t.Errorf("expected 52 cards. got %v", len(loadedDeck))
	}

	os.Remove(file)
}
