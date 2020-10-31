package single

import (
	"fmt"
	"lostcity/pkg/game"
	"lostcity/pkg/players"
)

// VisibleState{Hand: g.nextToMove.Hand,
//		Table:         &(g.nextToMove.Table),
//		OpponentTable: &(g.opponent.Table),
//		Discards:      &(g.discardPiles),
//		DeckCardsLeft: len(g.Deck)}

func main() {
	//Setup VisibleState objects, ready to pass to player instance
	hand := make([]game.Card, 8)
	table := game.CardSet{Cards: make(map[string][]game.Card), OrderMatters: true}
	opponentTable := game.CardSet{Cards: make(map[string][]game.Card), OrderMatters: true}
	discards := game.CardSet{Cards: make(map[string][]game.Card), OrderMatters: false}
	deckCardsLeft := 12

	//Populate cards as wanted
	hand = append(hand, game.NewCard("BH"))
	hand = append(hand, game.NewCard("B2"))
	hand = append(hand, game.NewCard("B5"))
	hand = append(hand, game.NewCard("B7"))
	hand = append(hand, game.NewCard("Y8"))
	hand = append(hand, game.NewCard("R7"))
	hand = append(hand, game.NewCard("WH"))
	hand = append(hand, game.NewCard("G3"))
	table.Cards["B"] = append(table.Cards["B"], game.NewCard("BH"), game.NewCard("B3"))
	table.Cards["Y"] = append(table.Cards["Y"], game.NewCard("Y4"))
	table.Cards["R"] = append(table.Cards["R"], game.NewCard("R5"), game.NewCard("R6"))
	table.Cards["W"] = append(table.Cards["W"], game.NewCard("W2"))

	//Assemble VisibleState Object

	//Get player instance
	player := players.BasicLogicPlayer{Name: "Single"}
	fmt.Printf("Single invocation of: %v\n", player)

	//Request move from player

	//Output
}
