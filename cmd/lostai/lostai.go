package main

import (
	"fmt"

	"github.com/mykldog7/lostai/pkg/game"
	"github.com/mykldog7/lostai/pkg/players"
)

func main() {
	g := game.NewGame()
	g.AddPlayer(players.RandomLegalMovePlayer{Name: "Annie"}, 1)
	g.AddPlayer(players.BasicLogicPlayer{Name: "Bobby"}, 2)

	fmt.Println("Starting simulation...")
	fmt.Println("Shuffling...")
	g.Shuffle(2)
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
	fmt.Println("---Scores---")
	p1Score, p2Score := g.Score()
	fmt.Printf("%v Scored: %v\n%v Scored: %v\n", g.P1.ExtPlayer, p1Score, g.P2.ExtPlayer, p2Score)
}
