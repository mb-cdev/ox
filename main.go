package main

import (
	_ "mb-cdev/ox/web/auth"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)
}
