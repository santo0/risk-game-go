package engine_test

import (
	"reflect"
	"testing"

	constants "github.com/santo0/risk-game-go/internal"
	"github.com/santo0/risk-game-go/internal/engine"
	game_engine "github.com/santo0/risk-game-go/internal/engine"
	player "github.com/santo0/risk-game-go/internal/model"
)

func TestBasicGame(t *testing.T) {
	players := []player.Player{
		{PlayerName: "joe", PlayerId: 1},
		{PlayerName: "bill", PlayerId: 2},
		{PlayerName: "mike", PlayerId: 3},
	}
	btf := engine.CreateBattlefieldMap(constants.MAP_FILE, engine.DiceBattleConfiguration{})
	game := game_engine.CreateGame(players, btf)
	if !reflect.DeepEqual(game_engine.GetPlayers(game), players) {
		t.Fail()
	}
}
