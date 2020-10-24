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
	//fmt.Println("TODO: IMPLEMENT CHECK FOR LEGAL MOVE")
	return game.Move{C: vs.Hand[0], Discard: false, PickupChoice: "new"}
}
