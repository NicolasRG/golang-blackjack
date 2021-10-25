package main

import (
	"math/rand"
	"strconv"
	"time"
)

//setup for individual cards
var suites = []string{"Spades", "Clubs", "Hearts", "Diamonds"}
var faceCards = []string{"Jack", "Queen", "King", "Ace"}

type Card struct {
	name  string
	value int8
	isAce bool
}

//make a struct to be a deck
type Deck struct {
	cards []Card
}

func (deck *Deck) drawCard() Card {
	var card Card
	card, deck.cards = deck.cards[0], deck.cards[1:]
	return card
}

func InitDeck() Deck {
	//create deck of cards
	deck := Deck{
		cards: make([]Card, 0, 50),
	}
	deck.createCards()
	//shuffle
	deck.shuffleCards()
	return deck

}

func (deck *Deck) createCards() {
	//create number cards first
	for j := 0; j < len(suites); j++ {
		for i := 2; i < 11; i++ {
			tempCard := Card{
				name:  strconv.Itoa(i) + " of " + suites[j],
				value: int8(i),
				isAce: false,
			}
			deck.cards = append(deck.cards, tempCard)
		}
	}

	//create face cards
	for i := 0; i < len(suites); i++ {
		for j := 0; j < len(faceCards); j++ {
			tempCard := Card{
				name:  faceCards[j] + " of " + suites[i],
				value: 10,
				isAce: false,
			}

			if faceCards[j] == "Ace" {
				tempCard.isAce = true
			}
			deck.cards = append(deck.cards, tempCard)
		}
	}
}

func (deck *Deck) shuffleCards() {

	for i := 0; i < 5000; i++ {
		//generate random spots
		rand.Seed(time.Now().UnixMilli())

		index1, index2 := rand.Intn(50), rand.Intn(50)
		//swap places
		deck.cards[index1], deck.cards[index2] = deck.cards[index2], deck.cards[index1]
	}

}

//for debugging purposes
// func (deck *Deck) printDeck() {
// 	fmt.Println(deck)
// }
