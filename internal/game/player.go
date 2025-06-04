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

// Выводим информацию об игроке
func (player *Player) PrintInfo() {
	fmt.Println("Имя игрока: ", player.Name)
	fmt.Println("Карты игрока: ", player.Hand, "Сумма карт: ", player.CountHandValue(), " Ставка: ", player.CurrentBet)
}

// Добираем карту
func (player *Player) Hit() {
	player.Hand = append(player.Hand, player.Deck.DrawCard())
}

// Удваиваем ставку и добираем карту
func (player *Player) Double() bool {
	if player.Bet(player.CurrentBet) {
		player.Hit()
		return true
	} else {
		return false
	}

}

// Считаем значения руки
func (player *Player) CountHandValue() Value {
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

// Даем выбор игроку только если сумма карт меньше 21
func (player *Player) PlayerChoose() {
	if player.CountHandValue() < 21 {
		var choose string
		fmt.Println("Напишите если вы хотите:\n h - взять карту\n d - удвоить ставку и взять карту\n s - сбросить карты(Вам вернется половина от ставки)\n напишите любую другую букву чтобы продолжить")
		_, err := fmt.Scanln(&choose)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}

		switch choose {
		case "h":
			fmt.Println("вы взяли карту")
			player.Hit()
			player.PrintInfo()
			player.PlayerChoose()

		case "d":
			if player.Double() {
				fmt.Println("вы удвоили ставку и взяли карту")
				player.PrintInfo()
			}

		case "s":
			fmt.Println("вы сбросили карты")

		}
	} else {
		fmt.Println("Вы больше не можете брать карты")
	}

}
