package cards

import "fmt"

// ExampleSuits gibt die Liste der Farben aus.
func ExampleSuits() {
	fmt.Println(Suits)

	// Output:
	// [♣ ♠ ♥ ♦]
}

// ExampleSuitStrings gibt für jede Farbe den int-Wert und den String aus.
func ExampleSuitStrings() {
	for _, s := range Suits {
		fmt.Printf("%d: %s\n", s, s)
	}

	// Output:
	// 0: ♣
	// 1: ♠
	// 2: ♥
	// 3: ♦
}

// ExampleSuit_print gibt eine Farbe auf verschiedene Arten aus.
// Dies soll noch einmal demonstrieren, was sich hinter dem Typ Suit verbirgt.
func ExampleSuit_asInt() {
	fmt.Printf("%v\n", Spades) // "verbatim"
	fmt.Printf("%s\n", Spades) // als String
	fmt.Printf("%d\n", Spades) // als int
	fmt.Printf("%T\n", Spades) // den Typ

	// Output:
	// ♠
	// ♠
	// 1
	// cards.Suit
}
