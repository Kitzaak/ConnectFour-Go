package game_test

import (
	"connect_four/game"
	"testing"
)

func Test_new_game_is_empty(t *testing.T) {
	want := [6][7]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}

	got := game.New().CurrentState()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test_add_first_piece(t *testing.T) {
	want := [6][7]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
	}

	g := game.New()
	g.AddPiece(3)
	got := g.CurrentState()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
