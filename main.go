package main

import(
	"fmt"
)

type Bet int
const(
	NoBET Bet = iota
	TICHU
	GRAND_TICHU
)

type Suit int
const(
	Spade Suit = iota
	Heart
	Diamond
	Club
	Special
)

type Card struct {
	Suit Suit
	Rank int	// 2-10, plus J,Q,K,A (11,12,13,14)
}
const(
	Sparrow = Card{Special, 1}
	Dragon = Card{Special, 2}
	Pheonix = Card{Special, 3}
	Dog = Card{Special, 4}
)



type Player struct {
	Bet Bet
	Hand []Card
	Winnings []Card
}

type Game struct {
	Players [4]Player
	CurrentPlayer int
	// TODO handle three human players + dummy
}

func generate_deck() []Card {
	var deck []Card // 52 + 4
	deck[0] = Sparrow
	deck[1] = Dragon
	deck[2] = Pheonix
	deck[3] = Dog

	currentIndex := 4
	suits := []Suit {Spade, Heart, Diamond, Club}
	for _, suit := range suits {
		for rank := 2; rank <= 14; rank++ {
			deck[currentIndex] = Card{Suit: suit, Rank: rank}
			currentIndex += 1
		}
	}

	return deck
}

func shuffle_deck(deck *[]Card) {

}

func main() {
	// play a game between four Tichu players
	deck := generate_deck()
	shuffle_deck(&deck)

	var game Game
	for i := 0; i < 4; i++ {
		hand := deck[:14]
		deck = deck[14:]

		game.Players[i] = Player{Bet: NoBET, Hand: hand}
	}

	fmt.Println(game)
}