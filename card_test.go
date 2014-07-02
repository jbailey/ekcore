package ekcore

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_CalculateAttack(t *testing.T) {
  card := Card{
    Name: "Cerato",
    Level: 1,
    BaseAttack: 245,
    AttackGainPerLevel: 25,
  }

  expected, got := 270, card.CalculateAttack()
  if expected != got {
    t.Errorf("Expected CalculateAttack() to be %v, got %v", expected, got)
  }

  card.Level = 10  
  expected, got = 495, card.CalculateAttack()
  if expected != got {
    t.Errorf("Expected CalculateAttack() to be %v, got %v", expected, got)
  }
}

func Test_CalculateHitPoints(t *testing.T) {
  card := Card{
    Name: "Cerato",
    Level: 1,
    BaseHitPoints: 900,
    HitPointsGainPerLevel: 32,
  }

  expected, got := 932, card.CalculateHitPoints()
  if expected != got {
    t.Errorf("Expected CalculateHitPoints() to be %v, got %v", expected, got)
  }

  card.Level = 10  
  expected, got = 1220, card.CalculateHitPoints()
  if expected != got {
    t.Errorf("Expected CalculateHitPoints() to be %v, got %v", expected, got)
  }
}

func Test_Card_TakeDamage(t *testing.T) {
  card := &Card{ Name: "Test Card", Level: 3, CurrentHitPoints: 300 }

  card.TakeDamage(200)
  assert.Equal(t, 100, card.CurrentHitPoints, "Card should have 100 hit points remaining")
}

func Test_Card_Heal(t *testing.T) {
  card := &Card{ Name: "Test Card", Level: 3, CurrentHitPoints: 200, MaxHitPoints: 300 }

  card.Heal(100)
  assert.Equal(t, 300, card.CurrentHitPoints)

  card.CurrentHitPoints = 200
  card.Heal(200)
  assert.Equal(t, 300, card.CurrentHitPoints, "Heal shouldn't over-heal")
}