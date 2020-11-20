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
type CalculatingSingleMoveAIPlayer struct {
	Name string
}

func (p CalculatingSingleMoveAIPlayer) String() string {
	return fmt.Sprintf("%v[CalculatingAIPlayer]", p.Name)
}

//SelectMove returns the Move that the AI player is selecting,
func (p CalculatingSingleMoveAIPlayer) SelectMove(vs game.VisibleState) game.Move {
	//For each card in hand...
	var bestCard game.Card
	var bestExpectation = ((-20 * 4) * 5) //-400 //worst possible score all the investors(hands) but nothing else.
	for _, c := range vs.Hand {
		expectation := expectedFinalScore(vs, c, bestExpectation)
		fmt.Println("expectation", expectation, c)
		if expectation > bestExpectation {
			bestCard = c
			bestExpectation = expectation
		}
	}
	return game.Move{C: bestCard, Discard: false, PickupChoice: "new"}
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
func expectedFinalScore(vs game.VisibleState, c game.Card, bestScoreSoFar int) int {
	//
	stack := *(vs.Table[c.Col])
	fmt.Println("Expected score of stack is: %v", stack.Score)
	return c.ValNum
}
