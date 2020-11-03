package main

import (
	"fmt"

	"github.com/mykldog7/lostai/pkg/game"
	"github.com/mykldog7/lostai/pkg/players"
)

// VisibleState{Hand: g.nextToMove.Hand,
//		Table:         &(g.nextToMove.Table),
//		OpponentTable: &(g.opponent.Table),
//		Discards:      &(g.discardPiles),
//		DeckCardsLeft: len(g.Deck)}

func main() {
	//Setup VisibleState objects, ready to pass to player instance
	hand := make([]game.Card, 0)
	table := game.CardSet{Cards: make(map[string][]game.Card), OrderMatters: true}
	opponentTable := game.CardSet{Cards: make(map[string][]game.Card), OrderMatters: true}
	discards := game.CardSet{Cards: make(map[string][]game.Card), OrderMatters: false}

	//Cards already removed from deck(total 12 * 5 = 60)
	//15(tabled) + 16(players' hands) = 31
	deckCardsLeft := 29

	//Populate cards as wanted
	hand = append(hand, game.NewCard("BH"), game.NewCard("B2"),
		game.NewCard("B5"), game.NewCard("B7"),
		game.NewCard("Y8"), game.NewCard("R7"),
		game.NewCard("WH"), game.NewCard("G4"))

	table.Cards["B"] = append(table.Cards["B"], game.NewCard("BH"), game.NewCard("B3"))
	table.Cards["Y"] = append(table.Cards["Y"], game.NewCard("Y5"))
	table.Cards["R"] = append(table.Cards["R"], game.NewCard("R5"), game.NewCard("R6"))
	table.Cards["W"] = append(table.Cards["W"], game.NewCard("W2"))

	opponentTable.Cards["Y"] = append(opponentTable.Cards["Y"], game.NewCard("Y2"), game.NewCard("Y3"), game.NewCard("Y4"))
	opponentTable.Cards["R"] = append(opponentTable.Cards["R"], game.NewCard("R3"), game.NewCard("R4"))
	opponentTable.Cards["W"] = append(opponentTable.Cards["W"], game.NewCard("WH"))
	opponentTable.Cards["G"] = append(opponentTable.Cards["G"], game.NewCard("G3"))

	discards.Cards["G"] = append(discards.Cards["G"], game.NewCard("G2"))
	discards.Cards["R"] = append(discards.Cards["R"], game.NewCard("R2"))

	//Assemble VisibleState Object
	vs := game.VisibleState{Hand: hand,
		Table:                  &table,
		OpponentTable:          &opponentTable,
		OpponentHandKnownCards: make([]game.Card, 0),
		Discards:               &discards,
		DeckCardsLeft:          deckCardsLeft,
	}

	fmt.Println(vs)
	//Get player instance
	player := players.BasicLogicPlayer{Name: "Single"}

	//Request move from player
	move := player.SelectMove(vs)

	//Output
	fmt.Printf("%v has hand: %v\n", player, hand)
	fmt.Printf("%v, plays move %v\n", player, move)
}
