package model_test

import (
	"testing"

	troop "github.com/santo0/risk-game-go/internal/model"
)

func TestBasicBattle(t *testing.T) {
	playerOneId := 1
	playerTwoId := 2
	army1 := troop.Army{playerOneId, 10}
	army2 := troop.Army{playerTwoId, 15}
	if army1 == army2 {
		t.Fail()
	}
}
