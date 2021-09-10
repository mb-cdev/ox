package game_protocol

import (
	"log"
	"mb-cdev/ox/game"
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/websocket_response"
	"strconv"
)

type MakeMoveToken struct {
	x, y uint8
}

func (m *MakeMoveToken) Execute(p *player.Player, r *room.Room) {
	if r.CurrentGame == nil {
		log.Default().Println("Game is NULL")
		return
	}

	_, err := r.CurrentGame.MakeMove(p, m.x, m.y)
	resp := websocket_response.Response{}

	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	status := game.NewGameStatusResponse(r.CurrentGame)
	resp.Data = status
	resp.Operation = "GAMESTATUS"

	r.Broadcast(resp)
}

func (m *MakeMoveToken) SetArguments(args []string) {
	x, errX := strconv.Atoi(args[0])
	if errX != nil {
		// to force bad index error
		m.x = 100
	} else {
		m.x = uint8(x)
	}

	y, errY := strconv.Atoi(args[1])
	if errY != nil {
		//to force bad index error
		m.y = 100
	} else {
		m.y = uint8(y)
	}
}
