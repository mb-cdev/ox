package game_protocol

import (
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/websocket_response"
)

type GetRoomParticipantsToken struct {
}

func (g *GetRoomParticipantsToken) SetArguments(args []string) {

}

func (g *GetRoomParticipantsToken) Execute(p *player.Player, r *room.Room) {
	participants := r.GetParticipants()
	sParticipants := make([]string, 0)

	for login := range participants {
		sParticipants = append(sParticipants, login)
	}

	resp := websocket_response.Response{
		Operation: "GETROOMPARTICIPANTS",
		Errors:    nil,
		Data:      sParticipants,
	}

	r.Broadcast(resp)
}
