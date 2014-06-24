package ekcore

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_NewPlayerBoard(t *testing.T) {
  card1 := Card{ Name: "Cerato", Level: 10 }
  card2 := Card{ Name: "Terra Cotta Warrior", Level: 5 }
  card3 := Card{ Name: "Ice Dragon", Level: 10 }
  cards := []*Card{ &card1, &card2, &card3 }
  player := Player{ Name: "The Player", Level: 40 }
  deck := Deck{ Cards: cards }
  playerBoard := NewPlayerBoard(player, deck)
  
  assert.Equal(t, 3, playerBoard.Deck.Size(), "PlayerBoard deck should contain 3 cards")
  assert.True(t, playerBoard.Deck.DoesContain(&card1), "hand should contain card1")
  assert.True(t, playerBoard.Deck.DoesContain(&card2), "hand should contain card2")
  assert.True(t, playerBoard.Deck.DoesContain(&card3), "hand should contain card3")
}

func Test_PlayerBoard_DrawCard(t *testing.T) {
  card1 := Card{ Name: "Cerato", Level: 10 }
  card2 := Card{ Name: "Terra Cotta Warrior", Level: 5 }
  card3 := Card{ Name: "Ice Dragon", Level: 10 }
  cards := []*Card{ &card1, &card2, &card3 }
  player := Player{ Name: "The Player", Level: 40 }
  deck := Deck{ Cards: cards }
  playerBoard := NewPlayerBoard(player, deck)
  playerBoard.DrawCard()
  assert.Equal(t, playerBoard.Deck.Size(), 2, "deck should have 2 cards")
  assert.Equal(t, playerBoard.Hand.Size(), 1, "hard should have 1 card")
  playerBoard.DrawCard()
  assert.Equal(t, playerBoard.Deck.Size(), 1, "deck should have 1 card")
  assert.Equal(t, playerBoard.Hand.Size(), 2, "hand should have 2 cards")

  // After last draw, should have all 3 cards in the hand
  playerBoard.DrawCard()
  assert.True(t, playerBoard.Hand.DoesContain(&card1), "hand should contain card1")
  assert.True(t, playerBoard.Hand.DoesContain(&card2), "hand should contain card2")
  assert.True(t, playerBoard.Hand.DoesContain(&card3), "hand should contain card3")
}

func Test_PlayerBoard_KillCard(t *testing.T) {
  card1 := Card{ Name: "Cerato", Level: 10 }
  card2 := Card{ Name: "Terra Cotta Warrior", Level: 5 }
  card3 := Card{ Name: "Ice Dragon", Level: 10 }
  cards := []*Card{ &card1, &card2, &card3 }
  player := Player{ Name: "The Player", Level: 40 }
  deck := Deck{ Cards: cards }
  playerBoard := NewPlayerBoard(player, deck)
  playerBoard.Battlefield.AddCard(&card1)
  playerBoard.Battlefield.AddCard(&card2)
  playerBoard.Battlefield.AddCard(&card3)

  playerBoard.KillCard(1)
  assert.True(t, playerBoard.Battlefield.DoesContain(&card1), "Expected the battlefield to contain card1")
  assert.False(t, playerBoard.Battlefield.DoesContain(&card2), "Expected the battlefield NOT to contain card2")
  assert.True(t, playerBoard.Battlefield.DoesContain(&card3), "Expected the battlefield to contain card3")
  assert.True(t, playerBoard.Cemetery.DoesContain(&card2), "Expected the cemetery to contain card2")

}

func Test_PlayerBoard_DecrementWaits(t *testing.T) {
  card := Card{ Name: "Cerato", Wait: 6 }
  player := Player{ Name: "The Player", Level: 40 }
  cards := []*Card{ &card }
  deck := Deck{ Cards: cards }
  playerBoard := NewPlayerBoard(player, deck)

  playerBoard.DrawCard()
  playerBoard.DecrementWaits(1)
  assert.Equal(t, 5, card.Wait, "Cards's Wait should have been decremented")
  playerBoard.DecrementWaits(2)
  assert.Equal(t, 3, card.Wait, "Cards's Wait should have been decremented")
  playerBoard.DecrementWaits(3)
  assert.Equal(t, 0, card.Wait, "Cards's Wait should have been decremented")
}

func Test_PlayerBoard_MoveExpiredHandCards(t *testing.T) {
  card := Card{ Name: "Cerato", Wait: 0 }
  player := Player{ Name: "The Player", Level: 40 }
  cards := []*Card{ &card }
  deck := Deck{ Cards: cards }
  playerBoard := NewPlayerBoard(player, deck)
  playerBoard.DrawCard()
  playerBoard.MoveExpiredHandCards()
  assert.Equal(t, 0, playerBoard.Hand.Size(), "playerBoard hand should be empty")
  assert.Equal(t, 1, playerBoard.Battlefield.Size(), "playerBoard battlefield should be have 1 card")
  assert.True(t, playerBoard.Battlefield.DoesContain(&card), "playerBoard battlefield should contain Cerato")
}

func printTestHeader(name string) {
  fmt.Println("------------------------------------------------")
  fmt.Println(name)
  fmt.Println("------------------------------------------------")
}
