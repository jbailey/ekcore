package ekcore

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_PlayerBoard_DrawCard(t *testing.T) {
  card1 := Card{ Name: "Cerato", Level: 10 }
  card2 := Card{ Name: "Terra Cotta Warrior", Level: 5 }
  card3 := Card{ Name: "Ice Dragon", Level: 10 }
  cards := []Card{ card1, card2, card3 }
  player := Player{ Name: "The Player", Level: 40 }
  deck := Deck{ Cards: cards }
  playerBoard := NewPlayerBoard(player, deck)

  playerBoard.DrawCard()
  assert.Equal(t, playerBoard.Deck.Size(), 2)
  assert.Equal(t, playerBoard.Hand.Size(), 1)
  playerBoard.DrawCard()
  assert.Equal(t, playerBoard.Deck.Size(), 1)
  assert.Equal(t, playerBoard.Hand.Size(), 2)

  // After last draw, should have all 3 cards in the hand
  playerBoard.DrawCard()
  playerBoard.Hand.PrintCards()
  assert.True(t, playerBoard.Hand.DoesContain(&card1), "hand should contain card1")
  assert.True(t, playerBoard.Hand.DoesContain(&card2), "hand should contain card2")
  assert.True(t, playerBoard.Hand.DoesContain(&card3), "hand should contain card3")
}

func Test_PlayerBoard_KillCard(t *testing.T) {
  card1 := Card{ Name: "Cerato", Level: 10 }
  card2 := Card{ Name: "Terra Cotta Warrior", Level: 5 }
  card3 := Card{ Name: "Ice Dragon", Level: 10 }
  cards := []Card{ card1, card2, card3 }
  player := Player{ Name: "The Player", Level: 40 }
  deck := Deck{ Cards: cards }
  playerBoard := NewPlayerBoard(player, deck)
  // for x := 0; x <= len(playerBoard.Deck.Cards); x++ {
  //   card := playerBoard.Deck.DrawCard()
  //   playerBoard.Battlefield.AddCard(card)
  // }
  playerBoard.Battlefield.AddCard(&card1)
  playerBoard.Battlefield.AddCard(&card2)
  playerBoard.Battlefield.AddCard(&card3)

  playerBoard.KillCard(1)
  assert.True(t, playerBoard.Battlefield.DoesContain(&card1), "Expected the battlefield to contain card1")
  assert.False(t, playerBoard.Battlefield.DoesContain(&card2), "Expected the battlefield NOT to contain card2")
  assert.True(t, playerBoard.Battlefield.DoesContain(&card3), "Expected the battlefield to contain card3")
  assert.True(t, playerBoard.Cemetery.DoesContain(&card2), "Expected the cemetery to contain card2")

}
