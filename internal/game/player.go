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
	Deck       *Deck
	Ai         bool
}

// Дилер
type Dealer struct {
	Player
}

// Выводим список игроков
func (game *Game) PrintPlayers() {
	fmt.Println(game.Players)
}

// Делаем ставку
func (player *Player) Bet(bet int) bool {
	if player.Balance < bet {
		fmt.Println("У вас недостаточно средств")
		return false
	} else {
		player.Balance -= bet
		player.CurrentBet += bet
		return true
	}
}

// Добираем карту
func (player *Player) Hit() {
	player.Hand = append(player.Hand, player.Deck.DrawCard())
}

// Удваиваем ставку и добираем карту
func (player *Player) Double() {
	if player.Bet(player.CurrentBet) {
		player.Hit()
	}

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
