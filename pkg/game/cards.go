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

//GiveRandomExampleCardset is a utility function, it will present a random cardset.
// count provides the number of random cards to return
func GiveRandomExampleCardset(count int) []Card {
	c := make([]Card, 0, count)
	for a := 0; a < count; a++ {
		c = append(c, GiveRandomExampleCard())
	}
	return c
}

//Shuffle returns a shuffled []Card (deck) ready to deal and play.
func (g *Game) Shuffle(seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(g.Deck), func(i, j int) {
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	})
}

//InitalizeDeck returns a []Card with one of each card. It is a perfect deck.
func (g *Game) InitalizeDeck() {
	g.Deck = make([]Card, 0, len(CardColors)*len(CardValues)) //one of each value of each color
	for _, col := range CardColors {
		for _, val := range CardValues {
			g.Deck = append(g.Deck, Card{Col: col, Val: val})
		}
	}
}

//Remove is used to take a card out of a card set.
func Remove(cs []Card, c Card) []Card {
	for i := 0; i < len(cs); i++ {
		if cs[i].Col == c.Col && cs[i].Val == c.Val {
			cs[i] = cs[len(cs)-1]
			//cs[len(cs)-1] = nil
			cs = cs[:len(cs)-1]
			break
		}
	}
	return cs
}

//Add is used to add a card to a cardset, usually a hand recieving a Card after pickup
func Add(cs []Card, c Card) []Card {
	return append(cs, c)
}
