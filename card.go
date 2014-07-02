package ekcore

import (
  "fmt"
  "log"
  "math"
)

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
  MaxHitPoints     int
  Level0SkillName   string
  Level0Skill       Ability
  Level5SkillName   string
  Level5Skill       Ability
  Level10SkillName   string
  Level10Skill      Ability
  Level99SkillName   string
}

func NewCard(name string, level int) (Card, bool) {
  card, found := MasterCardList[name]
  if found == true {
    card.Level = level
    card.Init()
    card.setAbility(0, card.Level0SkillName)
    if card.Level >= 5 { card.setAbility(5, card.Level5SkillName) }
    if card.Level >= 10 { card.setAbility(10, card.Level10SkillName) }
  }

  return card, found
}

func (card *Card) Init() {
  card.CurrentAttack = card.CalculateAttack()
  card.MaxHitPoints = card.CalculateHitPoints()
  card.CurrentHitPoints = card.MaxHitPoints
  // TODO: Initialize abilities here
}

func (card *Card) setAbility(lvlSkill int, abilityName string) {
  ability, ok := FindAbility(abilityName)
  if ok == false {
    log.Printf("UNIMPLEMENTED ABILITY FOUND: |%v|\n", abilityName)
    ability, ok = FindAbility("Unimplemented")
  } else {
    log.Printf("%v has ability %v", card.Name, abilityName)
  }

  switch lvlSkill {
    case 0: card.Level0Skill = ability
    case 5: card.Level5Skill = ability
    case 10: card.Level10Skill = ability
  }
}

func (card *Card) AttackAndHitPointsString() string {
  return fmt.Sprintf("[%d/%d]", card.CurrentAttack, card.CurrentHitPoints)
}

func (card *Card) CalculateAttack() int {
  return card.BaseAttack + (card.AttackGainPerLevel * card.Level)
}

func (card *Card) CalculateHitPoints() int {
  return card.BaseHitPoints + (card.HitPointsGainPerLevel * card.Level)
}

func (card *Card) LoseAttack(amount int) (remainingAtk int) {
  card.CurrentHitPoints = int(math.Dim(
    float64(card.CurrentAttack),
    float64(amount)))

  remainingAtk = card.CurrentAttack
  return
}

func (card *Card) TakeDamage(amount int) (remainingHP int) {
  card.CurrentHitPoints = int(math.Dim(
    float64(card.CurrentHitPoints),
    float64(amount)))

  remainingHP = card.CurrentHitPoints
  return
}

func (card *Card) Heal(amount int) {
  card.CurrentHitPoints += amount
  if card.CurrentHitPoints > card.MaxHitPoints { card.CurrentHitPoints = card.MaxHitPoints }
}
