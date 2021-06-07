package websocket_chat

import (
	"log"
	"mb-cdev/ox/game_protocol"
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/web/auth"
	"mb-cdev/ox/web/web_room"
	"net/http"
	"strings"
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

	defer func() {
		w.connectedRoom.Chat.Unsubscribe(sub)
	}()

	w.connectedRoom.Chat.SendMessage(w.connectedPlayer, "CONNECTED!")
	for {
		select {
		case <-disconnect:
			log.Default().Println("Disconnecting!")
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
