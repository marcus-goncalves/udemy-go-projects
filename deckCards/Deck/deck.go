package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func NewDeck() deck {
	newDeck := deck{}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) Deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) Shuffle() deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for idx := range d {
		newPos := r.Intn(len(d) - 1)

		d[idx], d[newPos] = d[newPos], d[idx]
	}

	return d
}

func (d deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("deu ruim: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return s
}
