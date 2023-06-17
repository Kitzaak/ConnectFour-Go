package game

import (
	"errors"
	"strconv"
)

type Game struct {
	pieces    [6][7]int
	nextPiece int
}

func New() Game {
	g := Game{}
	g.nextPiece = 1
	return g
}

func NewLoadExisting(pieces [6][7]int, nextPiece int) Game {
	g := Game{}
	g.pieces = pieces
	g.nextPiece = nextPiece
	return g
}

func (g Game) CurrentState() [6][7]int {
	return g.pieces
}

func (g Game) NextPiece() int {
	return 0
}

func (g *Game) AddPiece(col int) error {
	piecePlaced := false
	for i := 5; i >= 0; i-- {
		if g.pieces[i][col-1] == 0 {
			g.pieces[i][col-1] = g.nextPiece
			i = -1
			piecePlaced = true
			if g.nextPiece == 1 {
				g.nextPiece = 2
			} else {
				g.nextPiece = 1
			}
		}
	}

	if piecePlaced {
		return nil
	} else {
		return errors.New("Column " + strconv.Itoa(col) + " is full")
	}
}
