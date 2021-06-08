package websocket_response

type Response struct {
	Operation string
	Errors    []string
	Data      interface{}
}
