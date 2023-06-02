package game

type gamePieces struct {
	pieces    [6][7]int
	nextPiece int
}

func New() gamePieces {
	g := gamePieces{}
	g.nextPiece = 1
	return g
}

func (g gamePieces) CurrentState() [6][7]int {
	return g.pieces
}

func (g *gamePieces) AddPiece(col int) {
	for i := 5; i >= 0; i-- {
		if g.pieces[i][col] == 0 {
			g.pieces[i][col] = g.nextPiece
			i = -1
		}
	}
}
