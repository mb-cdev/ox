package websocket_game

import (
	"mb-cdev/ox/game_protocol"
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/web/auth"
	"mb-cdev/ox/web/web_room"
	"net/http"
	"strings"
)

type WebsocketGameHandler struct {
	connectedPlayer     *player.Player
	connectedRoom       *room.Room
	connectedPlayerUUID string
	connectedRoomUUID   string
}

func (w *WebsocketGameHandler) ConfirmHandshake(ch *http.Request, sh *http.Response) bool {
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

	w.connectedRoom.AppendParticipant(p)

	return true
}

func (w *WebsocketGameHandler) ServeConnection(in chan string, out chan string, disconnect chan bool) {
	sub := w.connectedRoom.Chat.Subscribe(func(msg string) {
		out <- msg
	})
	subGame := w.connectedRoom.Subscribe(func(data string) {
		out <- data
	})
	getParticipantsToken := game_protocol.NewToken("GETROOMPARTICIPANTS", nil)

	defer func() {
		w.connectedRoom.Chat.Unsubscribe(sub)
		w.connectedRoom.Unsubscribe(subGame)
		w.connectedRoom.DeleteParticipant(w.connectedPlayer)
		getParticipantsToken.Execute(w.connectedPlayer, w.connectedRoom)
	}()

	w.connectedRoom.Chat.SendMessage(w.connectedPlayer, "CONNECTED!")
	//force send player list
	getParticipantsToken.Execute(w.connectedPlayer, w.connectedRoom)
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
