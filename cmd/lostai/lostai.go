package main

import (
	"fmt"
	"lostcity/pkg/game"
	"lostcity/pkg/players"
)

func main() {
	g := &game.Game{}
	g.AddPlayer(players.RandomLegalMovePlayer{Name: "Annie"}, 1)
	g.AddPlayer(players.RandomLegalMovePlayer{Name: "Bobby"}, 2)
	fmt.Println("Starting...")
	//Shuffle the deck
	g.InitalizeDeck()
	fmt.Println("Shuffling...")
	g.Shuffle(0)
	//Deal the hands
	fmt.Println("Dealing...")
	g.Deal()
	fmt.Println("Staring play...")
	for {
		fmt.Printf("Turn %v:\n", g.Turn)
		if len(g.Deck) == 0 {
			fmt.Println("Deck exhausted, Game over.")
		}
		playersMove := g.NextPlayer().SelectMove(g.GetVisibleState())
		g.Apply(playersMove)
		//fmt.Println(g)
	}
	//fmt.Println(g)
	//fmt.Println(game.Deck)
	//fmt.Println(game.GiveRandomExampleCard())
	//g.InitalizeDeck()
}
