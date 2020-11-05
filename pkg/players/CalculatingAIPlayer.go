package players

import (
	"fmt"

	"github.com/mykldog7/lostai/pkg/game"
)

//Player implementation goes in this file
/* Implements:
SelectMove(hand []Card, game VisibleState) Move
String() string
*/

//CalculatingAIPlayer is a player for the LostCity game
// It calculates the best card to play
type CalculatingAIPlayer struct {
	Name string
}

//SelectMove returns the Move that the AI player is selecting,
func (p CalculatingAIPlayer) SelectMove(vs game.VisibleState) game.Move {
	//For each card in hand...
	var bestCard game.Card
	var bestExpectation = -400 //worst possible score all the investors(hands) but nothing else.
	for _, v := range vs.Hand {
		expectation := expectedFinalScore(v, bestExpectation)
		fmt.Println("expectation", expectation, v)
		if expectation > bestExpectation {
			bestCard = v
			bestExpectation = expectation
		}
	}
	return game.Move{C: bestCard, Discard: false, PickupChoice: "new"}
}

func (p CalculatingAIPlayer) String() string {
	return fmt.Sprintf("%v[CalculatingAIPlayer]", p.Name)
}

//Utility function which returns the expected final score if this card is played
// The expected final score can be determined by:
// Determine the score of the current state(use game.CardSet.Score method)
// Add the points from this current card played to table
// For each remaining card possibly still in the deck(remove visible cards, and cards known to be in opponent hand),
//  add that card(but discount its score by the likelihood its pickedup, and playable)
//  if the other card is already in the players hand(and enough turns remain, its 100%)
// TODO manage edgecases:
//  cards that could be picked up from discards and then scored
//  have a clear/consise method for managing prioirty of card play as the end of the game arrives.
func expectedFinalScore(c game.Card, bestScoreSoFar int) int {
	//
	return c.ValNum
}
