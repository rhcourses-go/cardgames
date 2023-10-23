package hands

import (
	"fmt"
	"strings"

	"github.com/rhcourses-go/cardgames/cards"
)

// Hand ist ein Typ für eine Hand von Karten.
type Hand struct {
	cards []cards.Card
}

// New erzeugt eine neue Hand.
func New() Hand {
	return Hand{}
}

// String gibt die Hand als String zurück.
func (h Hand) String() string {
	return fmt.Sprintf("%s", []cards.Card(h.cards))
}

// Add fügt eine oder mehrere Karten zur Hand hinzu.
func (h *Hand) Add(c ...cards.Card) {
	h.cards = append(h.cards, c...)
}

// Remove erwartet eine Position und entfernt
// diese Karte aus der Hand, falls sie existiert.
// Falls die Position nicht existiert, passiert nichts.
func (h *Hand) Remove(pos int) {
	if pos < 0 || pos >= len(h.cards) {
		return
	}
	h.cards = append(h.cards[:pos], h.cards[pos+1:]...)
}

// Draw erwartet eine Position, gibt diese Karte der Hand zurück
// und entfernt sie aus der Hand.
// Falls die Position nicht existiert, wird eine Panic ausgelöst.
func (h *Hand) Draw(pos int) cards.Card {
	if pos < 0 || pos >= len(h.cards) {
		panic(fmt.Sprintf("ungültige Position: %d", pos))
	}
	c := h.cards[pos]
	h.Remove(pos)
	return c
}

// Image liefert einen String, der die Bilder der Karten enthält.
func (h Hand) Image() string {
	imagelines := make([]string, 10)
	for _, c := range h.cards {
		cardimage := c.Image()
		cardlines := strings.Split(cardimage, "\n")
		for i, line := range cardlines {
			imagelines[i] += line
		}
	}
	return strings.Join(imagelines, "\n")
}
