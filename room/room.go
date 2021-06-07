package room

import (
	"mb-cdev/ox/chat"
	"mb-cdev/ox/game"
	"mb-cdev/ox/player"
	"sync"

	"github.com/google/uuid"
)

type Room struct {
	mu           sync.Mutex
	Chat         *chat.Chat
	Owner        *player.Player
	Participants map[string]*player.Player
	CurrentGame  *game.Game
}

func NewRoom(owner *player.Player) (string, error) {
	uid, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	r := &Room{
		Chat:         chat.NewChat(),
		Owner:        owner,
		Participants: make(map[string]*player.Player),
	}

	uidS := uid.String()
	RoomList.Rooms.Store(uidS, r)

	return uidS, nil
}

func (r *Room) AppendParticipant(uuid string, p *player.Player) {
	r.mu.Lock()
	if _, ok := r.Participants[uuid]; !ok {
		r.Participants[uuid] = p
	}
	r.mu.Unlock()
}

func (r *Room) DeleteParticipant(uuid string) {
	r.mu.Lock()
	if _, ok := r.Participants[uuid]; !ok {
		delete(r.Participants, uuid)
	}
	r.mu.Unlock()
}
