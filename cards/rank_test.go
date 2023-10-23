package cards

import "fmt"

// ExampleRanks gibt die Liste der Ränge aus.
func ExampleRanks() {
	for _, r := range Ranks {
		fmt.Println(r)
	}

	// Output:
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// J
	// Q
	// K
	// A
}

// ExampleRankStrings gibt für jeden Rang den int-Wert und den String aus.
func ExampleRankStrings() {
	for _, r := range Ranks {
		fmt.Printf("%d: %s\n", r, r)
	}

	// Output:
	// 2: 2
	// 3: 3
	// 4: 4
	// 5: 5
	// 6: 6
	// 7: 7
	// 8: 8
	// 9: 9
	// 10: 10
	// 11: J
	// 12: Q
	// 13: K
	// 14: A
}

// ExampleRank_print gibt einen Rang auf verschiedene Arten aus.
// Dies soll noch einmal demonstrieren, was sich hinter dem Typ Rank verbirgt.
func ExampleRank_asInt() {
	fmt.Printf("%v\n", Ace) // "verbatim"
	fmt.Printf("%s\n", Ace) // als String
	fmt.Printf("%d\n", Ace) // als int
	fmt.Printf("%T\n", Ace) // den Typ

	// Output:
	// A
	// A
	// 14
	// cards.Rank
}
