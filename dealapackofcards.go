package piscine

import "fmt"

type player struct {
	cards [3]int
}

// Player 1: 1, 2, 3$
func DealAPackOfCards(deck []int) {
	players := [4]player{}

	for i := 0; i < 12; i++ {
		players[i/3].cards[i%3] = deck[i]
	}

	for playeri := 0; playeri < 4; playeri++ {
		fmt.Printf("Player %d: %d, %d, %d\n", playeri+1, players[playeri].cards[0], players[playeri].cards[1], players[playeri].cards[2])
	}
}
