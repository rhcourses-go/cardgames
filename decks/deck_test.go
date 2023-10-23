package decks

import "fmt"

// ExampleNew32 erzeugt ein Skat-Blatt und gibt es aus.
// Es werden also für jede Farbe die Karten von 7 bis Ass ausgegeben.
func ExampleNew32() {
	d1 := New32()
	fmt.Printf("Kreuz: %v\n", d1.cards[0:8])
	fmt.Printf("Pik:   %v\n", d1.cards[8:16])
	fmt.Printf("Herz:  %v\n", d1.cards[16:24])
	fmt.Printf("Karo:  %v\n", d1.cards[24:32])

	// Output:
	// Kreuz: [(♣ 7) (♣ 8) (♣ 9) (♣ 10) (♣ J) (♣ Q) (♣ K) (♣ A)]
	// Pik:   [(♠ 7) (♠ 8) (♠ 9) (♠ 10) (♠ J) (♠ Q) (♠ K) (♠ A)]
	// Herz:  [(♥ 7) (♥ 8) (♥ 9) (♥ 10) (♥ J) (♥ Q) (♥ K) (♥ A)]
	// Karo:  [(♦ 7) (♦ 8) (♦ 9) (♦ 10) (♦ J) (♦ Q) (♦ K) (♦ A)]
}

// ExampleNew52 erzeugt ein Poker-Blatt und gibt es aus.
// Es werden also für jede Farbe die Karten von 2 bis Ass ausgegeben.
func ExampleNew52() {
	d1 := New52()
	fmt.Printf("Kreuz: %v\n", d1.cards[0:13])
	fmt.Printf("Pik:   %v\n", d1.cards[13:26])
	fmt.Printf("Herz:  %v\n", d1.cards[26:39])
	fmt.Printf("Karo:  %v\n", d1.cards[39:52])

	// Output:
	// Kreuz: [(♣ 2) (♣ 3) (♣ 4) (♣ 5) (♣ 6) (♣ 7) (♣ 8) (♣ 9) (♣ 10) (♣ J) (♣ Q) (♣ K) (♣ A)]
	// Pik:   [(♠ 2) (♠ 3) (♠ 4) (♠ 5) (♠ 6) (♠ 7) (♠ 8) (♠ 9) (♠ 10) (♠ J) (♠ Q) (♠ K) (♠ A)]
	// Herz:  [(♥ 2) (♥ 3) (♥ 4) (♥ 5) (♥ 6) (♥ 7) (♥ 8) (♥ 9) (♥ 10) (♥ J) (♥ Q) (♥ K) (♥ A)]
	// Karo:  [(♦ 2) (♦ 3) (♦ 4) (♦ 5) (♦ 6) (♦ 7) (♦ 8) (♦ 9) (♦ 10) (♦ J) (♦ Q) (♦ K) (♦ A)]
}
