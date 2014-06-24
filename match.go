package ekcore

import (
	"fmt"
  "log"
  "math"
)

type Match struct {
  Hero            Player
  Opponent        Player
	HeroBoard       PlayerBoard
	OpponentBoard   PlayerBoard
  AttackerBoard   *PlayerBoard
  DefenderBoard   *PlayerBoard
	Round           int
  Summary         *MatchSummary
}

func NewMatch(hero Player, opponent Player, heroDeck Deck, opponentDeck Deck) *Match {
  match := new(Match)
  match.Hero = hero
  match.Opponent = opponent
  match.HeroBoard = NewPlayerBoard(hero, heroDeck)
  match.OpponentBoard = NewPlayerBoard(opponent, opponentDeck)
  match.Summary = NewMatchSummary()

  return match
}

// init
func (match *Match) init() {
  match.HeroBoard.CurrentHitPoints = match.Hero.MaxHitPoints()
  match.OpponentBoard.CurrentHitPoints = match.Opponent.MaxHitPoints()
  match.shuffleDecks()

  match.AttackerBoard = &match.HeroBoard
  match.DefenderBoard = &match.OpponentBoard
}

// Run
func (match *Match) Run() {
	fmt.Printf("%v vs. %v\n", match.Hero.Name, match.Opponent.Name)

  match.init()

	for {
		match.Round += 1
    // match.Opponent.CurrentHitPoints -= 1000

    match.DoRound()

    if match.isGameOver() == true {
      break
    } else {
      match.AttackerBoard, match.DefenderBoard = match.DefenderBoard, match.AttackerBoard
    }
	}

  match.printMatchSummary()
}

// UpdateHandleWaits
func (match *Match) UpdateHandWaits() {
  match.AttackerBoard.DecrementWaits(1)
  match.DefenderBoard.DecrementWaits(1)
}

// DoRound
func (match *Match) DoRound() {
  fmt.Println("-------------------------------------------")
  fmt.Println("ROUND: ", match.Round)
  fmt.Println("-------------------------------------------")

  // Update all card waits
  match.UpdateHandWaits()

  // Draw a card from the deck
  match.AttackerBoard.DrawCard()

  // Move any cards from the hand to battefield whose timer is up
  match.AttackerBoard.MoveExpiredHandCards()

  // Print out the boards
  match.AttackerBoard.PrintBoard()
  match.DefenderBoard.PrintBoard()


  // Activate Runes

  // Cards attack
  match.CardsAttack()

  fmt.Println("---===  END ROUND  ===---")
}

func (match *Match) CardsAttack() {
  for idx, card := range match.AttackerBoard.Battlefield.Cards {
    if len(match.DefenderBoard.Battlefield.Cards) <= idx {
      match.PlayerAttack(card, match.DefenderBoard)
    } else {
      defendingCard := match.DefenderBoard.Battlefield.Cards[idx]
      cardIsDead := match.AttackCard(card, defendingCard)
      if cardIsDead {
        match.DefenderBoard.KillCard(idx)
      }
    }
  }
}

func (match *Match) AttackCard(attackingCard *Card, defendingCard *Card) (cardIsDead bool) {
  defendingCard.CurrentHitPoints = int(math.Dim(
    float64(defendingCard.CurrentHitPoints),
    float64(attackingCard.CurrentAttack)))
  log.Printf("%v hits %v for %d damage, %d HP remaining", attackingCard.Name, defendingCard.Name, attackingCard.CurrentAttack, defendingCard.CurrentHitPoints)
  if defendingCard.CurrentHitPoints < 0 { defendingCard.CurrentHitPoints = 0 }
  if defendingCard.CurrentHitPoints > 0 {
    cardIsDead = false
  } else {
    cardIsDead = true
  }
  return cardIsDead
}

// Attack the opposing player directly
func (match *Match) PlayerAttack(attackingCard *Card, defendingBoard *PlayerBoard) {
  defendingBoard.CurrentHitPoints = int(math.Dim(
    float64(defendingBoard.CurrentHitPoints), 
    float64(attackingCard.CurrentAttack)))
  fmt.Printf("%v attacks %v doing %d damage, leaving %d HP remaining\n",
    attackingCard.Name, defendingBoard.Player.Name, attackingCard.CurrentAttack, defendingBoard.CurrentHitPoints)
}

func (match *Match) isGameOver() bool {
  gameOver := false

  if match.HeroBoard.CurrentHitPoints <= 0 {
    match.Summary.Winner = &match.Opponent
    match.Summary.WinType = "Player Death"
    gameOver = true
  } else if match.OpponentBoard.CurrentHitPoints <= 0 {
    match.Summary.Winner = &match.Hero
    match.Summary.WinType = "Opponent Death"
    gameOver = true
  } else {
  }
  return gameOver
}

func (match *Match) shuffleDecks() {
  match.HeroBoard.Deck.Shuffle()
  match.OpponentBoard.Deck.Shuffle()
}

func (match *Match) printMatchSummary() {
  fmt.Println("MATCH SUMMARY")
  match.printWinnerLine()
  fmt.Printf("Match lasted %d rounds\n", match.Round)
  fmt.Printf("Player HP: %d/%d\n", match.HeroBoard.CurrentHitPoints, match.Hero.MaxHitPoints())
  fmt.Printf("Opponent HP: %d/%d\n", match.OpponentBoard.CurrentHitPoints, match.Opponent.MaxHitPoints())
}

func (match *Match) printWinnerLine() {
  fmt.Printf("%v wins by %v\n", match.Summary.Winner.Name, match.Summary.WinType)
}
