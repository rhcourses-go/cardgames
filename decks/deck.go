package decks

import "github.com/rhcourses-go/cardgames/cards"

// Deck ist ein Typ für einen Kartenstapel.
type Deck struct {
	cards []cards.Card
}

// New32 gibt einen neuen Kartenstapel zurück.
// Der Kartenstapel ist ein Skat-Blatt mit 32 Karten.
func New32() Deck {
	var deck Deck
	for _, suit := range cards.Suits {
		for rank := cards.Seven; rank <= cards.King; rank++ {
			deck.cards = append(deck.cards, cards.New(rank, suit))
		}
		deck.cards = append(deck.cards, cards.New(cards.Ace, suit))
	}
	return deck
}

// New52 gibt einen neuen Kartenstapel zurück.
// Der Kartenstapel ist ein Poker-Blatt mit 52 Karten.
func New52() Deck {
	var deck Deck
	for _, suit := range cards.Suits {
		for _, rank := range cards.Ranks {
			deck.cards = append(deck.cards, cards.New(rank, suit))
		}
	}
	return deck
}
