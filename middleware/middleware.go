package middleware

import "net/http"

type MiddlewareFunc interface {
	Execute(http.ResponseWriter, *http.Request) bool
}

func Middleware(next http.HandlerFunc, m ...MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, f := range m {
			if !f.Execute(w, r) {
				return
			}
		}

		next(w, r)
	}
}
