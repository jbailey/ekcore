package ekcore

import (
	"fmt"
  "log"
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
	board.Deck = NewCardContainer()
	board.Hand = NewCardContainer()
	board.Battlefield = NewCardContainer()
	board.Cemetery = NewCardContainer()

	for _, card := range deck.Cards {
		board.Deck.AddCard(card)
	}

	return board
}

func (pb *PlayerBoard) DrawCard() {
	card := pb.Deck.DrawCard()
	if card != nil { 
    pb.Hand.AddCard(card) 
    log.Printf("%v draws %v\n", pb.Player.Name, card.Name)
  }
}

func (pb *PlayerBoard) KillCard(idx int) {
	card := pb.Battlefield.Cards[idx]
	pb.Battlefield.RemoveCard(card)
	pb.Cemetery.AddCard(card)
  log.Printf("%v died", card.Name)
}

func (pb *PlayerBoard) DecrementWaits(amount int) {
  for _, card := range pb.Hand.Cards {
    card.Wait -= amount
  }
}

func (pb *PlayerBoard) MoveExpiredHandCards() {
  for _, card := range pb.Hand.Cards {
    if card.Wait <= 0 {
      fmt.Println("Card added to battlefield: ", card.Name)
      pb.Hand.RemoveCard(card)
      pb.Battlefield.AddCard(card)
    }
  }
}

func (pb *PlayerBoard) PrintBoard() {
  fmt.Printf("||| %v's Board(%d HP) |||\n", pb.Player.Name, pb.CurrentHitPoints)
  fmt.Printf("Deck: ")
  for _, card := range pb.Deck.Cards {
    fmt.Printf("%v(%d)  ", card.Name, card.Wait) 
  }
  fmt.Println()
  fmt.Printf("Hand: ")
  for _, card := range pb.Hand.Cards {
    fmt.Printf("%v(%d)  ", card.Name, card.Wait) 
  }
  fmt.Println()
  fmt.Printf("Battlefield: ")
  for idx, card := range pb.Battlefield.Cards {
    fmt.Printf("[%d]%v(%d/%d)  ", idx+1, card.Name, card.CurrentAttack, card.CurrentHitPoints) 
  }
  fmt.Println()
  fmt.Printf("Cemetery: ")
  for _, card := range pb.Cemetery.Cards {
    fmt.Printf("%v(%d/%d)  ", card.Name, card.CurrentAttack, card.CurrentHitPoints) 
  }
  fmt.Println()
}