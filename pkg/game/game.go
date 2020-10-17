/*Package game is a utility package that contains all the game items
and helpful functions for using them.
It is used to manage the state and scores of any particular LostCity game. */
package game

import "fmt"

//Player provides moves to the game engine.
// The game engine requires a Player to Provide a Move given the gamestate,
// and a selection of cards(representing the available cards to select)
// The Player must select a card and return a move.
type Player interface {
	SelectMove(game VisibleState) Move
	GetName() string
}

//Move describes a move that a player makes. Which card is played, and is it discarded.
// If not discarded, then the move is assumed to be scored.
type Move struct {
	C            Card   // Card that is being placed on the table this move.
	PickupChoice string // One of: ["new", "blue", "yellow", etc]
	Discard      bool   // Is the C(Card) a discard?
}

//CardSet stores the set of cards that a player will score or a set of cards discarded.
// DiscardArea stores the set of cards that have been discarded.
// While playing the game in real-life this sits between both players scoing cards.
type CardSet struct {
	Blue, Red, Yellow, White, Green []Card //maybe better to use a map so that the number of colours can be adjusted.
	OrderMatters                    bool   //does the value of the card have to increase.. table vs discard
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
	Hand          []Card
	Table         CardSet
	OpponentTable CardSet
	Discards      CardSet
	DeckCardsLeft int
}

//Contender tracks all data relevant to a player of the Game.
// This is distinct from Player(which provides the descision making)
type Contender struct {
	player Player
	Table  CardSet
	Hand   []Card
	ID     string
}

//Game tracks all the items currently involved in the game
type Game struct {
	P1, P2       Contender
	discardPiles CardSet
	Deck         []Card
	nextToMove   *Contender
	opponent     *Contender
	Turn         int
}

func (g *Game) String() string {
	deck := "Deck:" + fmt.Sprint(g.Deck)
	return deck + fmt.Sprintf("\nGamestate:\nP1ID:%v\nP1Hand:%v\nP1Table:%v\nDiscards%v\nP2Table:%v\nP2Hand:%v\nP2ID:%v\n",
		g.P1.ID, g.P1.Hand, g.P1.Table, g.discardPiles, g.P2.Table, g.P2.Hand, g.P2.ID)
}

//NextPlayer returns the player who needs to play the next move.
func (g *Game) NextPlayer() Player {
	return g.nextToMove.player
}

//Apply the given Move to the gamestate
func (g *Game) Apply(m Move) {
	//check that player(to play) actually has Card they are wanting to play.
	//TODO if g.nextToMove.Hand
	if m.Discard {
		g.discardPiles.PlaceCard(m.C)
	} else {
		g.P1.Table.PlaceCard(m.C)
	}
	if m.PickupChoice == "new" {
		g.nextToMove.Hand = append(g.nextToMove.Hand, g.Deck[0])
		g.Deck = g.Deck[1:] //Drop card 0 from Deck.
	}
	g.opponent, g.nextToMove = g.nextToMove, g.opponent // Swap active and non-active player
	g.Turn++
}

//AddPlayer sets a player to the given position,
// Currently we only support 2 players in position 1 or 2.
func (g *Game) AddPlayer(p Player, pos int) {
	switch pos {
	case 1:
		g.P1.player = p
	case 2:
		g.P2.player = p
	default:
		fmt.Printf("FAILED TO ADD player to position:%v", pos)
	}
}

//Deal is called to setup before the first move.
func (g *Game) Deal() {
	g.nextToMove = &g.P1
	g.opponent = &g.P2
	for i := 0; i < 8; i++ {
		//Take two cards from the top of the deck..
		c1, c2 := g.Deck[0], g.Deck[1]
		//Remove first two elements from Deck
		g.Deck = g.Deck[2:]
		g.P1.Hand = append(g.P1.Hand, c1)
		g.P2.Hand = append(g.P2.Hand, c2)
	}
}

//GetVisibleState returns the dataset that the next agent is allowed to use
func (g *Game) GetVisibleState() VisibleState {
	return VisibleState{Hand: g.nextToMove.Hand,
		Table:         g.nextToMove.Table,
		OpponentTable: g.opponent.Table,
		Discards:      g.discardPiles,
		DeckCardsLeft: len(g.Deck)}
}
