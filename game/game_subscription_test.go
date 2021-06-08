package game_test

import (
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"testing"
)

func TestSubscription(t *testing.T) {
	p1 := player.NewPlayer("p1")
	p2 := player.NewPlayer("p2")

	rUuid, err := room.NewRoom(&p1)

	if err != nil {
		t.Error(err)
	}

	r, ok := room.RoomList.Rooms.Load(rUuid)
	if !ok {
		t.Error("Error while getting room")
	}
	room := r.(*room.Room)

	room.AppendParticipant(&p1)
	room.AppendParticipant(&p2)
}
