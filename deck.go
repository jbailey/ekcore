package ekcore

import (
  "fmt"
)

type Deck struct {
  Name string
  Cards []Card
  Runes []Rune
}

func (deck *Deck) AddCard(card Card) {
  deck.Cards = append(deck.Cards, card)
}

func (deck *Deck) AddRune(theRune Rune) {
  deck.Runes = append(deck.Runes, theRune)  
}

func (deck *Deck) Shuffle() {
  
}

func (deck *Deck) PrintDeck() {
  fmt.Println("Cards")
   for _, card := range deck.Cards {
    fmt.Println("\t", card.Name)
  } 

  fmt.Println("Runes")
  for _, rune := range deck.Runes {
    fmt.Println("\t", rune.Name)
  }
}