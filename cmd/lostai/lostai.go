package main

import (
	"fmt"
	"lostcity/pkg/game"
)

func main() {
	g := &game.Game{}
	fmt.Println("Hello, Michael!")
	m := game.Move{C: game.GiveRandomExampleCard(), Discard: true}
	fmt.Println(g)
	g.Apply(m)
	fmt.Println(g)
	//fmt.Println(game.Deck)
	//fmt.Println(game.GiveRandomExampleCard())
	g.InitalizeDeck()
}
