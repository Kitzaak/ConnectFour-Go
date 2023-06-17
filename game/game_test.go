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
		t.Errorf("\ngot  >>> %v\nwant >>> %v", got, want)
	}
}

func Test_first_piece_is_1(t *testing.T) {
	want := 1
	got := game.New().NextPiece()

	if got != want {
		t.Errorf("\ngot  >>> %v\nwant >>> %v", got, want)
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
	g.AddPiece(4)
	got := g.CurrentState()

	if got != want {
		t.Errorf("\ngot  >>> %v\nwant >>> %v", got, want)
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
	g.AddPiece(4)
	g.AddPiece(4)
	got := g.CurrentState()

	if got != want {
		t.Errorf("\ngot  >>> %v\nwant >>> %v", got, want)
	}
}

func Test_can_load_game_and_play_next_piece(t *testing.T) {
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
		t.Errorf("\ngot  >>> %v\nwant >>> %v", got, want)
	}

	want[0][3] = 2
	g.AddPiece(4)
	got = g.CurrentState()

	if got != want {
		t.Errorf("\ngot  >>> %v\nwant >>> %v", got, want)
	}
}
func Test_error_thrown_on_over_filled_col(t *testing.T) {
	want := [6][7]int{
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
	}

	g := game.NewLoadExisting(want, 1)
	err := g.AddPiece(4)

	if err == nil {
		t.Errorf("Need error for full column")
	}

	errMessage := "Column 4 is full"
	if err.Error() != errMessage {
		t.Errorf("\ngot  >>> %q\nwant >>> %q", err, errMessage)
	}

}
