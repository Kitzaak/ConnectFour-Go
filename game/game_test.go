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
		t.Errorf("\ngot  >>> %q\nwant >>> %q", got, want)
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
		t.Errorf("\ngot  >>> %q\nwant >>> %q", got, want)
	}
}

func Test_add_second_piece(t *testing.T) {
	want := [6][7]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
	}

	g := game.New()
	g.AddPiece(3)
	g.AddPiece(3)
	got := g.CurrentState()

	if got != want {
		t.Errorf("\ngot  >>> %q\nwant >>> %q", got, want)
	}
}

func Test_can_load_game(t *testing.T) {
	want := [6][7]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
	}

	g := game.NewLoadExisting(want, 2)
	got := g.CurrentState()

	if got != want {
		t.Errorf("\ngot  >>> %q\nwant >>> %q", got, want)
	}

	want[0][3] = 2
	g.AddPiece(3)
	got = g.CurrentState()

	if got != want {
		t.Errorf("\ngot  >>> %q\nwant >>> %q", got, want)
	}

	err := g.AddPiece(3)

	if err == nil {
		t.Errorf("Need error for full column")
	}

}
