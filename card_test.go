package ekcore

import (
  "testing"
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