package ekcore

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_NewCardContainerWithSlots(t *testing.T) {

}

func Test_CardContainer_CardAtSlot(t *testing.T) {
  cc := NewCardContainer()
  card := &Card{Name: "God of Pain"}
  card2 := &Card{Name: "Taiga General"} 
  cc.AddCard(card)
  cc.AddCard(card2)

  assert.Equal(t, card, cc.CardAtSlot(1), "Card at slot 1 should be God of Pain")
  assert.Equal(t, card2, cc.CardAtSlot(2), "Card at slot 2 should be Taiga General")

  // When a card is killed, should return nil
  cc.KillCard(card)
  assert.Nil(t, cc.CardAtSlot(1), "Should be nil since the card that used to be there has been killed")
}

func Test_CardContainer_AddCard(t *testing.T) {
  cc := new(CardContainer)

  card := Card{Name: "Cerato"} 
  cc.AddCard(&card)
  assert.Equal(t, 1, cc.Size(), "should be equal")

  card2 := Card{Name: "Cerato"} 
  cc.AddCard(&card2)
  assert.Equal(t, 2, cc.Size(), "should be equal")

  card3 := Card{Name: "Ice Dragon"} 
  cc.AddCard(&card3)
  assert.Equal(t, 3, cc.Size(), "should be equal")
}

func Test_CardContainer_DoesContain(t *testing.T) {
  cc := new(CardContainer)
  card := Card{Name: "God of Pain"}
  card2 := Card{Name: "Taiga General"} 
  cc.AddCard(&card)
  cc.AddCard(&card2)

  assert.True(t, cc.DoesContain(&card))
  assert.True(t, cc.DoesContain(&card2))
}

func Test_CardContainer_DrawCard(t *testing.T) {
  cc := new(CardContainer)
  card := Card{Name: "Cerato"}
  card2 := Card{Name: "Cerato"} 
  cc.AddCard(&card)
  cc.AddCard(&card2)
  assert.Equal(t, 2, cc.Size(), "should be equal")

  drawnCard := cc.DrawCard()
  assert.NotNil(t, drawnCard, "Drawn Card should not be nil")
  assert.Equal(t, 1, cc.Size(), "Drawn card should have been removed")

  secondDrawnCard := cc.DrawCard()
  assert.NotNil(t, secondDrawnCard, "Second drawn Card should not be nil")
  assert.Equal(t, 0, cc.Size(), "Drawn card should have been removed")

  // with empty hand
  cc = NewCardContainer()
  assert.Nil(t, cc.DrawCard(), "Should be nil when container is empty")
}

func Test_CardContainer_RemoveCard(t *testing.T) {
  cc := new(CardContainer)
  card := Card{Name: "Cerato"}
  card2 := Card{Name: "Cerato"} 
  cc.AddCard(&card)
  cc.AddCard(&card2)

  cc.RemoveCard(&card)
  assert.Equal(t, 1, cc.Size(), "Size should have decreased")
  assert.False(t, cc.DoesContain(&card), "Container should not have card")
  assert.True(t, cc.DoesContain(&card2), "Container should have card2")
  cc.RemoveCard(&card2)
  assert.Equal(t, 0, cc.Size(), "Container should be empty")
}

// KillCard
func Test_CardContainer_KillCard(t *testing.T) {
  cc := new(CardContainer)
  card := Card{Name: "Cerato"}
  card2 := Card{Name: "Cerato"} 
  cc.AddCard(&card)
  cc.AddCard(&card2)

  cc.KillCard(&card)
  assert.Equal(t, 2, cc.Size(), "Container should still have size of 2, but with only 1 card")
  assert.False(t, cc.DoesContain(&card), "Container should not have card")
  assert.True(t, cc.DoesContain(&card2), "Container should have card2")
  assert.Nil(t, cc.CardAtSlot(1), "Slot 1 should be nil")
}

// Collapse
func Test_CardContainer_Collapse(t *testing.T) {
  cc := NewCardContainer()
  card := &Card{ Name: "Giant Crab" }
  card2 := &Card{ Name: "Orc Warrior" }
  cc.AddCard(card)
  cc.AddCard(card2)

  cc.KillCard(card)
  assert.Equal(t, 2, cc.Size(), "Container size should still be 2, 1 is nil")
  cc.Collapse()
  assert.Equal(t, 1, cc.Size(), "Container should only have 1 card")
}

// SelectRandomCards
func Test_CardContainer_SelectRandomCards(t *testing.T) {
  cc := NewCardContainer()
  card1 := &Card{ Name: "Card1" }
  card2 := &Card{ Name: "Card2" }
  card3 := &Card{ Name: "Card3" }
  card4 := &Card{ Name: "Card4" }
  cc.AddCard(card1)
  cc.AddCard(card2)
  cc.AddCard(card3)
  cc.AddCard(card4)

  // Basic test, 4 cards, select 3
  randoms := cc.SelectRandomCards(3)
  assert.Equal(t, 3, len(randoms), "CC size should be 3")

  // Test that nils aren't being included
  cc.KillCard(card1)
  randoms = cc.SelectRandomCards(3)
  assert.Equal(t, 3, len(randoms), "CC size should be 3")

  // Test that if the number of available cards is less then request, it gives you what it can
  cc.RemoveCard(card2)
  randoms = cc.SelectRandomCards(3)
  assert.Equal(t, 2, len(randoms), "CC size should be 2")
}

// GetRandomCard
func Test_CardContainer_GetRandomCard(t *testing.T) {

}