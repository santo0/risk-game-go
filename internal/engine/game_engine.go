package engine

import player "github.com/santo0/risk-game-go/internal/model"

type Game struct {
	players     []player.Player
	round       int
	player_turn int
}

func CreateGame(players []player.Player) Game {
	//len(players)
	return Game{players, 0, 0}
}

func NextTurn(game Game) player.Player {
	game.player_turn += 1
	if game.player_turn >= len(game.players) {
		game.player_turn = 0
		game.round += 1
	}
	return game.players[game.player_turn]
}

func DoPlayerAction(action string, player player.Player) {
	// Desde aqui es cridaria a la resta de funcionalitats
}

func GetPlayers(game Game) []player.Player {
	return game.players
}

func GetRound(game Game) int {
	return game.round
}

func GetPlayerTurn(game Game) int {
	return game.player_turn
}
