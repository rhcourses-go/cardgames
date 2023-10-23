package cards

import "fmt"

// Card ist ein Typ für eine einzelne Spielkarte.
type Card struct {
	Rank
	Suit
}

// New gibt eine neue Spielkarte zurück.
func New(r Rank, s Suit) Card {
	return Card{r, s}
}

// String gibt den Namen der Spielkarte zurück.
func (c Card) String() string {
	return fmt.Sprintf("(%s %s)", c.Suit, c.Rank)
}

// Image gibt ein Bild der Spielkarte zurück.
func (c Card) Image() string {
	template := `
┌─────────┐
│%s       │
│%s        │
│         │
│    %s   │
│         │
│        %s│
│       %s│
└─────────┘`
	ranktop := RankTop(c.Rank)
	rankbottom := RankBottom(c.Rank)
	suit := c.Suit.String()
	return fmt.Sprintf(template, ranktop, suit, ranktop, suit, rankbottom)
}

// RankTop erwartet einen Rank und gibt ihn zurück,
// wie er oben und in der Mitte der Karte angezeigt werden soll.
// Für alle Karten außer der 10 ist das der Rank,
// gefolgt von einem Leerzeichen.
// Für die 10 ist das der Rank ohne Leerzeichen.
func RankTop(r Rank) string {
	if r == Ten {
		return r.String()
	}
	return fmt.Sprintf("%s ", r)
}

// RankBottom erwartet einen Rank und gibt ihn zurück,
// wie er unten in der Karte angezeigt werden soll.
// Für alle Karten außer der 10 ist das der Rank
// mit einem Leerzeichen davor.
// Für die 10 ist das der Rank ohne Leerzeichen.
func RankBottom(r Rank) string {
	if r == Ten {
		return r.String()
	}
	return fmt.Sprintf(" %s", r)
}
