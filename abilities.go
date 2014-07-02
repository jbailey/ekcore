package ekcore

import (
  "log"
  "math/rand"
  "time"
)

// Unimplmented
type UnimplementedAbility struct{}
func (u *UnimplementedAbility) SetLevel(level int) { }
func (u *UnimplementedAbility) Execute(castingCard *Card, playerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  log.Println("UNIMPLMENTED ABILITY")
}

// Fireball
type Fireball struct{ Level int }
func (fb *Fireball) SetLevel(level int) { fb.Level = level }
func (fb *Fireball) Execute(castingCard *Card, playerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  targetCard := defendingBoard.Battlefield.SelectRandomCard()
  if targetCard != nil {
    minDamage := fb.Level * 25
    maxDamage := fb.Level * 50
    damage := rand.Intn(maxDamage - minDamage) + minDamage
    targetCard.TakeDamage(damage)
    log.Printf("%v casts Fireball %d for %d damage to %v, %d HP remaining\n", castingCard.Name, fb.Level, damage, targetCard.Name, targetCard.CurrentHitPoints)
  }
}

// Iceball
type Iceball struct{ Level int }
func (a *Iceball) SetLevel(level int) { a.Level = level }
func (a *Iceball) Execute(castingCard *Card, playerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  targetCard := defendingBoard.Battlefield.SelectRandomCard()
  if targetCard != nil {
    damage := a.Level * 20
    targetCard.TakeDamage(damage)
    var freeze bool
    if Random(1, 100) < 45 { freeze = true } else { freeze = false }
    // TODO: Implement the freeze part
    log.Printf("%v cast Iceball %d for %d damage to %v, %v", castingCard.Name, a.Level, damage, targetCard.Name, targetCard.AttackAndHitPointsString())
    if freeze == true { log.Printf("%v would be frozen if thise was implemented\n", targetCard.Name) }
  }
}

// Thunderbolt
type Thunderbolt struct{ Level int }
func (a *Thunderbolt) SetLevel(level int) { a.Level = level }
func (a *Thunderbolt) Execute(castingCard *Card, playerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  damage := a.Level * 25
  targetCard := defendingBoard.Battlefield.SelectRandomCard()
  if targetCard != nil {
    targetCard.TakeDamage(damage)
    var stunned bool
    if Random(1, 100) < 50 { stunned = true } else { stunned = false }
    // TODO: Implement the stunned part
    log.Printf("%v cast Thunderbolt %d for %d damage to %v, %v", castingCard.Name, a.Level, damage, targetCard.Name, targetCard.AttackAndHitPointsString())
    if stunned == true { log.Printf("%v would be stunned if thise was implemented\n", targetCard.Name) }
  }
}

// Blight
type Blight struct{ Level int }
func (a *Blight) SetLevel(level int) { a.Level = level }
func (a *Blight) Execute(castingCard *Card, playerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  targetCard := defendingBoard.Battlefield.SelectRandomCard()
  if targetCard != nil {
    loss := a.Level * 10
    targetCard.LoseAttack(loss)
    targetCard.TakeDamage(loss)
    log.Printf("%v cast Blight %d, %v loses %d attack and %d hit points, %v", castingCard.Name, a.Level, targetCard.Name, loss, loss, targetCard.AttackAndHitPointsString())
  } else {
    log.Printf("No card to cast Blight %d on\n", a.Level)
  }
}

// Backstab
type Backstab struct{ Level int }
func (a *Backstab) SetLevel(level int) { a.Level = level }
func (a *Backstab) Execute(castingCard *Card, PlayerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  castingCard.CurrentAttack += a.Level * 40
}

// Rejuvenation
type Rejuvenation struct{ Level int }
func (r *Rejuvenation) SetLevel(level int) { r.Level = level }
func (r *Rejuvenation) Execute(castingCard *Card, playerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  healingAmount := 30 * r.Level
  castingCard.Heal(healingAmount)
  log.Printf("Casting level %d Rejuvenation, healed %v for %d hit points", r.Level, castingCard.Name, healingAmount)
}

// Immuniity
type Immunity struct{}
func (i *Immunity) SetLevel(level int) { }
func (i *Immunity) Execute(castingCard *Card, playerBoard *PlayerBoard, defendingBoard *PlayerBoard) {
  log.Println("Immunitity bitch")
}

func init() {
  // Seed the RNG
  rand.Seed(time.Now().UTC().UnixNano())

  // Register each ability
  Register("Unimplemented", func() Ability { return &UnimplementedAbility{} })
  Register("Fireball", func() Ability { return &Fireball{} })
  Register("Iceball", func() Ability { return &Iceball{} })
  Register("Blight", func() Ability { return &Blight{} })
  Register("Rejuvenation", func() Ability { return &Rejuvenation{} })
  Register("Immunity", func() Ability { return &Immunity{} })
}
