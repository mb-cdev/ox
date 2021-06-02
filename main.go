package main

import (
	_ "mb-cdev/ox/web/auth"
	_ "mb-cdev/ox/web/web_room"
	"mb-cdev/ox/websocket/websocket_chat"
	"net/http"

	"canisdev.pl/websocket"
)

func main() {
	//HTTP
	go http.ListenAndServe(":8080", nil)

	//WebSocket
	mux := websocket.NewWebSocketMux()
	mux.Handle("/chat", &websocket_chat.WebsocketChatHandler{})
	websocket.ListenAndServe("0.0.0.0:9999", mux)
}
