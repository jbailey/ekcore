package ekcore

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

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