package ekcore

import (
  "bufio"
  "os"
  // "fmt"
  "strings"
)

func LoadDeck(filepath string) *Deck {
  deck := new(Deck)

  // read the file
  file, err := os.Open(filepath)
  if err != nil { panic(err) }
  defer file.Close()

  // read each line not starting with a # into cards until a blank line
  var readingCards = true
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if len(line) == 0 {
      readingCards = false
    } else if strings.Index(line, "#") != 0 {
      if readingCards {
        cardName := line
        card := Card{ Name: cardName, Level: 10 }
        deck.AddCard(card)
        // fmt.Println("Card: ", card.Name)
        // card = Card.FindByName(cardName)
      } else {
        runeName := line
        rune := Rune{ Name: runeName, Level: 10 }
        deck.AddRune(rune)
        // fmt.Println("Rune: ", rune.Name)
        // rune = Rune.FindByName(runeName)
      }
    }
  }

  return deck
}