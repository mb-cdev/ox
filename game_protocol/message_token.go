package game_protocol

import (
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
)

type MessageToken struct {
	msg string
}

func (m *MessageToken) SetArguments(args []string) {
	m.msg = args[0]
}

func (m *MessageToken) Execute(p *player.Player, r *room.Room) {
	r.Chat.SendMessage(p, m.msg)
}

func (m *MessageToken) GetMessage() string {
	return m.msg
}

func (m *MessageToken) SetMessage(ms string) *MessageToken {
	m.msg = ms
	return m
}
