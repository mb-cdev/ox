package room

import (
	"mb-cdev/ox/chat"

	"github.com/google/uuid"
)

type Room struct {
	Chat *chat.Chat
}

func NewRoom() (string, error) {
	uid, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	r := &Room{
		Chat: chat.NewChat(),
	}

	uidS := uid.String()
	RoomList.rooms.Store(uidS, r)

	return uidS, nil
}
