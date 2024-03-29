package player

import "time"

type Player struct {
	Name      string `validator:"notEmpty"`
	CreatedAt time.Time
}

func NewPlayer(name string) Player {
	return Player{Name: name, CreatedAt: time.Now()}
}
