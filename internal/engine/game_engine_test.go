package engine_test

import (
	"reflect"
	"testing"

	constants "github.com/santo0/risk-game-go/internal"
	"github.com/santo0/risk-game-go/internal/engine"
	player "github.com/santo0/risk-game-go/internal/model"
)

func TestBasicGameEngine(t *testing.T) {
	players := []player.Player{
		{PlayerName: "joe", PlayerId: 1},
		{PlayerName: "bill", PlayerId: 2},
		{PlayerName: "mike", PlayerId: 3},
	}
	btf := engine.CreateBattlefieldMap(constants.MAP_FILE, engine.DiceBattleConfiguration{})
	game := engine.CreateGame(players, btf)
	if !reflect.DeepEqual(game.GetPlayers(), players) {
		t.Fail()
	}
	numPlayers := len(players)
	for i := 0; i <= 9; i++ {
		currentPlayer := game.GetCurrentPlayer()
		if currentPlayer != players[i%numPlayers] {
			t.FailNow()
		}
		playerTurn := game.GetPlayerTurn()
		if playerTurn != i%numPlayers {
			t.FailNow()
		}
		nextPlayer := game.NextTurn()
		if nextPlayer != players[(i+1)%numPlayers] {
			t.FailNow()
		}

	}
}
