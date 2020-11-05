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
	String() string
}

//Move describes a move that a player makes. Which card is played, and is it discarded.
// If not discarded, then the move is assumed to be scored.
type Move struct {
	C            Card   // Card that is being placed on the table this move.
	PickupChoice string // One of: ["new", "B", "Y", values of game.CardColors]
	Discard      bool   // Is the C(Card) a discard?
}

func (m Move) String() string {
	var playDest string
	if m.Discard {
		playDest = "discard"
	} else {
		playDest = "table"
	}
	return fmt.Sprintf("[%v -> %v, pickup: %v]", m.C, playDest, m.PickupChoice)
}

//CardSet stores the set of cards that a player will score or a set of cards discarded.
// DiscardArea stores the set of cards that have been discarded.
// While playing the game in real-life this sits between both players scoing cards.
type CardSet struct {
	Cards        map[string][]Card
	OrderMatters bool //does the value of the card have to increase.. table vs discard
}

func (d *CardSet) String() string {
	report := ""
	for col, cards := range d.Cards {
		report += fmt.Sprintf(" %v:%v\n", col, cards)
	}
	return report
}

//PlaceCard places that Card into the CardSet
func (d *CardSet) PlaceCard(c Card) {
	d.Cards[c.Col] = append(d.Cards[c.Col], c)
}

//VisibleState is a struct of pointers that Players can use to make descisions about what to play
// It represents part of the game which is publicly visible to the players
// It shows cards that have been revealed.
type VisibleState struct {
	Hand                   []Card
	OpponentHandKnownCards []Card //This is usually empty, but in the case when opponent picks up discard, it will show that card until the opponent plays that card to the table
	Table                  *CardSet
	OpponentTable          *CardSet
	Discards               *CardSet
	DeckCardsLeft          int
}

func (vs VisibleState) String() string {
	return fmt.Sprintf("VisibleState:\nPlayersHand:%v\nTable:\n%vOpTable:\n%vOpHandCardsKnown:%v\nDiscards:\n%vDeckCardsLeft:%v",
		vs.Hand, vs.Table, vs.OpponentTable, vs.OpponentHandKnownCards, vs.Discards, vs.DeckCardsLeft)
}

//Contender tracks all data relevant to a player of the Game.
// This is distinct from Player(which provides the descision making)
type Contender struct {
	ExtPlayer      Player
	Table          CardSet
	Hand           []Card
	DiscardPickups []Card //Contains a record of cards that this player has picked up from the discard but not played yet.
	ID             string
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

//NewGame returns a newly initialized Game object, ready for players to be added
func NewGame() Game {
	g := Game{}
	g.Init()
	return g
}

//Init sets or resets the Game object to a new state.
func (g *Game) Init() {
	g.InitializeDeck()
	g.InitializeStorage()
	g.Turn = 0
	//TODO return all hands table to zero/pristine state
}

//InitializeStorage gets the CardSets ready
func (g *Game) InitializeStorage() {
	g.P1.Table.Cards = make(map[string][]Card)
	g.P2.Table.Cards = make(map[string][]Card)
	g.discardPiles.Cards = make(map[string][]Card)
}

//NextPlayer returns the player who needs to play the next move.
func (g *Game) NextPlayer() Player {
	return g.nextToMove.ExtPlayer
}

//Apply the given Move to the gamestate
func (g *Game) Apply(m Move) {
	//check that player(to play) actually has Card they are wanting to play.
	if !contains(g.nextToMove.Hand, m.C) {
		panic(fmt.Sprintf("Player:%v played a card that is not in their hand, cheater!\n Card played was: %v, the hand contains: %v", g.nextToMove.ExtPlayer, m.C, g.nextToMove.Hand))
	}
	g.nextToMove.Hand = Remove(g.nextToMove.Hand, m.C)
	//Apply move to modify gamestate...
	if m.Discard { //discards are always legal
		g.discardPiles.PlaceCard(m.C)
	} else {
		//check that the play is legal,
		// legal moves involve playing a card of higher value onto a card of lower value.
		targetLocation := g.nextToMove.Table.Cards[m.C.Col]
		validMove := true //if the targetLocation is empty then the move is valid
		if len(targetLocation) > 0 {
			if !m.C.CanStackOn(targetLocation[len(targetLocation)-1]) {
				validMove = false
			}
			if !validMove {
				panic(fmt.Sprintf("Player %v tried to play illegal move. They put card %v, on top of %v.",
					g.nextToMove.ExtPlayer, m.C, targetLocation[len(targetLocation)-1]))
			}
		}
		if validMove {
			g.nextToMove.Table.PlaceCard(m.C)
			//Check if this card was a discarded card that the player picked up.
			// If it was then we can remove it from the KnownCards set, as it will now be represented on the table.
			if contains(g.nextToMove.DiscardPickups, m.C) {
				Remove(g.nextToMove.DiscardPickups, m.C)
			}
		}
	} //Finished updating gamestate
	//Update players hand with new card
	if m.PickupChoice == "new" {
		g.nextToMove.Hand = append(g.nextToMove.Hand, g.Deck[0])
		g.Deck = g.Deck[1:] //Drop card 0 from Deck.
	} else {
		//place the card that has the highest index from the discard[color] into the players hand
		i := len(g.discardPiles.Cards[m.PickupChoice])
		if i == 0 {
			panic(fmt.Sprintf("player %v, tried to pickup from empty discard.", g.nextToMove))
		}
		topCard := g.discardPiles.Cards[m.PickupChoice][i-1]
		g.discardPiles.Cards[m.PickupChoice] = g.discardPiles.Cards[m.PickupChoice][:i-1]
		g.nextToMove.Hand = append(g.nextToMove.Hand, topCard)
		g.nextToMove.DiscardPickups = append(g.nextToMove.DiscardPickups, topCard)
	}
	g.opponent, g.nextToMove = g.nextToMove, g.opponent // Swap active and non-active player
	g.Turn++
}

//AddPlayer sets a player to the given position,
// Currently we only support 2 players in position 1 or 2.
func (g *Game) AddPlayer(p Player, pos int) {
	switch pos {
	case 1:
		g.P1.ExtPlayer = p
		g.P1.Table.OrderMatters = true //Required to be set on initialization of the contender
		g.nextToMove = &g.P1
	case 2:
		g.P2.ExtPlayer = p
		g.P2.Table.OrderMatters = true //Required to be set on initialization of the contender
		g.opponent = &g.P2
	default:
		fmt.Printf("FAILED TO ADD player to position:%v", pos)
	}
}

//Deal is called to setup before the first move.
func (g *Game) Deal() {
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
		Table:                  &(g.nextToMove.Table),
		OpponentTable:          &(g.opponent.Table),
		OpponentHandKnownCards: g.opponent.DiscardPickups,
		Discards:               &(g.discardPiles),
		DeckCardsLeft:          len(g.Deck)}

}
