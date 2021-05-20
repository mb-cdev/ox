package auth

import (
	"net/http"
	"sync"
)

func init() {
	var once sync.Once

	once.Do(func() {
		registerHttpHandlers()
	})
}

func registerHttpHandlers() {
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

	})
}
