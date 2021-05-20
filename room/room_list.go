package room

import "sync"

var RoomList list

type list struct {
	rooms sync.Map
}

func init() {
	var once sync.Once

	once.Do(func() {
		RoomList = list{rooms: sync.Map{}}
	})
}
