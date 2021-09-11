package game_protocol

import (
	"mb-cdev/ox/game"
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/websocket_response"
)

type GameStatusToken struct{}

func (m *GameStatusToken) Execute(p *player.Player, r *room.Room) {

	resp := websocket_response.Response{}
	resp.Operation = "GAMESTATUS"

	if r.CurrentGame == nil {
		resp.Data = nil
	} else {
		status := game.NewGameStatusResponse(r.CurrentGame)
		resp.Data = status
	}

	r.Broadcast(resp)
}

func (m *GameStatusToken) SetArguments(args []string) {

}
