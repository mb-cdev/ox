package room

import (
	clist "container/list"
	"encoding/json"
	"errors"
	"log"
	"mb-cdev/ox/chat"
	"mb-cdev/ox/game"
	"mb-cdev/ox/player"
	"sync"

	"github.com/google/uuid"
)

var errUserNotExists = errors.New("one of user not exists")
var errTwoSameUsers = errors.New("can not play with two same users")

type Room struct {
	mu                sync.Mutex
	Chat              *chat.Chat
	Owner             *player.Player
	Participants      map[string]*player.Player
	CurrentGame       *game.Game
	gameSubscriptions *clist.List
}

func NewRoom(owner *player.Player) (string, error) {
	uid, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	r := &Room{
		Chat:              chat.NewChat(),
		Owner:             owner,
		Participants:      make(map[string]*player.Player),
		gameSubscriptions: &clist.List{},
	}

	r.gameSubscriptions.Init()

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

	if p1 == p2 {
		return errTwoSameUsers
	}

	if p1 == nil || p2 == nil {
		return errUserNotExists
	}

	r.CurrentGame = nil
	r.CurrentGame = game.NewGame(p1, p2)

	return nil
}

func (r *Room) Broadcast(response interface{}) {
	d, err := json.Marshal(response)
	if err != nil {
		log.Default().Println("Error while marshalling game response", err)
	}

	for e := r.gameSubscriptions.Front(); e != nil; e = e.Next() {
		e.Value.(GameObserver)(string(d))
	}
}

func (r *Room) Subscribe(f GameObserver) *Subscription {
	return &Subscription{r.gameSubscriptions.PushBack(f)}
}

func (r *Room) Unsubscribe(s *Subscription) {
	r.gameSubscriptions.Remove(s.Element)
}
