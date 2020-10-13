/*Package game is a utility package that contains all the game items
and helpful functions for using them.
It is used to manage the state and scores of any particular LostCity game. */
package game

import "fmt"

//Move describes a move that a player makes. Which card is played, and is it discarded.
// If not discarded, then the move is assumed to be scored.
type Move struct {
	C       Card
	Discard bool
}

//Player provides moves to the game engine.
// The game engine requires a Player to Provide a Move given the gamestate,
// and a selection of cards(representing the available cards to select)
// The Player must select a card and return a move.
type Player interface {
	SelectMove(hand []Card, game VisibleState) Move
}

//CardSet stores the set of cards that a player will score or a set of cards discarded.
// DiscardArea stores the set of cards that have been discarded.
// While playing the game in real-life this sits between both players scoing cards.
type CardSet struct {
	Blue, Red, Yellow, White, Green []Card //maybe better to use a map so that the number of colours can be adjusted.
}

func (d *CardSet) String() string {
	return fmt.Sprintf("B:%v\nR:%v\nY:%v\nW:%v\nG:%v\n", d.Blue, d.Red, d.Yellow, d.White, d.Green)
}

//PlaceCard places that Card into the CardSet
func (d *CardSet) PlaceCard(c Card) {
	switch c.Col {
	case "B":
		d.Blue = append(d.Blue, c)
	case "R":
		d.Red = append(d.Red, c)
	case "G":
		d.Green = append(d.Green, c)
	case "Y":
		d.Yellow = append(d.Yellow, c)
	case "W":
		d.White = append(d.White, c)
	default:
		panic("Unknown color")
	}
}

//VisibleState is a struct of pointers that Players can use to make descisions about what to play
// It represents part of the game which is publicly visible to the players
// It shows cards that have been revealed.
type VisibleState struct {
	Area         *CardSet
	OpponentArea *CardSet
	Discards     *CardSet
}

//Game tracks all the items currently involved in the game
type Game struct {
	p1Table, p2Table CardSet
	discardPiles     CardSet
	Deck             []Card
	p1Hand, p2Hand   []Card
}

func (g *Game) String() string {
	return fmt.Sprintf("P1Hand:%v\nP1Table:%v\nDiscards%v\nP2Table:%v\nP2Hand:%v", g.p1Hand, g.p1Table, g.discardPiles, g.p2Table, g.p2Hand)
}

//Apply the given Move to the gamestate
func (g *Game) Apply(m Move) {
	if m.Discard {
		g.discardPiles.PlaceCard(m.C)
	} else {
		g.p1Table.PlaceCard(m.C)
	}
	return
}
