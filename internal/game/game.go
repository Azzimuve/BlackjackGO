package game

import (
	. "blackjack/internal/deck"
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
	dealer := &Dealer{Player: Player{Name: "Dealer"}}
	return &Game{
		Deck:   deck,
		Dealer: dealer,
	}
}

// Добавление в игру игрока
func (game *Game) AddPlayer(name string, balance int) {
	game.Players = append(game.Players, &Player{Name: name, Balance: balance})
}

// Раздача 2 карт всем участникам
func (game Game) DealCards() {
	//Раздача карт игрокам
	for player := range game.Players {
		game.Players[player].Hand = append(game.Players[player].Hand, game.Deck.DrawCard(), game.Deck.DrawCard())
	}
	//Раздача карт дилеру
	game.Dealer.Hand = append(game.Dealer.Hand, game.Deck.DrawCard(), game.Deck.DrawCard())
}
