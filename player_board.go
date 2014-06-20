package ekcore

import (
	"fmt"
)

type PlayerBoard struct {
	Player           Player
	CurrentHitPoints int
	Deck             *CardContainer
	Hand             *CardContainer
	Battlefield      *CardContainer
	Cemetery         *CardContainer
}

func NewPlayerBoard(player Player, deck Deck) PlayerBoard {
	board := PlayerBoard{Player: player}
	board.CurrentHitPoints = player.MaxHitPoints()
	board.Deck = new(CardContainer)
	board.Hand = new(CardContainer)
	board.Battlefield = new(CardContainer)
	board.Cemetery = new(CardContainer)
	for _, card := range deck.Cards {
		board.Deck.AddCard(&card)
	}
	return board
}

func (pb *PlayerBoard) DrawCard() {
	card := pb.Deck.DrawCard()
	pb.Hand.AddCard(card)
	fmt.Printf("Added %v to the hand\n", card.Name)
}

func (pb *PlayerBoard) KillCard(idx int) {
	card := pb.Battlefield.Cards[idx]
	result := pb.Battlefield.RemoveCard(card)
	if result != true {
		panic("Couldn't find card to remove")
	}
	pb.Cemetery.AddCard(card)
}
