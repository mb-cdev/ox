package middleware

import (
	"log"
	"mb-cdev/ox/player"
	"mb-cdev/ox/web/auth"
	"net/http"
)

type IsLogged struct{}

func (i *IsLogged) Execute(w http.ResponseWriter, r *http.Request) bool {
	var h string

	if r.Method == http.MethodOptions {
		return true
	}

	if h = r.Header.Get(auth.HTTP_HEADER_UUID); h == "" {
		w.WriteHeader(http.StatusUnauthorized)
		log.Default().Println("IsLogged middleware - unauthorized#1")
		return false
	}
	p, err := player.Logged.GetPlayer(h)

	if p == nil || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Default().Println("IsLogged middleware - unauthorized#2")
		return false
	}

	return true
}
