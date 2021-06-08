package room

import (
	"errors"
	"mb-cdev/ox/chat"
	"mb-cdev/ox/game"
	"mb-cdev/ox/player"
	"sync"

	"github.com/google/uuid"
)

var errUserNotExists = errors.New("one of user not exists")

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

func (r *Room) AppendParticipant(p *player.Player) {
	r.mu.Lock()
	if _, ok := r.Participants[p.Name]; !ok {
		r.Participants[p.Name] = p
	}
	r.mu.Unlock()
}

func (r *Room) DeleteParticipant(p *player.Player) {
	r.mu.Lock()
	if _, ok := r.Participants[p.Name]; !ok {
		delete(r.Participants, p.Name)
	}
	r.mu.Unlock()
}

func (r *Room) NewGame(player1_login, player2_login string) error {

	var p1 *player.Player
	var p2 *player.Player

	if p, ok := r.Participants[player1_login]; ok {
		p1 = p
	}
	if p, ok := r.Participants[player2_login]; ok {
		p2 = p
	}

	if p1 == nil || p2 == nil {
		return errUserNotExists
	}

	newGame := game.NewGame(p1, p2)
	if r.CurrentGame != nil {
		newGame.SetSubscriptionList(r.CurrentGame.GetSubscriptionList())
	}

	return nil
}
