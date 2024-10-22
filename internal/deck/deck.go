package deck

import (
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

type DeckService struct {
	Deck *Deck
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

// Перемешиваем колоду
func (deck *Deck) ShuffleDeck() {
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

// Вытягиваем карту из колоды
func (deck *Deck) DrawCard() Card {
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}
