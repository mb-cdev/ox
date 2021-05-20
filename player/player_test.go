package player

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	player := NewPlayer("TestLogin")
	uuid, err := Logged.AddPlayer(&player)
	if err != nil {
		t.Error("Error#1", err)
	}

	p, err2 := Logged.GetPlayer(uuid)

	if err2 != nil {
		t.Error("Error#2", err2)
	}

	if p.Name != "TestLogin" {
		t.Error("Error while adding Player")
	}
}
