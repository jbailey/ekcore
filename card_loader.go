package ekcore

import (
  "os"
  "io"
  "log"
  "encoding/csv"
  "strconv"
  "strings"
)

var MasterCardList = make(map[string]Card)

func FindCardByName(name string) (Card, bool) {
   card, found := MasterCardList[name]
   return card, found
}

func LoadCards(cardFile string) {
  file, err := os.Open(cardFile)
  if err != nil { log.Fatalf("Error reading all lines: ", err) }
  defer file.Close()

  reader := csv.NewReader(file)
  reader.Comma = '\t'

  for {
    line, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      log.Print(err)
      os.Exit(-1)
    }

    card := parseLine(line)
    MasterCardList[card.Name] = card
  }
}

func parseLine(line []string) Card {
  var name = strings.Trim(line[0], " \n")
  var class = line[1]
  var stars, _ = strconv.Atoi(line[2])
  var cost, _ = strconv.Atoi(line[3])
  var wait, _ = strconv.Atoi(line[4])
  var lvl0SkillName = line[5]
  var lvl5SkillName = line[6]
  var lvl10SkillName = line[7]
  var lvl99SkillName = line[8]
  var attProgression = line[25]
  var hpProgression = line[26] 

  baseAtk, atkPerLevel := CalcAttack(attProgression)
  baseHP, hpPerLevel := CalcHP(hpProgression)

  // lvl0Skill := ekcore.FindSkillByName(lvl0SkillName)
  // lvl5Skill := ekcore.FindSkillByName(lvl5SkillName)
  // lvl10Skill := ekcore.FindSkillByName(lvl10SkillName)

  card := Card {
    Name: name,
    Class: class,
    Stars: stars,
    Cost: cost,
    Wait: wait,
    BaseAttack: baseAtk,
    AttackGainPerLevel: atkPerLevel,
    BaseHitPoints: baseHP,
    HitPointsGainPerLevel: hpPerLevel,
    Level0SkillName: lvl0SkillName,
    Level5SkillName: lvl5SkillName,
    Level10SkillName: lvl10SkillName,
    Level99SkillName: lvl99SkillName,
  }

  return card
}

func CalcAttack(prog string) (int, int) {
  base, perLevel := 0, 0
  var atks = strings.Split(prog, ",")
  if len(atks) > 2 {
    base, _ = strconv.Atoi(atks[0])
    second, _ := strconv.Atoi(atks[1])
    perLevel = second - base
  }

  return base, perLevel
}

func CalcHP(prog string) (int, int) {
  base, perLevel := 0, 0
  var hps = strings.Split(prog, ",")
  if len(hps) > 2 {
    base, _ = strconv.Atoi(hps[0])
    second, _ := strconv.Atoi(hps[1])
    perLevel = second - base
  }

  return base, perLevel
}