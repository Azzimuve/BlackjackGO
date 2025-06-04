package main

import (
	g "blackjack/internal/game"
	"fmt"
)

// стартуем
func main() {
	game := g.CreateGame()
	for i := 0; i < 2; i++ {
		game.AddPlayer("AI "+fmt.Sprint(i), 1000, true)
	}
	for i := 0; i < 2; i++ {
		game.AddPlayer("Stepan "+fmt.Sprint(i), 1000, false)
	}
	game.DealCards()

	game.StartRound()

}
