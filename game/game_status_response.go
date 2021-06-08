package game

type GameStatusResponse struct {
	Board    [3][3]string
	Finished bool
	Winned   bool
	Winner   string
}

func NewGameStatusResponse(g *Game) *GameStatusResponse {
	s := &GameStatusResponse{}
	s.GenerateBoard(g)
	s.Finished = g.finished
	s.Winned = g.Winned
	s.GetWinner(g)

	return s
}

func (s *GameStatusResponse) GenerateBoard(g *Game) {
	for ix, x := range g.board {
		for iy, val := range x {

			if val&cross == cross {
				s.Board[ix][iy] = "x"
			} else if val&circle == circle {
				s.Board[ix][iy] = "o"
			} else {
				s.Board[ix][iy] = "e"
			}

		}
	}
}

func (s *GameStatusResponse) GetWinner(g *Game) {
	if !g.Winned {
		s.Winner = ""
		return
	}

	s.Winner = g.players[g.turnForPlayerIndex].Name
}
