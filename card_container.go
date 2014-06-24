package ekcore

import (
  "fmt"
  "math/rand"
  "time"
)

type CardContainer struct {
  Cards []*Card
}

func NewCardContainer() *CardContainer {
  rand.Seed(time.Now().UTC().UnixNano())
  return &CardContainer{ Cards: make([]*Card, 0) }
}

func (cc *CardContainer) Size() int {
  return len(cc.Cards)
}

func (cc *CardContainer) AddCard(card *Card) {
  // fmt.Printf("Adding %v to the container\n", card.Name)
  cc.Cards = append(cc.Cards, card)
}

func (cc *CardContainer) DrawCard() *Card {
  if len(cc.Cards) > 0 {
    card := cc.Cards[0]
    cc.Cards = cc.Cards[1:]
    return card
  }
  return nil
}

func (cc *CardContainer) RemoveCard(card *Card) {
  for i, c := range cc.Cards {
    if c == card {
      copy(cc.Cards[i:], cc.Cards[i+1:])  // Shift
      // cc.Cards[len(cc.Cards)-1] = nil     // remove the reference
      cc.Cards = cc.Cards[:len(cc.Cards)-1] // reslice
      return 
    }
  }
}

func (cc *CardContainer) Shuffle() {
  for i := 1; i < len(cc.Cards); i++ {
    r := rand.Intn(i + 1)
    if i != r {
      cc.Cards[r], cc.Cards[i] = cc.Cards[i], cc.Cards[r]
    }
  }
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
  // fmt.Println("Cards in Container:")
  for _, card := range cc.Cards {
    fmt.Println(card.Name)
  }
}