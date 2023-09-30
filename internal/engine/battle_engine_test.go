package engine_test

import (
	"reflect"
	"testing"

	"github.com/santo0/risk-game-go/internal/engine"
	"github.com/santo0/risk-game-go/internal/model"

	constants "github.com/santo0/risk-game-go/internal"
)

func TestBattleEngine(t *testing.T) {
	btf := engine.CreateBattlefieldMap(constants.MAP_FILE, engine.DiceBattleConfiguration{})
	expectedPositions := map[int][]int{0: {2, 3, 4}, 1: {2}, 2: {0, 1}, 3: {0}, 4: {0}}
	if !reflect.DeepEqual(btf.GetPositions(), expectedPositions) {
		t.Fail()
	}
	army := model.Army{PlayerId: 1, Quantity: 10}
	btf.SetPositionArmy(0, army)
	if btf.GetPositionArmy(0) != army {
		t.Fail()
	}
}
