package game

import (
	"errors"
	"math/rand"
	"mb-cdev/ox/player"
	"time"
)

var errBadPlayerIndex = errors.New("bad player index")
var errBadBoardIndex = errors.New("bad board index")
var errGameIsFinished = errors.New("game is finished")
var errNotPlayerTurn = errors.New("not player turn")
var errFieldUsed = errors.New("field already used")

type boardField byte

const (
	empty  boardField = 0x0
	circle boardField = 0x40
	cross  boardField = 0x80
)

type Game struct {
	players       [2]*player.Player
	playersSymbol [2]boardField

	// one board field:
	// LSB and next bit = player index
	// MSB and previous = boardField
	// 01000001 - player with index 1 selected circle
	// 10000000 - player with index 0 selected cross
	board [3][3]boardField

	turnForPlayerIndex uint8
	finished           bool
	Winned             bool
}

func NewGame(player1 *player.Player, player2 *player.Player) *Game {

	r := rand.New(rand.NewSource(time.Now().Unix()))
	var bf [2]boardField
	var pl [2]*player.Player

	if r.Intn(10) > 5 {
		bf[0] = cross
		bf[1] = circle
	} else {
		bf[0] = circle
		bf[1] = cross
	}

	if r.Intn(10) > 5 {
		pl[0] = player1
		pl[1] = player2
	} else {
		pl[0] = player2
		pl[1] = player1
	}

	g := &Game{
		board:              [3][3]boardField{},
		players:            pl,
		playersSymbol:      bf,
		turnForPlayerIndex: 0,
		finished:           false,
	}

	return g
}

//returning is movement winning and error
func (g *Game) MakeMove(pl *player.Player, x, y uint8) (bool, error) {

	if g.finished {
		return false, errGameIsFinished
	}

	pIndex := -1

	for i, p := range g.players {
		if p == pl {
			pIndex = i
		}
	}

	if pIndex == -1 || pIndex > 1 {
		return false, errBadPlayerIndex
	}

	if g.turnForPlayerIndex != uint8(pIndex) {
		return false, errNotPlayerTurn
	}

	if x > 2 || y > 2 {
		return false, errBadBoardIndex
	}

	if g.board[x][y] != empty {
		return false, errFieldUsed
	}

	moveType := g.playersSymbol[pIndex]
	movement := moveType | boardField(pIndex)

	g.board[x][y] |= movement

	if g.isMovementWinning(movement, x, y) {
		g.Winned = true
		g.finished = true
		return true, nil
	}

	g.turnForPlayerIndex++
	if (g.turnForPlayerIndex) > 1 {
		g.turnForPlayerIndex = 0
	}

	return false, nil
}

func (g *Game) isMovementWinning(movement boardField, x, y uint8) bool {

	wonX := true
	wonY := true
	wonDiagonal := true
	wonDiagonal2 := true

	for i := 0; i <= 2; i++ {
		if wonX && g.board[i][y] != movement {
			wonX = false
		}
		if wonY && g.board[x][i] != movement {
			wonY = false
		}
		if wonDiagonal && g.board[i][i] != movement {
			wonDiagonal = false
		}
		if wonDiagonal2 && g.board[2-i][i] != movement {
			wonDiagonal2 = false
		}

		if !wonX && !wonY && !wonDiagonal && !wonDiagonal2 {
			return false
		}
	}

	return wonX || wonY || wonDiagonal || wonDiagonal2
}
