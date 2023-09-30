package engine_test

import (
	"fmt"
	"testing"

	"github.com/santo0/risk-game-go/internal/engine"

	constants "github.com/santo0/risk-game-go/internal"
)

func TestBattleEngine(t *testing.T) {
	btf := engine.CreateBattlefieldMap(constants.MAP_FILE, engine.DiceBattleConfiguration{})
	fmt.Println(btf.GetPositions())
}
