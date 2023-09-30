package engine

import model "github.com/santo0/risk-game-go/internal/model"

type Game struct {
	players     []model.Player
	round       int
	player_turn int
	battlefield Battlefield
}

func CreateGame(players []model.Player, battlefield Battlefield) Game {
	return Game{players: players, round: 0, player_turn: 0, battlefield: battlefield}
}

func (game *Game) NextTurn() model.Player {
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
func (game Game) GetCurrentPlayer() model.Player {
	return game.players[game.player_turn%len(game.players)]
}

func (game Game) GetPlayers() []model.Player {
	return game.players
}

func (game Game) GetRound() int {
	return game.round
}

func (game Game) GetPlayerTurn() int {
	return game.player_turn
}
