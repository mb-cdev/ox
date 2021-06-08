package game_protocol

import (
	"mb-cdev/ox/game"
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/websocket_response"
)

type NewGameToken struct {
	player1_login string
	player2_login string
}

func (n *NewGameToken) Execute(p *player.Player, r *room.Room) {
	err := r.NewGame(n.player1_login, n.player2_login)
	if err != nil {
		r.Broadcast(websocket_response.Response{
			Errors:    []string{err.Error()},
			Operation: "NEWGAME",
		})
		return
	}

	data := game.NewGameStatusResponse(r.CurrentGame)
	r.Broadcast(websocket_response.Response{
		Operation: "GAMESTATUS",
		Data:      data,
	})
}

func (n *NewGameToken) SetArguments(args []string) {
	n.player1_login = args[0]
	n.player2_login = args[1]
}
