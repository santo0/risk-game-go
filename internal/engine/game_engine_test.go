package engine_test

import (
	"reflect"
	"testing"

	game_engine "github.com/santo0/risk-game-go/internal/engine"
	player "github.com/santo0/risk-game-go/internal/model"
)

func TestBasicGame(t *testing.T) {
	players := []player.Player{
		{PlayerName: "joe", PlayerId: 1},
		{PlayerName: "bill", PlayerId: 2},
		{PlayerName: "mike", PlayerId: 3},
	}
	game := game_engine.CreateGame(players)
	// bad test aposta
	if !reflect.DeepEqual(game_engine.GetPlayers(game), players) {
		t.Fail()
	}
}
