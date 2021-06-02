package room

import (
	"mb-cdev/ox/chat"
	"mb-cdev/ox/player"

	"github.com/google/uuid"
)

type Room struct {
	Chat  *chat.Chat
	Owner *player.Player
}

func NewRoom(owner *player.Player) (string, error) {
	uid, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	r := &Room{
		Chat:  chat.NewChat(),
		Owner: owner,
	}

	uidS := uid.String()
	RoomList.Rooms.Store(uidS, r)

	return uidS, nil
}
