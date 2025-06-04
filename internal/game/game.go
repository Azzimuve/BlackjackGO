package game

import (
	. "blackjack/internal/deck"
	"fmt"
)

// Игра
type Game struct {
	Players []*Player
	Dealer  *Dealer
	Deck    *Deck
	Bank    int
}

// Добавляем в игру колоду и дилера
func CreateGame() *Game {
	deck := CreateDeck()
	deck.ShuffleDeck()
	dealer := &Dealer{Player: Player{Name: "Dealer", Deck: deck}}
	return &Game{
		Deck:   deck,
		Dealer: dealer,
	}
}

// Добавление в игру игрока
func (game *Game) AddPlayer(name string, balance int, ai bool) {
	game.Players = append(game.Players, &Player{Name: name, Balance: balance, Ai: ai, Deck: game.Deck})
}

// Раздача 2 карт всем участникам
func (game Game) DealCards() {
	//Раздача карт игрокам
	for _, player := range game.Players {
		player.Hand = append(player.Hand, game.Deck.DrawCard(), game.Deck.DrawCard())
	}
	//Раздача карт дилеру
	game.Dealer.Hand = append(game.Dealer.Hand, game.Deck.DrawCard(), game.Deck.DrawCard())
}

func (game *Game) StartRound() {
	fmt.Println(game.Dealer.Name, " карты дилера: ", game.Dealer.Hand[0], ", ***")
	for _, player := range game.Players {
		player.Bet(200)
		fmt.Println("Игрок: ", player.Name)
		fmt.Println("Текущая рука: ", player.Hand, " Сумма карт: ", player.CountHandValue())
		if player.Ai {
			handValue := player.CountHandValue()
			if handValue > 14 && handValue <= 17 {
				player.Hit()
				fmt.Println(player.Name, " берет карту")
				fmt.Println("Текущая рука: ", player.Hand, " Сумма карт: ", player.CountHandValue(), " Ставка: ", player.CurrentBet)

			} else if handValue > 10 && handValue <= 14 {
				player.Double()
				fmt.Println(player.Name, " удваивает ставку и берет карту")
				fmt.Println("Текущая рука: ", player.Hand, " Сумма карт: ", player.CountHandValue(), " Ставка: ", player.CurrentBet)
			}
		}
		fmt.Println("-------------------------------------------")
	}
	fmt.Println("Карты дилера: ", game.Dealer.Hand, "Сумма карт: ")
	for game.Dealer.CountHandValue() < 17 {
		game.Dealer.Hit()
		fmt.Println("Диллер добирает карту")
		fmt.Println("Карты диллера: ", game.Dealer.Hand, " Сумма карт: ", game.Dealer.CountHandValue())
	}
}
