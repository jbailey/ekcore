package ekcore

import (
)

type MatchSummary struct {
  Winner *Player
  WinType string
}

func NewMatchSummary() *MatchSummary {
  return new(MatchSummary)
}