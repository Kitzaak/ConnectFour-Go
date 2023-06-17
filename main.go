package main

import (
	"connect_four/board"
	"connect_four/game"
	"fmt"
)

func Hello() string {
	return "Hello, world!"
}

func main() {
	var newPieceCol int
	var gameLogic = game.New()

	fmt.Println(Hello())
	fmt.Println(board.DisplayPieces(gameLogic.CurrentState()))
	fmt.Print("Enter a col for your next piece: ")
	fmt.Scanln(&newPieceCol)

	for newPieceCol != 0 {
		gameLogic.AddPiece(newPieceCol)
		fmt.Println(board.DisplayPieces(gameLogic.CurrentState()))
		fmt.Print("Enter a col for your next piece: ")
		fmt.Scanln(&newPieceCol)
	}
}
