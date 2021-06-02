package auth

import (
	"encoding/json"
	"mb-cdev/ox/player"
	"mb-cdev/ox/validator"
	"net/http"
	"sync"
)

const HTTP_HEADER_UUID string = "X-PLAYER-UUID"

func init() {
	var once sync.Once

	once.Do(func() {
		registerHttpHandlers()
	})
}

func registerHttpHandlers() {
	http.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		r.ParseForm()
		login := r.PostFormValue("login")

		jEnc := json.NewEncoder(w)

		p := player.NewPlayer(login)
		if ok, errVal := validator.IsModelValid(p); !ok {
			w.WriteHeader(http.StatusBadRequest)

			respError := struct {
				Error   string
				Success bool
			}{
				errVal.Error(),
				false,
			}

			jEnc.Encode(respError)
			return
		}

		uuid, err := player.Logged.AddPlayer(&p)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		res := struct {
			Uuid    string
			Header  string
			Success bool
		}{
			uuid,
			HTTP_HEADER_UUID,
			true,
		}

		jEnc.Encode(res)
	})
}
