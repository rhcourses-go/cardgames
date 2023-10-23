package cards

// RanksMatch erwarted zwei Karten und liefert true,
// wenn die Ränge gleich sind.
func RanksMatch(c1, c2 Card) bool {
	return c1.Rank == c2.Rank
}

// SuitsMatch erwarted zwei Karten und liefert true,
// wenn die Farben gleich sind.
func SuitsMatch(c1, c2 Card) bool {
	return c1.Suit == c2.Suit
}

// Match erwarted zwei Karten liefert true,
// wenn die Ränge oder die Farben gleich sind.
func Match(c1, c2 Card) bool {
	return RanksMatch(c1, c2) || SuitsMatch(c1, c2)
}
