package game

import "mb-cdev/ox/player"

type Game struct {
	player1 *player.Player
	player2 *player.Player

	board [3][3]byte

	turnFor *player.Player
}
