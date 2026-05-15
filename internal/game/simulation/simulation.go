package simulation

import (
	"simulation/internal/game/action"
	_map "simulation/internal/game/map"
)

type Simulation struct {
	IterationCounter uint64
	worldMap         *_map.Map
	//renderer
	initActions []action.Action
	turnActions []action.Action
}

func New(worldMap *_map.Map, initActions, turnActions []action.Action) *Simulation {
	simulation := &Simulation{
		IterationCounter: 1,
		worldMap:         worldMap,
		initActions:      initActions,
		turnActions:      turnActions,
	}
	simulation.init()
	return simulation
}

func (s *Simulation) init() {
	for _, initAction := range s.initActions {
		initAction.Execute(s.worldMap)
	}
}
