package players

import (
	"fmt"
	"mykldog7/lostai/pkg/game"
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
	gapSizes := make([]int, len(vs.Hand))
	minGap := 99 //Bigger than the biggest possible gap in the game.
	minGapIndx := -1
	for i, v := range vs.Hand {
		color := v.Col
		cardsOfColor := len(vs.Table.Cards[color])
		if cardsOfColor == 0 {
			switch v.Val {
			case "H":
				gapSizes[i] = 0
			default:
				gapSizes[i] = v.ValNum
			}
		} else {
			gapSizes[i] = v.ValNum - vs.Table.Cards[color][len(vs.Table.Cards[color])-1].ValNum
			//fmt.Printf("Card[%v] compared with card[%v], gapSize= %v\n", v, vs.Table.Cards[color][len(vs.Table.Cards[color])-1], gapSizes[i])
		}
		if minGap > gapSizes[i] && gapSizes[i] >= 0 {
			minGap = gapSizes[i]
			minGapIndx = i
		}
	}
	if minGapIndx == -1 {
		//No legally playable card found, discard instead.
		return game.Move{C: vs.Hand[0], Discard: true, PickupChoice: "new"}
	}
	return game.Move{C: vs.Hand[minGapIndx], Discard: false, PickupChoice: "new"}
}

func (p BasicLogicPlayer) String() string {
	return fmt.Sprintf("%v[BasicLogicPlayer]", p.Name)
}
