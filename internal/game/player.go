package game

import (
	. "blackjack/internal/deck"
	"fmt"
)

// Игрок
type Player struct {
	Name       string
	Hand       []Card
	Balance    int
	CurrentBet int
}

// Дилер
type Dealer struct {
	Player
}

// Выводим список игроков
func (game *Game) PrintPlayers() {
	fmt.Println(game.Players)
}

// Считаем значения руки
func (player Player) CountHandValue() Value {
	var value Value
	var ace Value

	//суммируем значения карт, если был туз добавляем к счетчику
	for _, card := range player.Hand {
		value += card.Value
		if card.Value == Ace {
			ace++
		}
	}

	//если суммарное значения карт больше 21 то туз будет равен 1
	for ace > 0 && value > 21 {
		ace -= 10
	}

	return value
}
