package chat

import (
	"container/list"
	"encoding/json"
	"mb-cdev/ox/player"
	"mb-cdev/ox/websocket_response"
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

	msgStruct := struct{ Msg string }{Msg: msg.String()}
	resp := websocket_response.Response{
		Operation: "MESSAGE",
		Data:      msgStruct,
	}
	m, _ := json.Marshal(resp)

	for e := c.subscribers.Front(); e != nil; e = e.Next() {
		e.Value.(ChatObserver)(string(m))
	}
}

func (c *Chat) Subscribe(o ChatObserver) *Subscription {
	return &Subscription{c.subscribers.PushBack(o)}
}

func (c *Chat) Unsubscribe(s *Subscription) {
	c.subscribers.Remove(s.Element)
}
