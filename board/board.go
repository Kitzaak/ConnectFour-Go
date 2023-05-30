package board

type Board struct{}

func DisplayEmpty() string {
	return "  1   2   3   4   5   6   7\n" +
		"-----------------------------\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"============================="
}

func DisplayPieces(pieces [6][7]int) string {
	var display = "  1   2   3   4   5   6   7\n" +
		"-----------------------------\n"
	var pieceDisplay string
	for i := 0; i < 6; i++ {
		var row = ""
		for j := 0; j < 7; j++ {
			if pieces[i][j] == 1 {
				pieceDisplay = "X"
			} else if pieces[i][j] == 2 {
				pieceDisplay = "O"
			} else {
				pieceDisplay = "."
			}
			row += "| " + pieceDisplay + " "
		}
		display += row + "|\n"
	}
	display += "============================="
	return display
}
