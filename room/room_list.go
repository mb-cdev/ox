package room

import "sync"

var RoomList list

type list struct {
	Rooms sync.Map
}

func init() {
	var once sync.Once

	once.Do(func() {
		RoomList = list{Rooms: sync.Map{}}
	})
}
