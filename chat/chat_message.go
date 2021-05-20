package chat

import (
	"mb-cdev/ox/player"
	"strings"
	"time"
)

type message struct {
	sender    *player.Player
	text      string
	createdAt time.Time
}

func newMessage(sender *player.Player, text string) *message {
	return &message{sender: sender, text: text, createdAt: time.Now()}
}

func (m *message) String() string {
	sb := strings.Builder{}

	sb.WriteRune('[')
	sb.WriteString(m.createdAt.String())
	sb.WriteRune(']')
	sb.WriteRune(' ')

	sb.WriteString(m.sender.Name)
	sb.WriteRune(':')
	sb.WriteRune(' ')

	sb.WriteString(m.text)

	return sb.String()
}
