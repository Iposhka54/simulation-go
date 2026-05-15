package simulation

import (
	"simulation/internal/game/action"
	_map "simulation/internal/game/map"
	"simulation/internal/game/renderer"
)

type Simulation struct {
	IterationCounter uint64
	worldMap         *_map.Map
	renderer         renderer.Renderer
	initActions      []action.Action
	turnActions      []action.Action
}

func New(worldMap *_map.Map, render renderer.Renderer, initActions, turnActions []action.Action) *Simulation {
	simulation := &Simulation{
		IterationCounter: 1,
		worldMap:         worldMap,
		renderer:         render,
		initActions:      initActions,
		turnActions:      turnActions,
	}
	simulation.init()
	return simulation
}

func (s *Simulation) PrintSimulation() {
	s.renderer.Render(s.worldMap)
}

func (s *Simulation) init() {
	for _, initAction := range s.initActions {
		initAction.Execute(s.worldMap)
	}
}
