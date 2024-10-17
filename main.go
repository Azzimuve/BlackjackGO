package main

import (
	g "blackjack/internal/game"
	"fmt"
)

// стартуем
func main() {
	game := g.CreateGame()
	game.AddPlayer("Stepan", 1000)
	for i := 0; i < 5; i++ {
		game.AddPlayer("Stepan "+fmt.Sprint(i), 1000)
	}
	game.DealCards()
	game.PrintPlayers()

	for _, player := range game.Players {
		fmt.Println(player.Balance, "ставка ", player.CurrentBet)
		player.Bet(500)
		fmt.Println(player.Balance, "ставка ", player.CurrentBet)
	}

}
