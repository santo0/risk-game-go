package engine

import (
	"bufio"
	"errors"
	"log"
	"math/rand"
	"os"
	"strings"

	troop "github.com/santo0/risk-game-go/internal/model"
)

type IDecisionGenerator interface {
	Decide(int, int) (Decision, error)
}

type DiceBattleConfiguration struct {
}

// For testing purposes
type DeterministicBattleConfiguration struct {
	battleSteps []Decision
	steps_index int
}

type Decision struct {
	attackerDelta int
	defenderDelta int
}

// TODO: Implementar la logica dels daus correcta del risk
func (btlCfg DiceBattleConfiguration) Decide(attackerUnits int, defenderUnits int) (Decision, error) {
	prob := rand.Float32()
	if prob >= 0.5 {
		return Decision{0, -1}, nil
	} else {
		return Decision{-1, 0}, nil
	}
}

// uses pointer so changes in steps_index can be kept
func (btlCfg *DeterministicBattleConfiguration) Decide(attackerUnits int, defenderUnits int) (Decision, error) {
	if btlCfg.steps_index >= len(btlCfg.battleSteps) {
		return Decision{}, errors.New("battle steps exceeded")
	}
	step := btlCfg.battleSteps[btlCfg.steps_index]
	btlCfg.steps_index += 1
	return step, nil
}

type AttackResult struct {
	defenderLosses int
	attackerLosses int
}

type Battlefield interface {
	ExecuteAttack(int, int) (AttackResult, error)
}

type BattlefieldMap struct {
	// map of positions with their adjacent positions
	positions map[int][]int
	// map of positions with their current armies
	armies_positions map[int]troop.Army
	decider          IDecisionGenerator
}

func (btf BattlefieldMap) GetPositions() map[int][]int {
	return btf.positions
}

// TODO: Llegir el fitxer i crear el mapa
func CreateBattlefieldMap(btfMapPath string, decider IDecisionGenerator) BattlefieldMap {
	file, err := os.Open(btfMapPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	positions := make(map[int][]int)
	armies_positions := make(map[int]troop.Army)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		lineSlice := strings.Split(scanner.Text(), ",")
		var s []int
		armies_positions[i] = troop.Army{PlayerId: -1}
		for j, c := range lineSlice {
			if c == "1" && i != j {
				s = append(s, j)
			}
		}
		positions[i] = s
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return BattlefieldMap{positions: positions, armies_positions: armies_positions, decider: decider}
}

func (battlefield BattlefieldMap) GetPositionArmy(position int) troop.Army {
	return battlefield.armies_positions[position]
}

func (battlefield BattlefieldMap) GetAdjacentPositions(position int) []int {
	return battlefield.positions[position]
}

func (battlefield BattlefieldMap) SetPositionArmy(position int, army troop.Army) {
	battlefield.armies_positions[position] = army
}

func (battlefield BattlefieldMap) ExecuteAttack(attackingPosition int, defendingPosition int) (AttackResult, error) {
	// Checkejar de que se pugui fer l'atac: que sigon adjacents i que hi hagin tropes suficients
	attackArmy := battlefield.armies_positions[attackingPosition]
	defendArmy := battlefield.armies_positions[defendingPosition]
	decision, err := battlefield.decider.Decide(attackArmy.Quantity, defendArmy.Quantity)
	if err == nil {
		return AttackResult{}, err
	}
	return AttackResult{attackerLosses: decision.attackerDelta, defenderLosses: decision.defenderDelta}, nil
}
