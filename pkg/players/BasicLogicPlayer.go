package players

import (
	"fmt"
	"lostcity/pkg/game"
)

//Player implementation goes here
/* Implements:
SelectMove(hand []Card, game VisibleState) Move
String() string
*/

//BasicLogicPlayer is a player for the LostCity game
// It implements a basicly logical card to play.
type BasicLogicPlayer struct {
	Name string
}

//SelectMove returns the Move that the AI player is selecting,
func (p BasicLogicPlayer) SelectMove(vs game.VisibleState) game.Move {
	//For each card in hand, determine the gap it will create and
	// play the card with the smallest gap.
	gapSize := make([]int, len(vs.Hand))
	for i, v := range vs.Hand {
		color := v.Col
		cardsOfColor := len(vs.Table.Cards[color])
		if cardsOfColor == 0 {
			gapSize[i] = v.ValNum
		} else {
			gapSize[i] = v.ValNum - vs.Table.Cards[color][len(vs.Table.Cards[color])-1].ValNum
		}

	}
	return game.Move{C: vs.Hand[0], Discard: true, PickupChoice: "new"}
}

func (p BasicLogicPlayer) String() string {
	return fmt.Sprintf("%v[RandomLegalMovePlayer]", p.Name)
}
