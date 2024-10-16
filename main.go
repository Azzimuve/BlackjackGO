package main

import (
	"fmt"
	"math/rand"
)

type Suit string //Масть
type Value int   //Значение карты

// Масти
const (
	Hearts   Suit = "Черви"
	Diamonds Suit = "Бубны"
	Spades   Suit = "Пики"
	Clubs    Suit = "Крести"
)

// Номинал карт
const (
	Two   Value = 2
	Three Value = 3
	Four  Value = 4
	Five  Value = 5
	Six   Value = 6
	Seven Value = 7
	Eight Value = 8
	Nine  Value = 9
	Ten   Value = 10
	Jack  Value = 10
	Queen Value = 10
	King  Value = 10
	Ace   Value = 11
)

// Карта
type Card struct {
	Suit  Suit
	Value Value
}

// Колода
type Deck struct {
	Cards []Card
}

// Выводим колоду
func (deck *Deck) PrintDeck() {
	fmt.Println(deck.Cards)
}

// Выводим список игроков
func (game *Game) PrintPlayers() {
	fmt.Println(game.Players)
}

// Перемешиваем колоду
func (deck *Deck) ShuffleDeck() {
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

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

// Игра
type Game struct {
	Players []Player
	Dealer  *Dealer
	Deck    *Deck
	Bank    int
}

// Создание колоды
func CreateDeck() *Deck {
	suits := []Suit{Hearts, Diamonds, Spades, Clubs}
	values := []Value{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	var cards []Card

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, Card{Suit: suit, Value: value})
		}
	}
	return &Deck{Cards: cards}
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
	game.Players = append(game.Players, Player{Name: name, Balance: balance})
}

// вытягиваем карту из колоды
func (deck *Deck) DrawCard() Card {
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
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

// стартуем
func main() {
	game := CreateGame()
	game.AddPlayer("Stepan", 1000)
	for i := 0; i < 5; i++ {
		game.AddPlayer("Stepan "+fmt.Sprint(i), 1000)
	}
	game.DealCards()
	game.PrintPlayers()

	for _, player := range game.Players {
		fmt.Println(player.CountHandValue())
	}
	fmt.Println(game.Dealer.Hand)
	fmt.Println(game.Dealer.CountHandValue())

}
