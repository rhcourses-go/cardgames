package cards

import "fmt"

func ExampleRanksMatch() {
	c1 := New(Ace, Spades)
	c2 := New(Ace, Hearts)
	c3 := New(Ace, Diamonds)
	c4 := New(Ace, Clubs)
	c5 := New(Queen, Spades)

	fmt.Println(RanksMatch(c1, c2))
	fmt.Println(RanksMatch(c1, c3))
	fmt.Println(RanksMatch(c1, c4))
	fmt.Println(RanksMatch(c1, c5))

	// Output:
	// true
	// true
	// true
	// false
}

func ExampleSuitsMatch() {
	c1 := New(Ace, Spades)
	c2 := New(Ace, Hearts)
	c3 := New(Ace, Diamonds)
	c4 := New(Ace, Clubs)
	c5 := New(Queen, Spades)

	fmt.Println(SuitsMatch(c1, c2))
	fmt.Println(SuitsMatch(c1, c3))
	fmt.Println(SuitsMatch(c1, c4))
	fmt.Println(SuitsMatch(c1, c5))

	// Output:
	// false
	// false
	// false
	// true
}

func ExampleMatch() {
	c1 := New(Ace, Spades)
	c2 := New(Ace, Hearts)
	c3 := New(Ace, Diamonds)
	c4 := New(Ace, Clubs)
	c5 := New(Queen, Spades)
	c6 := New(Queen, Hearts)

	fmt.Println(Match(c1, c2))
	fmt.Println(Match(c1, c3))
	fmt.Println(Match(c1, c4))
	fmt.Println(Match(c1, c5))
	fmt.Println(Match(c1, c6))

	// Output:
	// true
	// true
	// true
	// true
	// false
}
