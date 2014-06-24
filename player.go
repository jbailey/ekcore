package ekcore

import (
  "math"
)

type Player struct {
  Name string
  Level int
}

func (player *Player) MaxHitPoints() int {
  breakingPoint := ((player.Level - 1) / 10 + 1) * 10
  baseInc := 60
  baseHP := 400
  baseHPInc := 200

  for i := 0; i < int(math.Floor(float64(breakingPoint / 10))); i++ {
    baseHP += (i + 3) * baseHPInc
  }

  finalHP := baseHP + (baseInc + breakingPoint) * (player.Level - (1 + breakingPoint - 10))
  
  return finalHP
}

func (player *Player) CostAllowed() int {
  cost := 10
  for i := 0; i < player.Level; i++ {
    if i < 20 {
      cost += 3
    } else if i < 50 {
      cost += 2
    } else {
      cost += 1
    }
  }

  return cost
}