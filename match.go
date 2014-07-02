package ekcore

import (
	"fmt"
  "log"
  "math"
  "bufio"
  "os"
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
  InteractiveMode bool
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
  match.InteractiveMode = false
}

// Run
func (match *Match) Run() {
  bio := bufio.NewReader(os.Stdin)
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

    if match.InteractiveMode {
      fmt.Println("Hit Enter to continue, or q to quit")
      input, _, _ := bio.ReadLine()
      if string(input[:]) == "q" { return }
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
  fmt.Printf("ROUND: %d, %v Attacking\n", match.Round, match.AttackerBoard.Player.Name)
  fmt.Println("-------------------------------------------")

  // Update all card waits
  match.UpdateHandWaits()

  // Draw a card from the deck
  match.AttackerBoard.DrawCard()

  // Move any cards from the hand to battefield whose timer is up
  match.AttackerBoard.MoveExpiredHandCards()

  // Print out the boards
  fmt.Printf("\n<-- BOARDS -->\n")
  fmt.Println("---------------------------------------------")
  match.HeroBoard.PrintBoard()
  fmt.Println("---------------------------------------------")
  match.OpponentBoard.PrintBoard()
  fmt.Println("---------------------------------------------")
  fmt.Println()


  // Activate Runes

  // Cards attack
  match.CardsAttack()

  fmt.Println("---===  END ROUND  ===---")
  fmt.Printf("\n\n")
}

func (match *Match) CardsAttack() {
  for idx, card := range match.AttackerBoard.Battlefield.Cards {
    attackingSlot := match.AttackerBoard.Battlefield.GetSlotNumForCard(card)

    // Execute lvl 0 ability
    if card.Level0Skill != nil {
      card.Level0Skill.Execute(card, match.AttackerBoard, match.DefenderBoard)
    } else {
      log.Printf("%v doesn't have a level 0 ability\n", card.Name)
    }

    defendingCard := match.DefenderBoard.Battlefield.CardAtSlot(attackingSlot)
    fmt.Printf("%v attacking slot %d", card.Name, attackingSlot)
    if defendingCard != nil {
      fmt.Printf(", slot %d is occupied by %v\n", attackingSlot, defendingCard.Name)
    } else {
      fmt.Printf(", slot %d is empty, attacking player\n", attackingSlot)
    }
    if defendingCard == nil {
      match.PlayerAttack(card, match.DefenderBoard)
    } else {
      cardIsDead := match.AttackCard(card, defendingCard)
      if cardIsDead {
        match.DefenderBoard.KillCard(idx)
      }
    }
  }
  match.DefenderBoard.Battlefield.Collapse()
}

func (match *Match) AttackCard(attackingCard *Card, defendingCard *Card) (cardIsDead bool) {
  defendingCard.TakeDamage(attackingCard.CurrentAttack)
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
  log.Printf("%v attacks %v doing %d damage, leaving %d HP remaining\n",
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
