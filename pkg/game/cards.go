package game

//This file setups up the card and manages card types.

import (
	"fmt"
	"math/rand"
	"strconv"
)

//Card type
type Card struct {
	ValNum int    //used to score should match the Val, field
	Val    string //shows the facevalue of the card
	Col    string //shows the color of the card
}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.Col, c.Val)
}

//CardColors contains the valid card colors
var CardColors = []string{"B", "R", "G", "Y", "W"}

//CardValues contains the valid card Values
var CardValues = []string{"2", "3", "4", "5", "6", "7", "8", "9", "X", "H", "H", "H"}

//CanStackOn returns true of the card can stack on top of the Card passed as argument.
func (c Card) CanStackOn(p Card) bool { //6<H
	return c.Col == p.Col && (((p.Val < c.Val) && (c.Val != "H") || (p.Val == "H" && c.Val != "H")) || p.Val == "H" && c.Val == "H")
}

//NewCard makes a Card out of a string
// the input string should be length 2 and have form/[B|R|G|Y|W][2..9|X|H]/
func NewCard(in string) Card {
	//Check that the color indicator is in the list of CardColors
	validColor := false
	for _, v := range CardColors {
		if string(in[0]) == v {
			validColor = true
			break //stop searching
		}
	}
	//If its not then stop.
	if !validColor {
		panic("Tried to make a card with an invalid color.")
	}
	//Everything seems fine(value errors handled in valToNum), make the card...
	return Card{Col: string(in[0]), Val: string(in[1]), ValNum: valToNum(string(in[1]))}
}

//utility function to convert a string value for a card to a number value
func valToNum(val string) int {
	valnum := 0
	var err error
	switch val {
	case "H":
		valnum = -1
	case "X":
		valnum = 10
	default:
		valnum, err = strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("error converting %v to int", val))
		}
	}
	return valnum
}

//GiveRandomExampleCard is a utility function, it will present a random card.
// This is done by picking a random colour and value, then presenting that card.
func GiveRandomExampleCard() Card {
	var col string = CardColors[rand.Intn(len(CardColors))]
	var val string = CardValues[rand.Intn(len(CardValues))]
	return NewCard(col + val)
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

//InitializeDeck returns a []Card with one of each card. It is a perfect deck.
func (g *Game) InitializeDeck() {
	g.Deck = make([]Card, 0, len(CardColors)*len(CardValues)) //one of each value of each color
	for _, col := range CardColors {
		for _, val := range CardValues {
			g.Deck = append(g.Deck, NewCard(col+val))
		}
	}
}

//Remove is used to take a card out of a card set. (remove a card from a players hand)
func Remove(cs []Card, c Card) []Card {
	for i := 0; i < len(cs); i++ {
		if cs[i].Col == c.Col && cs[i].Val == c.Val {
			cs[i] = cs[len(cs)-1]
			//cs[len(cs)-1] = nil
			cs = cs[:len(cs)-1]
			return cs
		}
	}
	panic("Tried to remove a card that wasn't in a hand.")
}

//Contains checks if the slice of cards contains a card matching the given card, if it does it returns true.
func contains(cs []Card, c Card) bool {
	for _, v := range cs {
		if v.Col == c.Col && v.Val == c.Val {
			return true
		}
	}
	//not found at all? then return false
	return false
}
