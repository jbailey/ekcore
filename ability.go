package ekcore

import (
  // "log"
  "sync"
  "strings"
  "strconv"
  "regexp"
)

type Ability interface {
  SetLevel(level int)
  Execute(castingCard *Card, playerBoard *PlayerBoard, 
    defendingBoard *PlayerBoard)
}

func FindAbility(name string) (ability Ability, ok bool) {
  ok = false
  re := regexp.MustCompile("^[^0-9]*")
  baseName := re.FindString(name)
  baseName = strings.Trim(baseName, " \n")

  ability = getAbility(baseName)
  if ability != nil {
    ok = true
    re = regexp.MustCompile("[0-9]*$")
    strLevel := re.FindString(name)
    if strLevel != "" {
      level, _ := strconv.Atoi(strLevel)
      ability.SetLevel(level)
    }
  }
  return ability, ok
}

var abilities = struct {
  m map[string]AbilityCtor
  sync.RWMutex
}{m: make(map[string]AbilityCtor)}

type AbilityCtor func() Ability

func Register(id string, newfunc AbilityCtor) {
  abilities.Lock()
  abilities.m[id] = newfunc
  abilities.Unlock()
}

func getAbility(id string) (a Ability) {
  abilities.RLock()
  ctor, ok := abilities.m[id]
  abilities.RUnlock()
  if ok {
    a = ctor()
  }
  return
}
