package websocket_chat

import (
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/web/auth"
	"mb-cdev/ox/web/web_room"
	"net/http"
)

type WebsocketChatHandler struct {
	connectedPlayer *player.Player
	connectedRoom   *room.Room
}

func (w *WebsocketChatHandler) ConfirmHandshake(ch *http.Request, sh *http.Response) bool {
	p, err := player.Logged.GetPlayer(ch.FormValue(auth.HTTP_HEADER_UUID))
	r, ok := room.RoomList.Rooms.Load(ch.FormValue(web_room.HTTP_HEADER_ROOM_UUID))

	if !ok || err != nil || p == nil {
		sh.StatusCode = http.StatusInternalServerError
		sh.Status = http.StatusText(http.StatusInternalServerError)
		return false
	}

	w.connectedPlayer = p
	w.connectedRoom = r.(*room.Room)

	return true
}

func (w *WebsocketChatHandler) ServeConnection(in chan string, out chan string, disconnect chan bool) {
	sub := w.connectedRoom.Chat.Subscribe(func(msg string) {
		out <- msg
	})
	defer w.connectedRoom.Chat.Unsubscribe(sub)

	w.connectedRoom.Chat.SendMessage(w.connectedPlayer, "CONNECTED!")
listeners:
	for {
		select {
		case <-disconnect:
			break listeners
		case msg := <-in:
			w.connectedRoom.Chat.SendMessage(w.connectedPlayer, msg)
		}
	}

}
