package ekcore

import (
)

type PlayerBoard struct {
  Player Player
  CurrentHitPoints   int
  Deck        []Card
  Hand        []Card
  Battlefield []Card
  Cemetery    []Card
}

func NewBoard(player Player, deck Deck) PlayerBoard {
  board := PlayerBoard{ Player: player, Deck: deck.Cards }
  board.CurrentHitPoints = player.MaxHitPoints()

  return board
}