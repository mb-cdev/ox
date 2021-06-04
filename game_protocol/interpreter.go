package game_protocol

import (
	"bufio"
	"io"
	"strings"
)

func ParseTokens(r io.Reader) []Token {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	var line uint8 = 0

	var token string
	arguments := make([]string, 0)
	tokens := make([]Token, 0)

	for s.Scan() {
		line++

		f := ExtractTokenOrArguments(s.Text(), line, &token, &arguments)

		if !f {
			t := newToken(token, arguments)
			if t == nil {
				continue
			}
			tokens = append(tokens, t)

			token = ""
			arguments = make([]string, 0)
			line = 0

		}
	}

	return tokens
}

//return false if there is no more arguments for token
func ExtractTokenOrArguments(line string, lineno uint8, token *string, arguments *[]string) bool {
	if lineno == 1 {
		*token = line
		return true
	}

	if lineno > 1 && len(line) > 0 {
		*arguments = append(*arguments, strings.Trim(line, "\r\n "))
		return true
	}

	return false
}
