package ekcore

import (
	"fmt"
  "log"
)

type Match struct {
  Hero            Player
  Opponent        Player
	HeroBoard       PlayerBoard
	OpponentBoard   PlayerBoard
  AttackerBoard   *PlayerBoard
  DefenderBoard   *PlayerBoard
	Round           int
}

func NewMatch(hero Player, opponent Player, heroDeck Deck, opponentDeck Deck) *Match {
  match := new(Match)
  match.Hero = hero
  match.Opponent = opponent
  match.HeroBoard = NewBoard(hero, heroDeck)
  match.OpponentBoard = NewBoard(opponent, opponentDeck)

  return match
}

func (match *Match) Run() {
	fmt.Printf("%v vs. %v\n", match.Hero.Name, match.Opponent.Name)

  match.init()

	for {
		match.Round += 1
    match.Opponent.CurrentHitPoints -= 1000

    match.DoRound()

    if match.isGameOver() == true { break }
	}
	fmt.Printf("Match lasted %d rounds\n", match.Round)
  fmt.Println("Player HP: ", match.Hero.CurrentHitPoints)
  fmt.Println("Opponent HP: ", match.Opponent.CurrentHitPoints)
}

func (match *Match) init() {
  match.Hero.CurrentHitPoints = match.Hero.MaxHitPoints()
  match.Opponent.CurrentHitPoints = match.Opponent.MaxHitPoints()
  match.shuffleDecks()

  match.AttackerBoard = &match.HeroBoard
  match.DefenderBoard = &match.OpponentBoard
}

func (match *Match) DoRound() {
  for idx, card := range match.AttackerBoard.Battlefield {
    if len(match.DefenderBoard.Battlefield) <= idx {
      match.DirectAttack(&card, match.DefenderBoard)
    } else {
      defendingCard := match.DefenderBoard.Battlefield[idx]
      match.AttackCard(&card, &defendingCard)
    }
  }
}

func (match *Match) AttackCard(attackingCard *Card, defendingCard *Card) {
  defendingCard.CurrentHitPoints -= attackingCard.CurrentAttack
  log.Printf("%v hits %v for %d damage", attackingCard.Name, defendingCard.Name, attackingCard.CurrentAttack)
  if defendingCard.CurrentHitPoints < 0 { defendingCard.CurrentHitPoints = 0 }
}

func (match *Match) DirectAttack(attackingCard *Card, defendingBoard *PlayerBoard) {
  defendingBoard.CurrentHitPoints -= attackingCard.CurrentAttack
}

func (match *Match) isGameOver() bool {
  if match.Round > 50 || 
    match.Hero.CurrentHitPoints <= 0 ||
    match.Opponent.CurrentHitPoints <= 0 {
    return true
  }
  return false
}

func (match *Match) shuffleDecks() {

}
