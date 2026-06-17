package simulation

import (
	"context"
	"simulation/internal/game/action"
	"simulation/internal/game/renderer"
	"simulation/internal/game/world"
	"time"
)

type Simulation struct {
	turn        uint64
	world       *world.World
	renderer    renderer.Renderer
	initActions []action.Action
	turnActions []action.Action
	interval    time.Duration
	pauseChan   chan struct{}
}

func New(world *world.World, delayMs int, render renderer.Renderer, initActions, turnActions []action.Action) *Simulation {
	simulation := &Simulation{
		turn:        0,
		world:       world,
		renderer:    render,
		initActions: initActions,
		turnActions: turnActions,
		interval:    time.Duration(delayMs) * time.Millisecond,
		pauseChan:   make(chan struct{}),
	}
	simulation.init()
	return simulation
}

// Start blocking operation, need invoke in a separate goroutine
func (s *Simulation) Start(ctx context.Context) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	paused := false
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.pauseChan:
			paused = !paused
		case <-ticker.C:
			if !paused {
				s.nextTurn()
			}
		}
	}
}

func (s *Simulation) Pause() {
	s.pauseChan <- struct{}{}
}

func (s *Simulation) Resume() {
	s.pauseChan <- struct{}{}
}

func (s *Simulation) nextTurn() {
	s.turn++
	for _, turnAction := range s.turnActions {
		turnAction.Execute(s.world)
	}

	s.printTurnHeader()
	s.renderer.Render(s.world)
}

func (s *Simulation) init() {
	for _, initAction := range s.initActions {
		initAction.Execute(s.world)
	}
}

func (s *Simulation) printTurnHeader() {
	println("Turn: ", s.turn)
}
