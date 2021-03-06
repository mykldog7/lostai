package game

//This file contains code for how scores are calculated (in the game)

import "fmt"

func ScoreStack(stack []Card) int {
	stackTotal := 0
	// zero cards means exit early
	if len(stack) == 0 {
		return stackTotal
	}
	//cost of starting a color, we have more than zero cards
	stackTotal -= 20
	//how many 'H' cards determines multiplier? [x2, x3, x4]
	multiplier := 1
	//For each card in the stack...
	for _, c := range stack {
		//'H' increases multiplier (investor card)
		if c.Val == "H" {
			multiplier++
		}
		//sum the value of the non-multiplier cards
		if c.Val != "H" {
			stackTotal += c.ValNum
		}
	}
	//apply multiplier to subtotal
	stackTotal *= multiplier
	//did they stack have 8 or more.. bonus 20
	if len(stack) >= 8 {
		stackTotal += 20
	}
	return stackTotal
}

//Score is used to generate/calculate the score from a CardSet
func (cs *CardSet) Score() int {
	//Only ordered CardSets can be scored.
	if !cs.OrderMatters {
		panic("Can't score a CardSet where order doesn't matter.")
	}
	total := 0
	for _, stack := range cs.Cards {
		total += ScoreStack(stack)
	}
	fmt.Printf("TableTotal: %v\n\n", total)
	return total
}

//Score is used to show the scores of the current game state
func (g *Game) Score() (int, int) {
	return g.P1.Table.Score(), g.P2.Table.Score()
}
