package hands

import (
	"fmt"

	"github.com/rhcourses-go/cardgames/cards"
)

// ExampleNew erzeugt eine neue Hand und gibt sie aus.
// Eine neue Hand ist leer.
func ExampleNew() {
	h := New()
	fmt.Println(h)

	// Output:
	// []
}

// ExampleHand_addCards fügt Karten zur Hand hinzu und gibt sie aus.
// Zu Demonstrationszwecken werden die Karten "manuell" hinzugefügt.
func ExampleHand_addCards() {
	h1 := New()
	cardlist := []cards.Card{
		cards.New(cards.Ace, cards.Spades),
		cards.New(cards.Two, cards.Hearts),
	}

	h1.cards = append(h1.cards, cardlist...)
	fmt.Println(h1)

	// Output:
	// [(♠ A) (♥ 2)]
}

// ExampleHand_Add fügt eine Karte zur Hand hinzu und gibt sie aus.
// Wie oben, aber mit der Methode Add.
func ExampleHand_Add() {
	h1 := New()
	cardlist := []cards.Card{
		cards.New(cards.Ace, cards.Spades),
		cards.New(cards.Two, cards.Hearts),
	}

	h1.Add(cardlist...)
	fmt.Println(h1)

	// Output:
	// [(♠ A) (♥ 2)]
}

// ExampleHand_Remove entfernt eine Karte aus einer Hand und gibt die Hand aus.
func ExampleHand_Remove() {
	h1 := New()
	cardlist := []cards.Card{
		cards.New(cards.Ace, cards.Spades),
		cards.New(cards.Two, cards.Hearts),
		cards.New(cards.Seven, cards.Hearts),
		cards.New(cards.King, cards.Clubs),
	}
	h1.Add(cardlist...)

	fmt.Println(h1)
	h1.Remove(1)
	fmt.Println(h1)
	h1.Remove(2)
	fmt.Println(h1)
	h1.Remove(2)
	fmt.Println(h1)
	h1.Remove(-1)
	fmt.Println(h1)

	// Output:
	// [(♠ A) (♥ 2) (♥ 7) (♣ K)]
	// [(♠ A) (♥ 7) (♣ K)]
	// [(♠ A) (♥ 7)]
	// [(♠ A) (♥ 7)]
	// [(♠ A) (♥ 7)]
}

// ExampleHand_Draw zieht eine Karte aus einer Hand und gibt die Hand und die Karte aus.
func ExampleHand_Draw() {
	h1 := New()
	cardlist := []cards.Card{
		cards.New(cards.Ace, cards.Spades),
		cards.New(cards.Two, cards.Hearts),
		cards.New(cards.Seven, cards.Hearts),
		cards.New(cards.King, cards.Clubs),
	}
	h1.Add(cardlist...)

	fmt.Println(h1)
	c1 := h1.Draw(1)
	fmt.Println(h1)
	fmt.Println(c1)

	// Output:
	// [(♠ A) (♥ 2) (♥ 7) (♣ K)]
	// [(♠ A) (♥ 7) (♣ K)]
	// (♥ 2)
}

// ExampleHand_Draw_panic zeigt, dass Draw eine Panic auslöst,
// wenn die Position nicht existiert.
func ExampleHand_Draw_panic() {
	h1 := New()
	cardlist := []cards.Card{
		cards.New(cards.Ace, cards.Spades),
		cards.New(cards.Two, cards.Hearts),
		cards.New(cards.Seven, cards.Hearts),
		cards.New(cards.King, cards.Clubs),
	}
	h1.Add(cardlist...)

	fmt.Println(h1)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()
	h1.Draw(5)

	// Output:
	// [(♠ A) (♥ 2) (♥ 7) (♣ K)]
	// Panic: ungültige Position: 5
}

// ExampleHand_Image gibt die Bilder der Karten einer Hand aus.
func ExampleHand_Image() {
	h1 := New()
	cardlist := []cards.Card{
		cards.New(cards.Ace, cards.Spades),
		cards.New(cards.Two, cards.Hearts),
		cards.New(cards.Seven, cards.Hearts),
		cards.New(cards.King, cards.Clubs),
	}
	h1.Add(cardlist...)

	fmt.Println(h1.Image())

	// Output:
	// ┌─────────┐┌─────────┐┌─────────┐┌─────────┐
	// │A        ││2        ││7        ││K        │
	// │♠        ││♥        ││♥        ││♣        │
	// │         ││         ││         ││         │
	// │    A    ││    2    ││    7    ││    K    │
	// │         ││         ││         ││         │
	// │        ♠││        ♥││        ♥││        ♣│
	// │        A││        2││        7││        K│
	// └─────────┘└─────────┘└─────────┘└─────────┘
}
