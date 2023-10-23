package cards

// Rank ist ein Typ für die Werte einer Karte.
type Rank int

// Enum für die Werte einer Karte.
const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Ranks ist eine Liste aller Werte.
var Ranks = []Rank{
	Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace,
}

// RankStrings ist eine Map, die jedem Wert einen String zuordnet.
// Beachte, dass die 10 breiter ist als die anderen Werte.
// Dies muss später bei der Darstellung von Karten berücksichtigt werden.
var RankStrings = map[Rank]string{
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
	Ace:   "A",
}

// String liefert die String-Repräsentation des Werts.
func (r Rank) String() string {
	return RankStrings[r]
}
