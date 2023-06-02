package board_test

import (
	"connect_four/board"
	"testing"
)

func TestDisplayEmpty(t *testing.T) {
	got := board.DisplayEmpty()
	want := "  1   2   3   4   5   6   7\n" +
		"-----------------------------\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"============================="

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDisplayPieces(t *testing.T) {
	var pieces = [6][7]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 1, 2, 2, 0, 0, 0},
	}

	got := board.DisplayPieces(pieces)
	want := "  1   2   3   4   5   6   7\n" +
		"-----------------------------\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | . | . | . | . | . |\n" +
		"| . | . | X | . | . | . | . |\n" +
		"| . | X | O | O | . | . | . |\n" +
		"============================="

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
