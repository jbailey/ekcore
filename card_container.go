package ekcore

import (
  "fmt"
)

type CardContainer struct {
  Cards []*Card
}

func (cc *CardContainer) Size() int {
  return len(cc.Cards)
}

func (cc *CardContainer) AddCard(card *Card) {
  cc.Cards = append(cc.Cards, card)
}

func (cc *CardContainer) DrawCard() *Card {
  card := cc.Cards[0]
  cc.Cards = cc.Cards[1:]
  cc.PrintCards()
  return card
}

func (cc *CardContainer) RemoveCard(card *Card) bool {
  for idx, c := range cc.Cards {
    if c == card {
      if idx + 1 == len(cc.Cards) {
        cc.Cards = cc.Cards[:idx]
      } else {
        cc.Cards = append(cc.Cards[:idx], cc.Cards[idx+1:]...)
      }
      return true
    }
  }
  return false
}

func (cc *CardContainer) Shuffle() {

}

func (cc *CardContainer) DoesContain(card *Card) bool {
  // fmt.Println("Looking for:", card)
  found := false
  for _, c := range cc.Cards {
    // fmt.Println(c)
    if c == card {
      // fmt.Println("Found it")
      found = true
      break
    }
  }
  // fmt.Println("Finished Looking")
  return found
}

func (cc *CardContainer) PrintCards() {
  fmt.Println("Cards in Container:")
  for _, card := range cc.Cards {
    fmt.Println(card.Name)
  }
}