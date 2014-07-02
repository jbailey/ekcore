package ekcore

import (
  "fmt"
  "math/rand"
  "time"
)

type CardContainer struct {
  Slot  int
  Cards []*Card
}

func NewCardContainer() *CardContainer {
  rand.Seed(time.Now().UTC().UnixNano())
  return &CardContainer{ Cards: make([]*Card, 0) }
}

// func NewCardContainerWithSlots(numSlot int) *CardContainer {
//   rand.Seed(time.Now().UTC().UnixNano())
//   return &CardContainer{ Cards: make([]*Card, numSlot) }
// }

func (cc *CardContainer) Size() int {
  return len(cc.Cards)
}

func (cc *CardContainer) GetSlotNumForCard(card *Card) int {
  for idx, c := range cc.Cards {
    if c == card {
      return idx + 1
    }
  }
  return -1
}

func (cc *CardContainer) CardAtSlot(slotNum int) *Card {
  var card *Card

  if len(cc.Cards) >= slotNum {
    card = cc.Cards[slotNum-1]
  } else {
    card = nil
  }

  return card
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
      cc.Cards = cc.Cards[:len(cc.Cards)-1] // reslice
      return 
    }
  }
}

// Removes the card but does not collapse the slice
func (cc *CardContainer) KillCard(card *Card) {
  for i, c := range cc.Cards {
    if c == card {
      cc.Cards[i] = nil
      return
    }
  }
}

func (cc *CardContainer) Collapse() {
  newCC := NewCardContainer()
  for _, card := range cc.Cards {
    if card != nil { newCC.AddCard(card) }
  }

  cc.Cards = newCC.Cards
}

func (cc *CardContainer) SelectRandomCard() *Card {
  var card *Card
  cards := cc.SelectRandomCards(1)
  if len(cards) > 0 {
    card = cards[0]
  } else {
    card = nil
  }

  return card
}

func (cc *CardContainer) SelectRandomCards(num int) []*Card {
  if len(cc.Cards) == 0 { return nil }
  tempCC := &CardContainer{ Cards: cc.nonNilCards() }
  tempCC.Shuffle()
  if num > len(tempCC.Cards) { num = len(tempCC.Cards) }
  return tempCC.Cards[:num]
}

func (cc *CardContainer) nonNilCards() []*Card {
  tempCC := NewCardContainer()
  for _, card := range cc.Cards {
    if card != nil {
      tempCC.AddCard(card)
    }
  }
  return tempCC.Cards
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