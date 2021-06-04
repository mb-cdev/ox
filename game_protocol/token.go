package game_protocol

import (
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"reflect"
	"sync"
)

var tokenMap map[string]reflect.Type

func init() {
	var once sync.Once

	once.Do(func() {
		tokenMap = make(map[string]reflect.Type)
		tokenMap["MESSAGE"] = reflect.TypeOf(MessageToken{})
	})
}

type Token interface {
	Execute(p *player.Player, r *room.Room)
	SetArguments(args []string)
}

func newToken(name string, args []string) Token {
	if t, ok := tokenMap[name]; ok {

		tok := reflect.New(t).Interface().(Token)
		tok.SetArguments(args)

		return tok
	}
	return nil
}
