package engine

import model "github.com/santo0/risk-game-go/internal/model"

type Game struct {
	players     []model.Player
	round       int
	player_turn int
	battlefield Battlefield
}

func CreateGame(players []model.Player, battlefield Battlefield) Game {
	//len(players)
	return Game{players, 0, 0, battlefield}
}

func NextTurn(game Game) model.Player {
	game.player_turn += 1
	if game.player_turn >= len(game.players) {
		game.player_turn = 0
		game.round += 1
	}
	return game.players[game.player_turn]
}

func DoPlayerAction(action string, player model.Player) {
	// Desde aqui es cridaria a la resta de funcionalitats
}

func GetPlayers(game Game) []model.Player {
	return game.players
}

func GetRound(game Game) int {
	return game.round
}

func GetPlayerTurn(game Game) int {
	return game.player_turn
}
