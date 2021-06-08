package game_test

import (
	"encoding/json"
	"fmt"
	"mb-cdev/ox/game"
	"mb-cdev/ox/player"
	"testing"
)

func TestMovement(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	_, err := g.MakeMove(&p2, 1, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestWinning1(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 0, 0)
	won1, err1 := g.MakeMove(&p2, 1, 0)
	won2, err2 := g.MakeMove(&p2, 2, 0)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}

func TestWinning2(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 0, 1)
	won1, err1 := g.MakeMove(&p2, 1, 1)
	won2, err2 := g.MakeMove(&p2, 2, 1)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}

func TestWinning3(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 0, 2)
	won1, err1 := g.MakeMove(&p2, 1, 2)
	won2, err2 := g.MakeMove(&p2, 2, 2)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}

func TestWinning4(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 0, 0)
	won1, err1 := g.MakeMove(&p2, 1, 0)
	won2, err2 := g.MakeMove(&p2, 2, 0)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}
func TestWinning5(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 0, 0)
	won1, err1 := g.MakeMove(&p2, 0, 1)
	won2, err2 := g.MakeMove(&p2, 0, 2)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}
func TestWinning6(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 1, 0)
	won1, err1 := g.MakeMove(&p2, 1, 1)
	won2, err2 := g.MakeMove(&p2, 1, 2)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}
func TestWinning7(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 2, 0)
	won1, err1 := g.MakeMove(&p2, 2, 1)
	won2, err2 := g.MakeMove(&p2, 2, 2)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}
func TestWinning8(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 0, 0)
	won1, err1 := g.MakeMove(&p2, 1, 1)
	won2, err2 := g.MakeMove(&p2, 2, 2)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}
func TestWinning9(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 2, 0)
	won1, err1 := g.MakeMove(&p2, 1, 1)
	won2, err2 := g.MakeMove(&p2, 0, 2)

	if err != nil || err1 != nil || err2 != nil {
		t.Error(err, err1, err2)
	}
	if won || won1 || !won2 {
		t.Error("Winning bool not correct")
	}
}
func TestBoardToRunes(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	g := game.NewGame(&p1, &p2)
	won, err := g.MakeMove(&p2, 2, 0)
	won1, err1 := g.MakeMove(&p2, 1, 1)

	if err != nil || err1 != nil {
		t.Error(err, err1)
	}
	if won || won1 {
		t.Error("Winning bool not correct")
	}

	s := game.NewGameStatusResponse(g)
	b, _ := json.Marshal(s)
	fmt.Printf("%v", string(b))
}
