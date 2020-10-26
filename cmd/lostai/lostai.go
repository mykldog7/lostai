package main

import (
	"fmt"
	"lostcity/pkg/game"
	"lostcity/pkg/players"
)

func main() {
	fmt.Println("H" < "6")
	g := game.NewGame()
	g.AddPlayer(players.RandomLegalMovePlayer{Name: "Annie"}, 1)
	g.AddPlayer(players.RandomLegalMovePlayer{Name: "Bobby"}, 2)

	fmt.Println("Starting simulation...")
	fmt.Println("Shuffling...")
	g.Shuffle(0)
	//Deal the hands
	fmt.Println("Dealing...")
	g.Deal()
	fmt.Println("Staring play...")
	for {
		if len(g.Deck) == 0 {
			fmt.Println("Deck exhausted, Game over.")
			break
		}
		fmt.Printf("Turn %v:", g.Turn)
		playersMove := g.NextPlayer().SelectMove(g.GetVisibleState())
		fmt.Printf("%v %v", g.NextPlayer(), playersMove)
		g.Apply(playersMove)
		fmt.Printf("\n")
	}
	fmt.Println("Simulation Completed.")
	fmt.Println("---Scores---")
	p1Score, p2Score := g.Score()
	fmt.Printf("%v Scored: %v\n%v Scored: %v", g.P1, p1Score, g.P2, p2Score)
	//g.InitalizeDeck()
}
