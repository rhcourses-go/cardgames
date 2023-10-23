package cards

// Suit ist ein Typ für die vier Farben eines Kartenspiels.
type Suit int

// Enum für die vier Farben eines Kartenspiels.
const (
	Clubs Suit = iota
	Spades
	Hearts
	Diamonds
)

// Suits ist eine Liste aller Farben.
var Suits = []Suit{Clubs, Spades, Hearts, Diamonds}

// SuitStrings ist eine Map, die jeder Farbe einen String zuordnet.
// In diesem Fall sind das die Unicode-Zeichen für die vier Farben.
var SuitStrings = map[Suit]string{
	Clubs:    "♣",
	Spades:   "♠",
	Hearts:   "♥",
	Diamonds: "♦",
}

// String liefert die String-Repräsentation der Farbe.
func (s Suit) String() string {
	return SuitStrings[s]
}
