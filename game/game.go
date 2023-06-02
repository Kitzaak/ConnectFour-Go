package game

import (
	"errors"
	"strconv"
)

type gamePieces struct {
	_pieces    [6][7]int
	_nextPiece int
}

func New() gamePieces {
	g := gamePieces{}
	g._nextPiece = 1
	return g
}

func NewLoadExisting(pieces [6][7]int, nextPiece int) gamePieces {
	g := gamePieces{}
	g._pieces = pieces
	g._nextPiece = nextPiece
	return g
}

func (g gamePieces) CurrentState() [6][7]int {
	return g._pieces
}

func (g *gamePieces) AddPiece(col int) error {
	piecePlaced := false
	for i := 5; i >= 0; i-- {
		if g._pieces[i][col] == 0 {
			g._pieces[i][col] = g._nextPiece
			i = -1
			piecePlaced = true
			if g._nextPiece == 1 {
				g._nextPiece = 2
			} else {
				g._nextPiece = 1
			}
		}
	}

	if piecePlaced {
		return nil
	} else {
		return errors.New("Column " + strconv.Itoa(col+1) + " is full")
	}
}
