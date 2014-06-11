package ekcore

import (
  "testing"
)

func Test_AddCard(t *testing.T) {
  deck := new(Deck)
  card := Card{ Name: "Cerato", Level: 10 }
  deck.AddCard(card)

  expected := 1
  got := len(deck.Cards)
  if expected != got {
    t.Errorf("Expected the lenth of the cards to be %v, got %v", expected, got)
  }
}

func Test_AddRune(t *testing.T) {
  deck := new(Deck)
  theRune := Rune{ Name: "Dirt", Level: 4 }
  deck.AddRune(theRune)

  expected := 1
  got := len(deck.Runes)
  if expected != got {
    t.Error("Expected the length of the runes to be %v, got %v", expected, got)
  }
}