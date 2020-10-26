package players

import (
	"fmt"
	"lostcity/pkg/game"
	"math/rand"
)

//Player implementation goes here
/* Implements:
SelectMove(hand []Card, game VisibleState) Move
String() string
*/

//RandomLegalMovePlayer is a player for the LostCity game
// It selects and random card as its move,
// without considering if the move is strategic or smart.
type RandomLegalMovePlayer struct {
	Name string
}

//GetName returns the name of this player
func (p RandomLegalMovePlayer) String() string {
	return fmt.Sprintf("%v[RandomLegalMovePlayer]", p.Name)
}

//SelectMove returns the Move that the AI player is selecting,
// in this case its a random legal move.
func (p RandomLegalMovePlayer) SelectMove(vs game.VisibleState) game.Move {
	rand.Shuffle(len(vs.Hand), func(i, j int) {
		vs.Hand[i], vs.Hand[j] = vs.Hand[j], vs.Hand[i]
	})
	//TODO: IMPLEMENT CHECK FOR LEGAL MOVE
	i := 0
	for ; i < len(vs.Hand); i++ {
		color := vs.Hand[i].Col
		cardsOfColor := len(vs.Table.Cards[color])
		if cardsOfColor == 0 {
			return game.Move{C: vs.Hand[i], Discard: false, PickupChoice: "new"}
		}
		if vs.Hand[i].CanStackOn(vs.Table.Cards[color][len(vs.Table.Cards[color])-1]) {
			return game.Move{C: vs.Hand[i], Discard: false, PickupChoice: "new"}
		}
	}
	//discard if we can't find a legal move to play to the table
	return game.Move{C: vs.Hand[0], Discard: true, PickupChoice: "new"}
}
