package ekcore

import (
  "testing"
)

func Test_MaxHitPoints(t *testing.T) {
  player := Player{ Name: "The Player", Level: 1 }
  testLevel(1000, player.MaxHitPoints(), t)
  player.Level = 10
  testLevel(1630, player.MaxHitPoints(), t)
  player.Level = 11
  testLevel(1800, player.MaxHitPoints(), t)
  player.Level = 21
  testLevel(2800, player.MaxHitPoints(), t)
  player.Level = 60
  testLevel(8080, player.MaxHitPoints(), t)
  player.Level = 61
  testLevel(8800, player.MaxHitPoints(), t)
}

func testLevel(expected int, got int, t *testing.T) {
  if expected != got {
    t.Errorf("Expected MaxHitPoints to be %v, got %v", expected, got)
  }
}

func Test_CostAllowed(t *testing.T) {
  player := Player{ Name: "The Player", Level: 1 }
  testCostAllowed(13, player.CostAllowed(), t)
  player.Level = 20
  testCostAllowed(70, player.CostAllowed(), t)
  player.Level = 21
  testCostAllowed(72, player.CostAllowed(), t)
  player.Level = 49
  testCostAllowed(128, player.CostAllowed(), t)
  player.Level = 50
  testCostAllowed(130, player.CostAllowed(), t)
}

func testCostAllowed(expected int, got int, t *testing.T) {
  if expected != got {
    t.Errorf("Expected MaxAllowed to be %v, got %v", expected, got)
  }
}