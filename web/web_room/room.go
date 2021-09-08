package web_room

import (
	"encoding/json"
	"log"
	"mb-cdev/ox/middleware"
	"mb-cdev/ox/player"
	"mb-cdev/ox/room"
	"mb-cdev/ox/web/auth"
	"net/http"
	"sync"
)

const HTTP_HEADER_ROOM_UUID string = "X-ROOM-UUID"

type errorResponse struct {
	Error   string
	Success bool
}

type successResponse struct {
	RoomUuid string
	Header   string
	Success  bool
}

func init() {
	var once sync.Once

	once.Do(func() {
		registerHttpHandlers()
	})
}

func registerHttpHandlers() {
	http.HandleFunc("/room/create", middleware.Middleware(handleCreateRoom, &middleware.IsLogged{}))
	http.HandleFunc("/room/join", middleware.Middleware(handleJoinRoom, &middleware.IsLogged{}))
}

func handleJoinRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		return
	}

	room_uuid := r.FormValue("room_uuid")

	ro, ok := room.RoomList.Rooms.Load(room_uuid)
	_, okCast := ro.(*room.Room)
	encoder := json.NewEncoder(w)

	if ok && okCast {
		resp := successResponse{
			RoomUuid: room_uuid,
			Header:   HTTP_HEADER_ROOM_UUID,
			Success:  true,
		}
		encoder.Encode(resp)
		return
	}

	resp := errorResponse{
		Error:   "Room not found",
		Success: false,
	}
	encoder.Encode(resp)

}

func handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		return
	}

	jsonEnc := json.NewEncoder(w)
	h := r.Header.Get(auth.HTTP_HEADER_UUID)

	p, err := player.Logged.GetPlayer(h)

	if err != nil || p == nil {
		w.WriteHeader(http.StatusInternalServerError)

		resp := errorResponse{Success: false, Error: "Player is empty"}
		jsonEnc.Encode(resp)

		log.Default().Println("Player is empty", err)
		return
	}

	rUid, err2 := room.NewRoom(p)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)

		resp := errorResponse{Success: false, Error: "Error while creating room"}
		jsonEnc.Encode(resp)

		log.Default().Println("Error while creating room", err)
		return
	}

	resp := successResponse{Success: true, RoomUuid: rUid, Header: HTTP_HEADER_ROOM_UUID}
	jsonEnc.Encode(resp)
}
