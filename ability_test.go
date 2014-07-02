package ekcore

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_Ability_Execute(t *testing.T) {
  // card := &Card{ Name: "Babbler", Level: 0 }

  // fb1, ok := FindAbility("Fireball 1")
  // if ok {
  //   fb1.Execute(card, nil, nil)
  // }

  // b3, ok := FindAbility("Blight 3")
  // if ok {
  //   b3.Execute(card, nil, nil)
  // }

  // assert.True(t, true, "WTF")
}

func Test_FindAbility(t *testing.T) {
  ability, ok := FindAbility("Fireball 3") 

  assert.True(t, ok, "Fireball ability should be found")
  assert.ObjectsAreEqual(&Fireball{ Level: 3}, ability)
}