package ekcore

import (
  "math/rand"
)

func Random(min int, max int) int {
  num := rand.Intn(max - min) + min
  return num
}