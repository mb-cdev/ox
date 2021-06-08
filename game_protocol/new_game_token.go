package game_protocol

import (
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
)

type NewGameToken struct {
	player1_login string
	player2_login string
}

func (n *NewGameToken) Execute(p *player.Player, r *room.Room) {
	r.NewGame(n.player1_login, n.player2_login)
}

func (n *NewGameToken) SetArguments(args []string) {
	n.player1_login = args[0]
	n.player2_login = args[1]
}
