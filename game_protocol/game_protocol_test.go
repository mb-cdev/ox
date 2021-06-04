package game_protocol_test

import (
	"bufio"
	"fmt"
	"mb-cdev/ox/game_protocol"
	"strings"
	"testing"
)

var message string = "MESSAGE\r\nHello Wazzup\r\n\r\nSECONDARY TOKEN\r\n\r\nThirdToken\r\nArg1\r\nArg2\r\nArg3\r\n\r\n"

func TestExtractTokens(t *testing.T) {
	r := strings.NewReader(message)
	s := bufio.NewScanner(r)

	var line uint8 = 0

	var token string
	arguments := make([]string, 0)

	for s.Scan() {
		line++

		f := game_protocol.ExtractTokenOrArguments(s.Text(), line, &token, &arguments)

		if !f {
			fmt.Printf("%#v %#v\n", token, arguments)

			token = ""
			arguments = make([]string, 0)
			line = 0
		}
	}
}

func TestNewTokens(t *testing.T) {
	r := strings.NewReader(message)
	tokens := game_protocol.ParseTokens(r)

	fmt.Printf("\n%#v\n", tokens[0].(*game_protocol.MessageToken))
}
