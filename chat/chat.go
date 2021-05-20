package chat

import (
	"container/list"
	"mb-cdev/ox/player"
)

type Chat struct {
	subscribers *list.List
}

func NewChat() *Chat {
	subs := &list.List{}
	subs.Init()

	return &Chat{subscribers: subs}
}

func (c *Chat) SendMessage(sender *player.Player, text string) {
	msg := newMessage(sender, text)

	for e := c.subscribers.Front(); e != nil; e = e.Next() {
		e.Value.(ChatObserver)(msg.String())
	}
}

func (c *Chat) Subscribe(o ChatObserver) *Subscription {
	return &Subscription{c.subscribers.PushBack(o)}
}

func (c *Chat) Unsubscribe(s *Subscription) {
	c.subscribers.Remove(s.Element)
}
