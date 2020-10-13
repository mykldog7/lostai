package game

//This file setups up the card and manages card types.

import (
	"fmt"
	"math/rand"
)

//Card type
type Card struct {
	Val string
	Col string
}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.Col, c.Val)
}

//CardColors contains the valid card colors
var CardColors = []string{"B", "R", "G", "Y", "W"}

//CardValues contains the valid card Values
var CardValues = []string{"2", "3", "4", "5", "6", "7", "8", "9", "H", "H", "H"}

//GiveRandomExampleCard is a utility function, it will present a random card.
// This is done by picking a random colour and value, then presenting that card.
func GiveRandomExampleCard() Card {
	var col string = CardColors[rand.Intn(len(CardColors))]
	var val string = CardValues[rand.Intn(len(CardValues))]
	return Card{Col: col, Val: val}
}

//initalShuffle returns a shuffled []Card (deck) ready to deal and play.
func (g *Game) InitalShuffle(seed int) {
	return
}

//initalizeDeck returns a []Card with one of each card. It is a perfect deck.
func (g *Game) InitalizeDeck() {
	g.Deck = make([]Card, 0, len(CardColors)*len(CardValues)) //one of each value of each color
	for _, col := range CardColors {
		for _, val := range CardValues {
			g.Deck = append(g.Deck, Card{Col: col, Val: val})
		}
	}
	fmt.Println(g.Deck)
}
