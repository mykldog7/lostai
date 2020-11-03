package players

import (
	"fmt"
	"math/rand"

	"github.com/mykldog7/lostai/pkg/game"
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

//SelectMove returns the Move that the AI player is selecting,
// in this case its a random legal move.
func (p RandomLegalMovePlayer) SelectMove(Vs game.VisibleState) game.Move {
	//Shuffle the hand
	rand.Shuffle(len(Vs.Hand), func(i, j int) {
		Vs.Hand[i], Vs.Hand[j] = Vs.Hand[j], Vs.Hand[i]
	})
	//Search through cards in hand.. play the first legal one. If none are legal discard the first one.
	i := 0
	for ; i < len(Vs.Hand); i++ {
		color := Vs.Hand[i].Col
		cardsOfColor := len(Vs.Table.Cards[color])
		if cardsOfColor == 0 {
			fmt.Println(Vs.Hand[i])
			return game.Move{C: Vs.Hand[i], Discard: false, PickupChoice: "new"}
		}
		if Vs.Hand[i].CanStackOn(Vs.Table.Cards[color][len(Vs.Table.Cards[color])-1]) {
			fmt.Println(Vs.Hand[i])
			return game.Move{C: Vs.Hand[i], Discard: false, PickupChoice: "new"}
		}
	}
	//discard if we can't find a legal move to play to the table
	return game.Move{C: Vs.Hand[0], Discard: true, PickupChoice: "new"}
}

func (p RandomLegalMovePlayer) String() string {
	return fmt.Sprintf("%v[RandomLegalMovePlayer]", p.Name)
}
