package ekcore

import (
)

const (
)

type Effect struct {
  EffectType int
  TargetType int
  AbilityType int // Ice, Fire, etc
  NumOfTargets int
  OccursOnce bool
  Duration int
}