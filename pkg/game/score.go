package game

//This file contains code for how scores are calculated (in the game)

//Score is used to generate/calculate the score from a CardSet
func (c *CardSet) Score() int {
	//Only ordered CardSets can be scored.
	if !c.OrderMatters {
		panic("Can't score a CardSet where order doesn't matter.")
	}
	return 0
}
