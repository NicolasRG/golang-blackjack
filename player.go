package main

import (
	"fmt"
)

type Player struct {
	nonAcesCards []Card
	acesCards    []Card
	total        int8
}

func (player *Player) printHand() {
	fmt.Println(player.nonAcesCards)
	fmt.Println(player.acesCards)
	fmt.Println(player.total)
}

func (player *Player) addRegularCardToHand(card Card) {
	player.nonAcesCards = append(player.nonAcesCards, card)
}

func (player *Player) addAceToHand(card Card) {
	player.acesCards = append(player.acesCards, card)
}

func (player *Player) recaculateHand() {
	player.total = 0

	//calculate regular cards first
	for _, card := range player.nonAcesCards {
		player.total = player.total + card.value
	}

	if len(player.acesCards) == 0 {
		return
	}

	//calculate aces last
	if len(player.acesCards) < 2 {
		if player.total <= 10 {
			player.total = player.total + 11
		} else {
			player.total++
		}
		return
	} else {
		//if you have more than one ace you can only have on be an 11 anyway
		if player.total >= int8(11-len(player.acesCards)) {
			player.total = player.total + 11
			player.total = player.total + int8(len(player.acesCards)) - 1
			return
		} else {
			player.total = player.total + int8(len(player.acesCards))
			return
		}
	}
}
