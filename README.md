# LostAI
Ever wanted to make the best plays when in a tricky situation involving the LostCities Card game. [About the Game](https://en.wikipedia.org/wiki/Lost_Cities) Now you can. 

This project is written in Go.  To use the code, install go, and clone the repo to your GOPATH/src folder. This can be done automagically with `go get github.com/mykldog7/lostai`. 

You should now have `single` and `lostai` binaries built into your GOPATH/bin location. You can also run a single source file(if its has `package main` in it) by navigating to it and `go run lostai.go`.

2 binaries can be built, but they are not yet parameterized: 
- `lostai` used to simulate a game and review how an AI performs (IN PROGRESS)
- `single` used to request an AI to make a move from a given state (WORKING)

Currently I've got 2 AIs in the repo:
- `RandomLegalMove` will play the first legal card in a random ordering of their hand
- `BasicLogiclayer` will play the card which results in the smallest gap. IE it will minimize the number of cards excluded by a particular play. It doesn't consider the oppoenets tabled cards (V2, will consider more)

### Goals
Create an AI that calculates the best card to play, and where to play it based on all available information.