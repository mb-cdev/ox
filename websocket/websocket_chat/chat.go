package websocket_chat

import (
	"mb-cdev/ox/game_protocol"
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/web/auth"
	"mb-cdev/ox/web/web_room"
	"net/http"
	"strings"
)

type WebsocketChatHandler struct {
	connectedPlayer     *player.Player
	connectedRoom       *room.Room
	connectedPlayerUUID string
	connectedRoomUUID   string
}

func (w *WebsocketChatHandler) ConfirmHandshake(ch *http.Request, sh *http.Response) bool {
	w.connectedPlayerUUID = ch.FormValue(auth.HTTP_HEADER_UUID)
	w.connectedRoomUUID = ch.FormValue(web_room.HTTP_HEADER_ROOM_UUID)

	p, err := player.Logged.GetPlayer(w.connectedPlayerUUID)
	r, ok := room.RoomList.Rooms.Load(w.connectedRoomUUID)

	if !ok || err != nil || p == nil {
		sh.StatusCode = http.StatusInternalServerError
		sh.Status = http.StatusText(http.StatusInternalServerError)
		return false
	}

	w.connectedPlayer = p
	w.connectedRoom = r.(*room.Room)

	w.connectedRoom.AppendParticipant(w.connectedPlayerUUID, p)

	return true
}

func (w *WebsocketChatHandler) ServeConnection(in chan string, out chan string, disconnect chan bool) {
	sub := w.connectedRoom.Chat.Subscribe(func(msg string) {
		out <- msg
	})

	defer func() {
		w.connectedRoom.Chat.Unsubscribe(sub)
		w.connectedRoom.DeleteParticipant(w.connectedPlayerUUID)
	}()

	w.connectedRoom.Chat.SendMessage(w.connectedPlayer, "CONNECTED!")
	for {
		select {
		case <-disconnect:
			w.connectedRoom.Chat.SendMessage(w.connectedPlayer, "DISCONNECTED")
			return
		case cmd := <-in:
			r := strings.NewReader(cmd)
			tokens := game_protocol.ParseTokens(r)

			for _, t := range tokens {
				t.Execute(w.connectedPlayer, w.connectedRoom)
			}
		}
	}

}
