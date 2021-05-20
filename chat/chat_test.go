package chat

import (
	"fmt"
	"mb-cdev/ox/player"
	"testing"
)

func TestSubsciption(t *testing.T) {
	p1 := player.NewPlayer("P1")
	p2 := player.NewPlayer("P2")

	ch := NewChat()

	test := 123
	sub := ch.Subscribe(func(msg string) {
		fmt.Println(msg, test)
	})

	ch.SendMessage(&p1, "Test message")
	ch.SendMessage(&p2, "Test message")

	ch.Unsubscribe(sub)
}
