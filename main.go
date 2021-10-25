package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Welcome to blackjack!!")

	//initialize deck
	deck := InitDeck()

	//create number to beat
	rand.Seed(time.Now().UnixMilli())
	cpuAns := rand.Intn(6) + 15

	//initialize player
	player := Player{
		nonAcesCards: []Card{},
		acesCards:    []Card{},
		total:        0,
	}

	//draw first two cards
	addCardToHand(&player, &deck)
	addCardToHand(&player, &deck)

	player.recaculateHand()
	player.printHand()

	str := ""
	//loop till busted or player stays
Loop:
	for {
		fmt.Println("Would you like to hit again? Type 'Y' for yes and anything else for no...")
		fmt.Scanln(&str)

		if str == "Y" || str == "y" {
			addCardToHand(&player, &deck)
			player.recaculateHand()
			player.printHand()

			if player.total == 21 {
				fmt.Println("You Won!!")
				return
			} else if player.total > 21 {
				fmt.Println("You busted")
				return
			}

		} else {
			fmt.Println("the cpus number is : " + strconv.Itoa(cpuAns))
			if player.total > int8(cpuAns) {
				fmt.Println("You Won!!")
			} else {
				fmt.Println("You Lost!!")
			}

			break Loop
		}

	}
}

func addCardToHand(player *Player, deck *Deck) {
	card := deck.drawCard()

	if card.isAce {
		player.addAceToHand(card)
	} else {
		player.addRegularCardToHand(card)
	}
}
