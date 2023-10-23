package cards

import "fmt"

// ExampleCard zeigt die Verwendung von Card.
func ExampleCard() {
	fmt.Println(New(Ace, Spades))
	fmt.Println(New(Queen, Hearts))
	fmt.Println(New(Two, Diamonds))
	fmt.Println(New(Ten, Clubs))

	// Output:
	// (♠ A)
	// (♥ Q)
	// (♦ 2)
	// (♣ 10)
}

// ExampleCard_Image zeigt die Verwendung von Card.Image.
func ExampleCard_Image() {
	fmt.Println(New(Ace, Spades).Image())
	fmt.Println(New(Queen, Hearts).Image())
	fmt.Println(New(Two, Diamonds).Image())
	fmt.Println(New(Ten, Clubs).Image())

	// Output:
	// ┌─────────┐
	// │A        │
	// │♠        │
	// │         │
	// │    A    │
	// │         │
	// │        ♠│
	// │        A│
	// └─────────┘
	//
	// ┌─────────┐
	// │Q        │
	// │♥        │
	// │         │
	// │    Q    │
	// │         │
	// │        ♥│
	// │        Q│
	// └─────────┘
	//
	// ┌─────────┐
	// │2        │
	// │♦        │
	// │         │
	// │    2    │
	// │         │
	// │        ♦│
	// │        2│
	// └─────────┘
	//
	// ┌─────────┐
	// │10       │
	// │♣        │
	// │         │
	// │    10   │
	// │         │
	// │        ♣│
	// │       10│
	// └─────────┘
}
