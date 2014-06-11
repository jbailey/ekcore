package ekcore

import ()

type Card struct {
	Name             string
  Class            string
  Stars            int
  Cost             int
  Wait             int
	Level            int
  BaseAttack       int
  AttackGainPerLevel int
	CurrentAttack    int
  BaseHitPoints    int
  HitPointsGainPerLevel int
	CurrentHitPoints int
  Level0SkillName   string
  Level5SkillName   string
  Level10SkillName   string
  Level99SkillName   string
}

func NewCard(name string, level int) (Card, bool) {
  card, found := MasterCardList[name]
  if found == true {
    card.Init()
  }

  return card, found
}

func (card *Card) Init() {
  card.CurrentAttack = card.CalculateAttack()
  card.CurrentHitPoints = card.CalculateHitPoints()
  // TODO: Initialize abilities here
}

func (card *Card) CalculateAttack() int {
  return card.BaseAttack + (card.AttackGainPerLevel * card.Level)
}

func (card *Card) CalculateHitPoints() int {
  return card.BaseHitPoints + (card.HitPointsGainPerLevel * card.Level)
}