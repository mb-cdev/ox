package player

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var Logged list

type list struct {
	players sync.Map
}

func (l *list) GetPlayer(uuid string) (*Player, error) {
	if v, ok := l.players.Load(uuid); ok {
		return v.(*Player), nil
	}

	return nil, errors.New("player not exists")
}

//return uuid new added player
func (l *list) AddPlayer(p *Player) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	idS := id.String()

	if _, ok := l.players.Load(idS); ok {
		return "", errors.New("player with this uuid exists")
	}

	l.players.Store(idS, p)

	return idS, nil
}

func init() {
	var once sync.Once

	once.Do(func() {
		Logged = list{players: sync.Map{}}
	})
}
