package model_test

import (
	"fmt"
	"testing"

	troop "github.com/santo0/risk-game-go/internal/model"
)

func TestBasicBattle(t *testing.T) {
	playerOneId := 1
	playerTwoId := 2
	army1 := troop.Army{playerOneId, 10}
	army2 := troop.Army{playerTwoId, 15}
	fmt.Printf("army1: %v\n", army1)
	fmt.Printf("army2: %v\n", army2)
}
